/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */package aggregate

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	analytics "github.com/SeMI-network/janus-spark-analytics/clients/go"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/semi-technologies/weaviate/adapters/connectors/janusgraph/gremlin"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/kinds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_QueryProcessor_AnalyticsEngine(t *testing.T) {
	t.Run("when params specify not to use the analytics engine", func(t *testing.T) {
		params := paramsWithAnalyticsProps(filters.AnalyticsProps{})
		janusResponse := &gremlin.Response{
			Data: []gremlin.Datum{
				gremlin.Datum{
					Datum: map[string]interface{}{
						"true": map[string]interface{}{
							"population": map[string]interface{}{
								"prop_18__count": 8,
							},
						},
					},
				},
			},
		}
		executor := &fakeExecutor{result: janusResponse}
		expectedResult := []interface{}{
			map[string]interface{}{
				"groupedBy": map[string]interface{}{
					"path":  []interface{}{"convertible"},
					"value": "true",
				},
				"population": map[string]interface{}{
					"count": 8,
				},
			},
		}

		etcd := &etcdClientMock{}
		analytics := &analyticsAPIMock{}

		result, err := NewProcessor(executor, etcd, analytics).
			Process(context.Background(), gremlin.New(), params.GroupBy, &params)

		require.Nil(t, err, "should not error")
		assert.Equal(t, expectedResult, result, "result should be directly computed without the engine")

		// Make sure that neither etcd, nor the analytics API were involved
		etcd.AssertNotCalled(t, "Get")
		analytics.AssertNotCalled(t, "Schedule")
	})

	t.Run("when analytics engine should be used and cache is present", func(t *testing.T) {
		params := paramsWithAnalyticsProps(filters.AnalyticsProps{UseAnaltyicsEngine: true})
		executorResponse := &gremlin.Response{Data: []gremlin.Datum{}}
		executor := &fakeExecutor{result: executorResponse}
		expectedResult := []interface{}{
			map[string]interface{}{
				"groupedBy": map[string]interface{}{
					"path":  []interface{}{"convertible"},
					"value": "true",
				},
				"population": map[string]interface{}{
					"count": float64(8),
				},
			},
		}
		hash, err := params.AnalyticsHash()
		require.Nil(t, err)

		etcd := &etcdClientMock{}
		analyticsResult := []interface{}{
			map[string]interface{}{
				"true": map[string]interface{}{
					"population": map[string]interface{}{
						"prop_18__count": 8,
					},
				},
			},
		}
		etcd.On("Get", mock.Anything, keyFromHash(hash),
			[]clientv3.OpOption(nil)).
			Return(etcdResponse(analyticsResult, hash, analytics.StatusSucceeded), nil)

		analytics := &analyticsAPIMock{}

		result, err := NewProcessor(executor, etcd, analytics).
			Process(context.Background(), gremlin.New(), params.GroupBy, &params)

		require.Nil(t, err, "should not error")
		assert.Equal(t, expectedResult, result, "result should be directly computed without the engine")

		// Make sure the analytics API wasn't called (since we retrieved the result from the cache)
		etcd.AssertExpectations(t)
		analytics.AssertNotCalled(t, "Schedule")
	})

	t.Run("when analytics engine should be used and there is no cache at all", func(t *testing.T) {
		params := paramsWithAnalyticsProps(filters.AnalyticsProps{UseAnaltyicsEngine: true})
		executorResponse := &gremlin.Response{Data: []gremlin.Datum{}}
		executor := &fakeExecutor{result: executorResponse}
		hash, err := params.AnalyticsHash()
		require.Nil(t, err)

		etcd := &etcdClientMock{}
		etcd.On("Get", mock.Anything, keyFromHash(hash),
			[]clientv3.OpOption(nil)).
			Return(emptyEtcdResponse(), nil)

		scheduleParams := analytics.QueryParams{
			ID:    hash,
			Query: "g.V().count()", // exact query doesn't matter for this test, only that it matches
		}
		analytics := &analyticsAPIMock{}
		analytics.On("Schedule", mock.Anything, scheduleParams).
			Return(nil)

		_, err = NewProcessor(executor, etcd, analytics).
			Process(context.Background(), gremlin.New().Raw("g.V().count()"), nil, &params)

		etcd.AssertExpectations(t)
		analytics.AssertExpectations(t)

		expectedMessage := fmt.Errorf("could not process aggregate query: new job started - check back later: "+
			"the requested analytics request could not be served from cache, so a new analytics job "+
			"was triggered. This job runs in the background and can take considerable time depending "+
			"on the size of your graph. Please check back later. The id of your analysis job is '%s'.",
			hash)
		assert.Equal(t, expectedMessage, err)
	})

	t.Run("when analytics engine should be used and a job is ongoing", func(t *testing.T) {
		params := paramsWithAnalyticsProps(filters.AnalyticsProps{UseAnaltyicsEngine: true})
		executorResponse := &gremlin.Response{Data: []gremlin.Datum{}}
		executor := &fakeExecutor{result: executorResponse}
		hash, err := params.AnalyticsHash()
		require.Nil(t, err)

		etcd := &etcdClientMock{}
		etcd.On("Get", mock.Anything, keyFromHash(hash),
			[]clientv3.OpOption(nil)).
			Return(etcdResponse(nil, hash, analytics.StatusInProgress), nil)

		analytics := &analyticsAPIMock{}

		_, err = NewProcessor(executor, etcd, analytics).
			Process(context.Background(), gremlin.New().Raw("g.V().count()"), nil, &params)

		etcd.AssertExpectations(t)
		analytics.AssertNotCalled(t, "Schedule")

		expectedMessage := fmt.Errorf("could not process aggregate query: an analysis job matching your query "+
			"is already running with id '%s'. However, it hasn't finished yet. Please check back later.",
			hash)

		assert.Equal(t, expectedMessage, err)
	})

	t.Run("when analytics engine has failed", func(t *testing.T) {
		params := paramsWithAnalyticsProps(filters.AnalyticsProps{UseAnaltyicsEngine: true})
		executorResponse := &gremlin.Response{Data: []gremlin.Datum{}}
		executor := &fakeExecutor{result: executorResponse}
		hash, err := params.AnalyticsHash()
		require.Nil(t, err)

		etcd := &etcdClientMock{}
		analyticsErrorResponse := []interface{}{"spark died unexpectedly"}
		etcd.On("Get", mock.Anything, keyFromHash(hash),
			[]clientv3.OpOption(nil)).
			Return(etcdResponse(analyticsErrorResponse, hash, analytics.StatusFailed), nil)

		analytics := &analyticsAPIMock{}

		_, err = NewProcessor(executor, etcd, analytics).
			Process(context.Background(), gremlin.New().Raw("g.V().count()"), nil, &params)

		etcd.AssertExpectations(t)
		analytics.AssertNotCalled(t, "Schedule")

		expectedMessage := fmt.Errorf("could not process aggregate query: the previous analyis job matching "+
			"your query with id '%s' failed. To try again, set 'forceRecalculate' to 'true', the error message "+
			"from the previous failure was: %v", hash, analyticsErrorResponse)

		assert.Equal(t, expectedMessage, err)
	})

	t.Run("when forceRecalculate is set", func(t *testing.T) {
		params := paramsWithAnalyticsProps(filters.AnalyticsProps{
			UseAnaltyicsEngine: true,
			ForceRecalculate:   true,
		})
		executorResponse := &gremlin.Response{Data: []gremlin.Datum{}}
		executor := &fakeExecutor{result: executorResponse}
		hash, err := params.AnalyticsHash()
		require.Nil(t, err)

		etcd := &etcdClientMock{}

		scheduleParams := analytics.QueryParams{
			ID:    hash,
			Query: "g.V().count()", // exact query doesn't matter for this test, only that it matches
		}
		analytics := &analyticsAPIMock{}
		analytics.On("Schedule", mock.Anything, scheduleParams).
			Return(nil)

		_, err = NewProcessor(executor, etcd, analytics).
			Process(context.Background(), gremlin.New().Raw("g.V().count()"), nil, &params)

		analytics.AssertExpectations(t)

		// It should skip calling etcd
		etcd.AssertNotCalled(t, "Get")

		expectedMessage := fmt.Errorf("could not process aggregate query: new job started - check back later: "+
			"the requested analytics request could not be served from cache, so a new analytics job "+
			"was triggered. This job runs in the background and can take considerable time depending "+
			"on the size of your graph. Please check back later. The id of your analysis job is '%s'.",
			hash)
		assert.Equal(t, expectedMessage, err)
	})
}

func paramsWithAnalyticsProps(a filters.AnalyticsProps) kinds.AggregateParams {
	return kinds.AggregateParams{
		Kind:      kind.Thing,
		ClassName: schema.ClassName("Car"),
		Properties: []kinds.AggregateProperty{
			kinds.AggregateProperty{
				Name: "horsepower",
				Aggregators: []kinds.Aggregator{
					kinds.MeanAggregator,
				},
			},
		},
		GroupBy: &filters.Path{
			Class:    schema.ClassName("Car"),
			Property: schema.PropertyName("convertible"),
		},
		Analytics: a,
	}
}

type etcdClientMock struct {
	mock.Mock
}

func (m *etcdClientMock) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	args := m.Called(ctx, key, opts)
	return args.Get(0).(*clientv3.GetResponse), args.Error(1)
}

type analyticsAPIMock struct {
	mock.Mock
}

func (m *analyticsAPIMock) Schedule(ctx context.Context, params analytics.QueryParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func etcdResponse(data []interface{}, id string, status analytics.Status) *clientv3.GetResponse {
	res := analytics.Result{
		ID:            id,
		Status:        status,
		Result:        data,
		OriginalQuery: analytics.QueryParams{}, // doesn't matter
	}

	resBytes, _ := json.Marshal(res)

	return &clientv3.GetResponse{
		Count: 1,
		Kvs: []*mvccpb.KeyValue{
			&mvccpb.KeyValue{
				Value: resBytes,
			},
		},
	}
}

func emptyEtcdResponse() *clientv3.GetResponse {
	return &clientv3.GetResponse{
		Count: 0,
		Kvs:   []*mvccpb.KeyValue{},
	}
}

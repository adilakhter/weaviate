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
 */
package fetch

import (
	"context"
	"testing"

	"github.com/semi-technologies/weaviate/adapters/connectors/janusgraph/gremlin"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_QueryProcessor(t *testing.T) {
	t.Run("when bool count and groupCount are requested", func(t *testing.T) {

		janusResponse := &gremlin.Response{
			Data: []gremlin.Datum{
				gremlin.Datum{
					Datum: map[string]interface{}{
						"classId": []interface{}{
							"class_18",
						},
						"uuid": []interface{}{
							"some-uuid",
						},
					},
				},
				gremlin.Datum{
					Datum: map[string]interface{}{
						"classId": []interface{}{
							"class_18",
						},
						"uuid": []interface{}{
							"some-other-uuid",
						},
					},
				},
			},
		}
		executor := &fakeExecutor{result: janusResponse}
		expectedResult := []interface{}{
			map[string]interface{}{
				"beacon":    "weaviate://my-super-peer/things/some-uuid",
				"className": "City",
			},
			map[string]interface{}{
				"beacon":    "weaviate://my-super-peer/things/some-other-uuid",
				"className": "City",
			},
		}

		k := kind.Thing
		peerName := "my-super-peer"

		result, err := NewProcessor(executor, k, peerName, &fakeNameSource{}).Process(context.Background(), gremlin.New())

		require.Nil(t, err, "should not error")
		assert.Equal(t, expectedResult, result, "result should be merged and post-processed")
	})

}

type fakeExecutor struct {
	result *gremlin.Response
}

func (f *fakeExecutor) Execute(ctx context.Context, query gremlin.Gremlin) (*gremlin.Response, error) {
	return f.result, nil
}

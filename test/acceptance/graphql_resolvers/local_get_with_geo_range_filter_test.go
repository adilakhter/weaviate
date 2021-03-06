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
package test

import (
	"encoding/json"
	"testing"

	"github.com/semi-technologies/weaviate/test/acceptance/helper"
	"github.com/stretchr/testify/assert"
)

func TestLocalGetWithWithinGeoRangeFilter(t *testing.T) {
	t.Run("Only Dusseldorf should be within 100km of Dusseldorf", func(t *testing.T) {
		query := `
		{
			Local{
				Get {
					Things {
						City(where:{
							operator: WithinGeoRange
							path: ["location"]
							valueGeoRange: { geoCoordinates: {latitude: 51.225556, longitude: 6.782778} distance: { max: 100 } }
						}){
							name
							location {
								latitude
								longitude
							}
						}
					}
				}
			}
		}
		`
		result := AssertGraphQL(t, helper.RootAuth, query)
		cities := result.Get("Local", "Get", "Things", "City").AsSlice()

		expectedResults := []interface{}{
			map[string]interface{}{
				"name": "Dusseldorf",
				"location": map[string]interface{}{
					"latitude":  json.Number("51.225555"),
					"longitude": json.Number("6.782778"),
				},
			},
		}

		assert.Equal(t, expectedResults, cities)
	})

	t.Run("Dusseldorf and Amsterdam should be within 200km of Dusseldorf", func(t *testing.T) {
		query := `
		{
			Local{
				Get {
					Things {
						City(where:{
							operator: WithinGeoRange
							path: ["location"]
							valueGeoRange: { geoCoordinates: {latitude: 51.225556, longitude: 6.782778} distance: { max: 200 } }
						}){
							name
							location {
								latitude
								longitude
							}
						}
					}
				}
			}
		}
		`
		result := AssertGraphQL(t, helper.RootAuth, query)
		cities := result.Get("Local", "Get", "Things", "City").AsSlice()

		expectedResults := []interface{}{
			map[string]interface{}{
				"name": "Dusseldorf",
				"location": map[string]interface{}{
					"latitude":  json.Number("51.225555"),
					"longitude": json.Number("6.782778"),
				},
			},
			map[string]interface{}{
				"name": "Amsterdam",
				"location": map[string]interface{}{
					"latitude":  json.Number("52.36667"),
					"longitude": json.Number("4.9"),
				},
			},
		}

		assert.ElementsMatch(t, expectedResults, cities)
	})
}

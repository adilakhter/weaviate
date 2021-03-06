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

package get

import (
	"testing"

	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/kinds"
)

func Test_QueryBuilder(t *testing.T) {
	tests := testCases{
		{
			name: "with a Thing.City with a single primitive prop 'name'",
			inputParams: kinds.LocalGetParams{
				ClassName: "City",
				Properties: []kinds.SelectProperty{
					kinds.SelectProperty{
						IsPrimitive: true,
						Name:        "name",
					},
				},
				Kind: kind.Thing,
				Pagination: &filters.Pagination{
					Limit: 33,
				},
			},
			expectedQuery: `
			g.V().has("kind", "thing").hasLabel("class_18")
				.limit(33).path().by(valueMap())
			`,
		},
		{
			name: "without an explicit limit specified",
			inputParams: kinds.LocalGetParams{
				ClassName: "City",
				Properties: []kinds.SelectProperty{
					kinds.SelectProperty{
						IsPrimitive: true,
						Name:        "name",
					},
				},
				Kind: kind.Thing,
			},
			expectedQuery: `
			g.V().has("kind", "thing").hasLabel("class_18")
				.limit(20).path().by(valueMap())
			`,
		},
		{
			name: "with a Thing.City with a single primitive prop 'name' and a where filter",
			inputParams: kinds.LocalGetParams{
				ClassName: "City",
				Properties: []kinds.SelectProperty{
					kinds.SelectProperty{
						IsPrimitive: true,
						Name:        "name",
					},
				},
				Kind: kind.Thing,
				Pagination: &filters.Pagination{
					Limit: 33,
				},
				Filters: &filters.LocalFilter{
					Root: &filters.Clause{
						Value: &filters.Value{
							Value: "Amsterdam",
							Type:  schema.DataTypeString,
						},
						On: &filters.Path{
							Class:    schema.ClassName("City"),
							Property: schema.PropertyName("name"),
						},
						Operator: filters.OperatorEqual,
					},
				},
			},
			expectedQuery: `
			g.V().has("kind", "thing").hasLabel("class_18")
			  .union(has("prop_1", eq("Amsterdam")))
				.limit(33).path().by(valueMap())
			`,
		},
		{
			name: "with a Thing.City with a ref prop one level deep",
			inputParams: kinds.LocalGetParams{
				ClassName: "City",
				Properties: []kinds.SelectProperty{
					kinds.SelectProperty{
						IsPrimitive: true,
						Name:        "name",
					},
					kinds.SelectProperty{
						IsPrimitive: false,
						Name:        "inCountry",
						Refs: []kinds.SelectClass{
							kinds.SelectClass{
								ClassName: "Country",
								RefProperties: []kinds.SelectProperty{
									kinds.SelectProperty{
										IsPrimitive: true,
										Name:        "name",
									},
								},
							},
						},
					},
				},
				Kind: kind.Thing,
				Pagination: &filters.Pagination{
					Limit: 33,
				},
			},
			expectedQuery: `
			g.V().has("kind", "thing").hasLabel("class_18")
				.union(
				  optional(
						outE("prop_3").inV().hasLabel("class_19")
					)
				)
				.limit(33).path().by(valueMap())
			`,
		},
		{
			name: "with a Thing.City with a network ref prop one level deep",
			inputParams: kinds.LocalGetParams{
				ClassName: "City",
				Properties: []kinds.SelectProperty{
					kinds.SelectProperty{
						IsPrimitive: true,
						Name:        "name",
					},
					kinds.SelectProperty{
						IsPrimitive: false,
						Name:        "inCountry",
						Refs: []kinds.SelectClass{
							kinds.SelectClass{
								ClassName: "WeaviateB__Country",
								RefProperties: []kinds.SelectProperty{
									kinds.SelectProperty{
										IsPrimitive: true,
										Name:        "name",
									},
								},
							},
						},
					},
				},
				Kind: kind.Thing,
				Pagination: &filters.Pagination{
					Limit: 33,
				},
			},
			expectedQuery: `
			g.V().has("kind", "thing").hasLabel("class_18")
				.union(
				  optional(
						outE("prop_3").inV()
					)
				)
				.limit(33).path().by(valueMap())
			`,
		},
		{
			name: "with a Thing.City with a ref prop three levels deep",
			inputParams: kinds.LocalGetParams{
				ClassName: "City",
				Properties: []kinds.SelectProperty{
					kinds.SelectProperty{
						IsPrimitive: true,
						Name:        "name",
					},
					kinds.SelectProperty{
						IsPrimitive: false,
						Name:        "inCountry",
						Refs: []kinds.SelectClass{
							kinds.SelectClass{
								ClassName: "Country",
								RefProperties: []kinds.SelectProperty{
									kinds.SelectProperty{
										IsPrimitive: false,
										Name:        "inContinent",
										Refs: []kinds.SelectClass{
											kinds.SelectClass{
												ClassName: "Continent",
												RefProperties: []kinds.SelectProperty{
													kinds.SelectProperty{
														IsPrimitive: false,
														Name:        "onPlanet",
														Refs: []kinds.SelectClass{
															kinds.SelectClass{
																ClassName: "Planet",
																RefProperties: []kinds.SelectProperty{
																	kinds.SelectProperty{
																		IsPrimitive: true,
																		Name:        "name",
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Kind: kind.Thing,
				Pagination: &filters.Pagination{
					Limit: 33,
				},
			},
			expectedQuery: `
			g.V().has("kind", "thing").hasLabel("class_18")
				.union(
				  optional(outE("prop_3").inV().hasLabel("class_19"))
					.optional( outE("prop_13").inV().hasLabel("class_20"))
					.optional( outE("prop_23").inV().hasLabel("class_21"))
				)
				.limit(33).path().by(valueMap())
			`,
		},
	}

	tests.AssertQuery(t)
}

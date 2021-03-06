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
package filters

import (
	"testing"
	"time"

	"github.com/semi-technologies/weaviate/adapters/connectors/janusgraph/state"
	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_EmptyFilters(t *testing.T) {
	result, err := New(nil, nil).String()

	require.Nil(t, err, "no error should have occurred")
	assert.Equal(t, "", result, "should return an empty string")
}

func Test_SingleProperties(t *testing.T) {
	t.Run("with propertyType Int", func(t *testing.T) {
		t.Run("with various operators and valid values", func(t *testing.T) {
			tests := testCases{
				{"'City.population == 10000'", filters.OperatorEqual, `.union(has("population", eq(10000)))`},
				{"'City.population != 10000'", filters.OperatorNotEqual, `.union(has("population", neq(10000)))`},
				{"'City.population < 10000'", filters.OperatorLessThan, `.union(has("population", lt(10000)))`},
				{"'City.population <= 10000'", filters.OperatorLessThanEqual, `.union(has("population", lte(10000)))`},
				{"'City.population > 10000'", filters.OperatorGreaterThan, `.union(has("population", gt(10000)))`},
				{"'City.population >= 10000'", filters.OperatorGreaterThanEqual, `.union(has("population", gte(10000)))`},
			}

			tests.AssertFilter(t, "population", int(10000), schema.DataTypeInt)
		})

		t.Run("an invalid value", func(t *testing.T) {
			tests := testCases{{"should fail with wrong type", filters.OperatorEqual, ""}}

			// Note the mismatch between the specified type (arg4) and the actual type (arg3)
			tests.AssertFilterErrors(t, "population", "200", schema.DataTypeInt)
		})
	})

	t.Run("with propertyType Number (float)", func(t *testing.T) {
		t.Run("with various operators and valid values", func(t *testing.T) {
			tests := testCases{
				{"'City.energyConsumption == 953.280000'", filters.OperatorEqual, `.union(has("energyConsumption", eq(953.280000)))`},
				{"'City.energyConsumption != 953.280000'", filters.OperatorNotEqual, `.union(has("energyConsumption", neq(953.280000)))`},
				{"'City.energyConsumption < 953.280000'", filters.OperatorLessThan, `.union(has("energyConsumption", lt(953.280000)))`},
				{"'City.energyConsumption <= 953.280000'", filters.OperatorLessThanEqual, `.union(has("energyConsumption", lte(953.280000)))`},
				{"'City.energyConsumption > 953.280000'", filters.OperatorGreaterThan, `.union(has("energyConsumption", gt(953.280000)))`},
				{"'City.energyConsumption >= 953.280000'", filters.OperatorGreaterThanEqual, `.union(has("energyConsumption", gte(953.280000)))`},
			}

			tests.AssertFilter(t, "energyConsumption", float64(953.28), schema.DataTypeNumber)
		})

		t.Run("an invalid value", func(t *testing.T) {
			tests := testCases{{"should fail with wrong type", filters.OperatorEqual, ""}}

			// Note the mismatch between the specified type (arg4) and the actual type (arg3)
			tests.AssertFilterErrors(t, "energyConsumption", "200", schema.DataTypeNumber)
		})
	})

	t.Run("with propertyType date (time.Time)", func(t *testing.T) {
		t.Run("with various operators and valid values", func(t *testing.T) {
			dateString := "2017-08-17T12:47:00+02:00"
			dateTime, err := time.Parse(time.RFC3339, dateString)
			require.Nil(t, err)

			tests := testCases{
				{`City.foundedWhen == "2017-08-17T12:47:00+02:00"`, filters.OperatorEqual,
					`.union(has("foundedWhen", eq("2017-08-17T12:47:00+02:00")))`},
				{`City.foundedWhen != "2017-08-17T12:47:00+02:00"`, filters.OperatorNotEqual,
					`.union(has("foundedWhen", neq("2017-08-17T12:47:00+02:00")))`},
				{`City.foundedWhen < "2017-08-17T12:47:00+02:00"`, filters.OperatorLessThan,
					`.union(has("foundedWhen", lt("2017-08-17T12:47:00+02:00")))`},
				{`City.foundedWhen <= "2017-08-17T12:47:00+02:00"`, filters.OperatorLessThanEqual,
					`.union(has("foundedWhen", lte("2017-08-17T12:47:00+02:00")))`},
				{`City.foundedWhen > "2017-08-17T12:47:00+02:00"`, filters.OperatorGreaterThan,
					`.union(has("foundedWhen", gt("2017-08-17T12:47:00+02:00")))`},
				{`City.foundedWhen >= "2017-08-17T12:47:00+02:00"`, filters.OperatorGreaterThanEqual,
					`.union(has("foundedWhen", gte("2017-08-17T12:47:00+02:00")))`},
			}

			tests.AssertFilter(t, "foundedWhen", dateTime, schema.DataTypeDate)
		})

		t.Run("an invalid value", func(t *testing.T) {
			tests := testCases{{"should fail with wrong type", filters.OperatorEqual, ""}}

			// Note the mismatch between the specified type (arg4) and the actual type (arg3)
			tests.AssertFilterErrors(t, "foundedWhen", "200", schema.DataTypeDate)
		})
	})

	t.Run("with propertyType string", func(t *testing.T) {
		t.Run("with various operators and valid values", func(t *testing.T) {
			tests := testCases{
				{`'City.name == "Berlin"'`, filters.OperatorEqual, `.union(has("name", eq("Berlin")))`},
				{`'City.name != "Berlin"'`, filters.OperatorNotEqual, `.union(has("name", neq("Berlin")))`},
			}

			tests.AssertFilter(t, "name", "Berlin", schema.DataTypeString)
		})

		t.Run("with an operator that does not make sense for this type", func(t *testing.T) {
			tests := testCases{
				{`City.name < "Berlin"`, filters.OperatorLessThan, ""},
				{`City.name <= "Berlin"`, filters.OperatorLessThanEqual, ""},
				{`City.name > "Berlin"`, filters.OperatorGreaterThan, ""},
				{`City.name >= "Berlin"`, filters.OperatorGreaterThanEqual, ""},
			}

			tests.AssertFilterErrors(t, "name", "Berlin", schema.DataTypeString)
		})

		t.Run("an invalid value", func(t *testing.T) {
			tests := testCases{{"should fail with wrong type", filters.OperatorEqual, ""}}

			// Note the mismatch between the specified type (arg4) and the actual type (arg3)
			tests.AssertFilterErrors(t, "name", int(200), schema.DataTypeString)
		})
	})

	t.Run("with propertyType bool", func(t *testing.T) {
		t.Run("with various operators and valid values", func(t *testing.T) {
			tests := testCases{
				{`'City.isCapital == true'`, filters.OperatorEqual, `.union(has("isCapital", eq(true)))`},
				{`'City.isCapital != true'`, filters.OperatorNotEqual, `.union(has("isCapital", neq(true)))`},
			}

			tests.AssertFilter(t, "isCapital", true, schema.DataTypeBoolean)
		})

		t.Run("with an operator that does not make sense for this type", func(t *testing.T) {
			tests := testCases{
				{`City.isCapital < true`, filters.OperatorLessThan, ""},
				{`City.isCapital <= true`, filters.OperatorLessThanEqual, ""},
				{`City.isCapital > true`, filters.OperatorGreaterThan, ""},
				{`City.isCapital >= true`, filters.OperatorGreaterThanEqual, ""},
			}

			tests.AssertFilterErrors(t, "isCapital", true, schema.DataTypeBoolean)
		})

		t.Run("an invalid value", func(t *testing.T) {
			tests := testCases{{"should fail with wrong type", filters.OperatorEqual, ""}}

			// Note the mismatch between the specified type (arg4) and the actual type (arg3)
			tests.AssertFilterErrors(t, "isCapital", int(200), schema.DataTypeBoolean)
		})
	})

	t.Run("with propertyType valueGeoRange", func(t *testing.T) {
		t.Run("with various operators and valid values", func(t *testing.T) {
			tests := testCases{
				{`'City.location within range'`, filters.OperatorWithinGeoRange,
					`.union(has("location", geoWithin(Geoshape.circle(51.200001, 6.700000, 190.000000))))`},
			}

			expectedValue := filters.GeoRange{
				GeoCoordinates: &models.GeoCoordinates{
					Latitude:  51.2,
					Longitude: 6.7,
				},
				Distance: 190,
			}
			tests.AssertFilter(t, "location", expectedValue, schema.DataTypeGeoCoordinates)
		})

		t.Run("with an operator that does not make sense for this type", func(t *testing.T) {
			tests := testCases{
				{`less than`, filters.OperatorLessThan, ""},
				{`less than equal`, filters.OperatorLessThanEqual, ""},
				{`greater than`, filters.OperatorGreaterThan, ""},
				{`greater than equal`, filters.OperatorGreaterThanEqual, ""},
				{`equal`, filters.OperatorEqual, ""},
				{`not equal`, filters.OperatorNotEqual, ""},
			}

			expectedValue := filters.GeoRange{
				GeoCoordinates: &models.GeoCoordinates{
					Latitude:  51.2,
					Longitude: 6.7,
				},
				Distance: 190,
			}

			tests.AssertFilterErrors(t, "location", expectedValue, schema.DataTypeGeoCoordinates)
		})

		t.Run("an invalid value", func(t *testing.T) {
			tests := testCases{{"should fail with wrong type", filters.OperatorWithinGeoRange, ""}}

			// Note the mismatch between the specified type (arg4) and the actual type (arg3)
			tests.AssertFilterErrors(t, "location", int(200), schema.DataTypeGeoCoordinates)
		})
	})
}

func Test_SinglePropertiesWithMappedNames(t *testing.T) {
	tests := testCases{
		{"'City.population == 10000'", filters.OperatorEqual, `.union(has("prop_20", eq(10000)))`},
		{"'City.population != 10000'", filters.OperatorNotEqual, `.union(has("prop_20", neq(10000)))`},
		{"'City.population < 10000'", filters.OperatorLessThan, `.union(has("prop_20", lt(10000)))`},
		{"'City.population <= 10000'", filters.OperatorLessThanEqual, `.union(has("prop_20", lte(10000)))`},
		{"'City.population > 10000'", filters.OperatorGreaterThan, `.union(has("prop_20", gt(10000)))`},
		{"'City.population >= 10000'", filters.OperatorGreaterThanEqual, `.union(has("prop_20", gte(10000)))`},
	}

	tests.AssertFilterWithNameSource(t, "population", int(10000), schema.DataTypeInt, &fakeNameSource{})
}

func Test_InvalidOperator(t *testing.T) {
	filter := buildFilter("population", "200", filters.Operator(27), schema.DataTypeInt)

	_, err := New(filter, nil).String()

	assert.NotNil(t, err, "it should error due to the wrong type")
}

func Test_MultipleConditions(t *testing.T) {
	buildOperandFilter := func(operator filters.Operator) *filters.LocalFilter {
		return &filters.LocalFilter{
			Root: &filters.Clause{
				Operator: operator,
				Operands: []filters.Clause{
					filters.Clause{
						Operator: filters.OperatorGreaterThan,
						On: &filters.Path{
							Class:    schema.ClassName("City"),
							Property: schema.PropertyName("population"),
						},
						Value: &filters.Value{
							Value: int(70000),
							Type:  schema.DataTypeInt,
						},
					},
					filters.Clause{
						Operator: filters.OperatorNotEqual,
						On: &filters.Path{
							Class:    schema.ClassName("City"),
							Property: schema.PropertyName("name"),
						},
						Value: &filters.Value{
							Value: "Rotterdam",
							Type:  schema.DataTypeString,
						},
					},
				},
			},
		}
	}

	t.Run("with operator and", func(t *testing.T) {
		filter := buildOperandFilter(filters.OperatorAnd)
		expectedResult := `.union(and(has("population", gt(70000)), has("name", neq("Rotterdam"))))`

		result, err := New(filter, nil).String()

		require.Nil(t, err, "should not error")
		assert.Equal(t, expectedResult, result, "should match the gremlin query")
	})

	t.Run("with operator or", func(t *testing.T) {
		filter := buildOperandFilter(filters.OperatorOr)
		expectedResult := `.union(or(has("population", gt(70000)), has("name", neq("Rotterdam"))))`

		result, err := New(filter, nil).String()

		require.Nil(t, err, "should not error")
		assert.Equal(t, expectedResult, result, "should match the gremlin query")
	})
}

func Test_MultipleNestedConditions(t *testing.T) {
	buildOperandFilter := func(operator filters.Operator) *filters.LocalFilter {
		return &filters.LocalFilter{
			Root: &filters.Clause{
				Operator: operator,
				Operands: []filters.Clause{
					filters.Clause{
						Operator: filters.OperatorGreaterThan,
						On: &filters.Path{
							Class:    schema.ClassName("City"),
							Property: schema.PropertyName("population"),
						},
						Value: &filters.Value{
							Value: int(70000),
							Type:  schema.DataTypeInt,
						},
					},
					filters.Clause{
						Operator: filters.OperatorOr,
						Operands: []filters.Clause{
							filters.Clause{
								Operator: filters.OperatorEqual,
								On: &filters.Path{
									Class:    schema.ClassName("City"),
									Property: schema.PropertyName("name"),
								},
								Value: &filters.Value{
									Value: "Rotterdam",
									Type:  schema.DataTypeString,
								},
							},
							filters.Clause{
								Operator: filters.OperatorEqual,
								On: &filters.Path{
									Class:    schema.ClassName("City"),
									Property: schema.PropertyName("name"),
								},
								Value: &filters.Value{
									Value: "Berlin",
									Type:  schema.DataTypeString,
								},
							},
						},
					},
				},
			},
		}
	}

	filter := buildOperandFilter(filters.OperatorAnd)
	expectedResult := `.union(and(has("population", gt(70000)), or(has("name", eq("Rotterdam")), has("name", eq("Berlin")))))`

	result, err := New(filter, nil).String()

	require.Nil(t, err, "should not error")
	assert.Equal(t, expectedResult, result, "should match the gremlin query")
}

func Test_FiltersOnRefProps(t *testing.T) {
	t.Run("one level deep", func(t *testing.T) {
		// city where InCountry->City->Country->name == "Germany"
		buildOperandFilter := func() *filters.LocalFilter {
			return &filters.LocalFilter{
				Root: &filters.Clause{
					Operator: filters.OperatorEqual,
					On: &filters.Path{
						Class:    schema.ClassName("City"),
						Property: schema.PropertyName("inCountry"),
						Child: &filters.Path{
							Class:    schema.ClassName("Country"),
							Property: schema.PropertyName("name"),
						},
					},
					Value: &filters.Value{
						Value: "Germany",
						Type:  schema.DataTypeString,
					},
				},
			}
		}

		filter := buildOperandFilter()
		expectedResult := `.union(where(outE("inCountry").inV().has("classId", "Country").has("name", eq("Germany"))))`

		result, err := New(filter, nil).String()

		require.Nil(t, err, "should not error")
		assert.Equal(t, expectedResult, result, "should match the gremlin query")
	})

	t.Run("multiple levels deep", func(t *testing.T) {
		// airport where InCity->City->InCountry->City->Country->name == "Germany"
		buildOperandFilter := func() *filters.LocalFilter {
			return &filters.LocalFilter{
				Root: &filters.Clause{
					Operator: filters.OperatorEqual,
					On: &filters.Path{
						Class:    schema.ClassName("Airport"),
						Property: schema.PropertyName("inCity"),
						Child: &filters.Path{
							Class:    schema.ClassName("City"),
							Property: schema.PropertyName("inCountry"),
							Child: &filters.Path{
								Class:    schema.ClassName("Country"),
								Property: schema.PropertyName("name"),
							},
						},
					},
					Value: &filters.Value{
						Value: "Germany",
						Type:  schema.DataTypeString,
					},
				},
			}
		}

		filter := buildOperandFilter()
		expectedResult := `.union(where(outE("inCity").inV().has("classId", "City")` +
			`.outE("inCountry").inV().has("classId", "Country").has("name", eq("Germany"))))`

		result, err := New(filter, nil).String()

		require.Nil(t, err, "should not error")
		assert.Equal(t, expectedResult, result, "should match the gremlin query")
	})

	t.Run("with mapped names", func(t *testing.T) {
		buildOperandFilter := func() *filters.LocalFilter {
			return &filters.LocalFilter{
				Root: &filters.Clause{
					Operator: filters.OperatorEqual,
					On: &filters.Path{
						Class:    schema.ClassName("City"),
						Property: schema.PropertyName("inCountry"),
						Child: &filters.Path{
							Class:    schema.ClassName("Country"),
							Property: schema.PropertyName("name"),
						},
					},
					Value: &filters.Value{
						Value: "Germany",
						Type:  schema.DataTypeString,
					},
				},
			}
		}

		filter := buildOperandFilter()
		expectedResult := `.union(where(outE("prop_15").inV().has("classId", "class_18").has("prop_20", eq("Germany"))))`

		result, err := New(filter, &fakeNameSource{}).String()

		require.Nil(t, err, "should not error")
		assert.Equal(t, expectedResult, result, "should match the gremlin query")
	})
}

func buildFilter(propName string, value interface{}, operator filters.Operator, schemaType schema.DataType) *filters.LocalFilter {
	return &filters.LocalFilter{
		Root: &filters.Clause{
			Operator: operator,
			On: &filters.Path{
				Class:    schema.ClassName("City"),
				Property: schema.PropertyName(propName),
			},
			Value: &filters.Value{
				Value: value,
				Type:  schemaType,
			},
		},
	}
}

type testCase struct {
	name           string
	operator       filters.Operator
	expectedResult string
}

type testCases []testCase

func (tests testCases) AssertFilter(t *testing.T, propName string, propValue interface{}, propType schema.DataType) {
	tests.AssertFilterWithNameSource(t, propName, propValue, propType, nil)
}

func (tests testCases) AssertFilterWithNameSource(t *testing.T, propName string, propValue interface{}, propType schema.DataType, nameSource nameSource) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filter := buildFilter(propName, propValue, test.operator, propType)

			result, err := New(filter, nameSource).String()

			require.Nil(t, err, "no error should have occurred")
			assert.Equal(t, test.expectedResult, result, "should form the right query")
		})
	}
}

func (tests testCases) AssertFilterErrors(t *testing.T, propName string, propValue interface{}, propType schema.DataType) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filter := buildFilter(propName, propValue, test.operator, propType)

			_, err := New(filter, nil).String()

			assert.NotNil(t, err, "should error")
		})
	}
}

type fakeNameSource struct{}

func (f *fakeNameSource) MustGetMappedPropertyName(className schema.ClassName,
	propName schema.PropertyName) state.MappedPropertyName {
	switch propName {
	case schema.PropertyName("inCountry"):
		return "prop_15"
	}
	return state.MappedPropertyName("prop_20")
}

func (f *fakeNameSource) MustGetMappedClassName(className schema.ClassName) state.MappedClassName {
	return state.MappedClassName("class_18")
}

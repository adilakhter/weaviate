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
package gremlin

import (
	"fmt"
)

type Vertex struct {
	Id         int
	Label      string
	Properties map[string]Property
}

func (v *Vertex) AssertPropertyValue(name string) *PropertyValue {
	prop := v.PropertyValue(name)

	if prop == nil {
		panic(fmt.Sprintf("Expected to find a property '%v' on vertex '%v', but no such property exists!", name, v.Id))
	}

	return prop
}

func (v *Vertex) PropertyValue(name string) *PropertyValue {
	val, ok := v.Properties[name]
	if !ok {
		return nil
	} else {
		return &val.Value
	}
}

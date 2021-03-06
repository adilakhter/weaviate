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
 */package kinds

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/usecases/network/common/peers"
)

type fakeRepo struct {
	GetThingResponse     *models.Thing
	UpdateThingParameter *models.Thing
}

func (f *fakeRepo) AddAction(ctx context.Context, class *models.Action, id strfmt.UUID) error {
	panic("not implemented")
}

func (f *fakeRepo) AddThing(ctx context.Context, class *models.Thing, id strfmt.UUID) error {
	panic("not implemented")
}

func (f *fakeRepo) GetThing(ctx context.Context, id strfmt.UUID, thing *models.Thing) error {
	*thing = *f.GetThingResponse
	return nil
}

func (f *fakeRepo) GetAction(context.Context, strfmt.UUID, *models.Action) error {
	panic("not implemented")
}

func (f *fakeRepo) ListThings(ctx context.Context, limit int, thingsResponse *models.ThingsListResponse) error {
	panic("not implemented")
}

func (f *fakeRepo) ListActions(ctx context.Context, limit int, actionsResponse *models.ActionsListResponse) error {
	panic("not implemented")
}

func (f *fakeRepo) UpdateAction(ctx context.Context, class *models.Action, id strfmt.UUID) error {
	panic("not implemented")
}

func (f *fakeRepo) UpdateThing(ctx context.Context, class *models.Thing, id strfmt.UUID) error {
	f.UpdateThingParameter = class
	return nil
}

func (f *fakeRepo) DeleteThing(ctx context.Context, thing *models.Thing, UUID strfmt.UUID) error {
	panic("not implemented")
}

func (f *fakeRepo) DeleteAction(ctx context.Context, thing *models.Action, UUID strfmt.UUID) error {
	panic("not implemented")
}

func (f *fakeRepo) AddThingsBatch(ctx context.Context, things BatchThings) error {
	panic("not implemented")
}

func (f *fakeRepo) AddActionsBatch(ctx context.Context, actions BatchActions) error {
	panic("not implemented")
}

func (f *fakeRepo) AddBatchReferences(ctx context.Context, references BatchReferences) error {
	panic("not implemented")
}

type fakeSchemaManager struct {
	CalledWith struct {
		kind      kind.Kind
		fromClass string
		property  string
		toClass   string
	}
	GetSchemaResponse schema.Schema
}

func (f *fakeSchemaManager) UpdatePropertyAddDataType(ctx context.Context, k kind.Kind, fromClass, property, toClass string) error {
	f.CalledWith = struct {
		kind      kind.Kind
		fromClass string
		property  string
		toClass   string
	}{
		kind:      k,
		fromClass: fromClass,
		property:  property,
		toClass:   toClass,
	}
	return nil
}

func (f *fakeSchemaManager) GetSchema() schema.Schema {
	return f.GetSchemaResponse
}

type fakeLocks struct{}

func (f *fakeLocks) LockConnector() (func() error, error) {
	return func() error { return nil }, nil
}

func (f *fakeLocks) LockSchema() (func() error, error) {
	return func() error { return nil }, nil
}

type fakeNetwork struct {
	peerURI string
}

func (f *fakeNetwork) ListPeers() (peers.Peers, error) {
	myPeers := peers.Peers{
		peers.Peer{
			Name: "BestWeaviate",
			URI:  strfmt.URI(f.peerURI),
			Schema: schema.Schema{
				Things: &models.SemanticSchema{
					Classes: []*models.SemanticSchemaClass{
						&models.SemanticSchemaClass{
							Class: "BestThing",
						},
					},
				},
			},
		},
	}

	return myPeers, nil
}

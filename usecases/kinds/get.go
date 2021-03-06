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
	"errors"
	"math"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
)

type getRepo interface {
	GetThing(context.Context, strfmt.UUID, *models.Thing) error
	GetAction(context.Context, strfmt.UUID, *models.Action) error

	ListThings(ctx context.Context, limit int, thingsResponse *models.ThingsListResponse) error
	ListActions(ctx context.Context, limit int, actionsResponse *models.ActionsListResponse) error
}

// GetThing Class from the connected DB
func (m *Manager) GetThing(ctx context.Context, id strfmt.UUID) (*models.Thing, error) {
	unlock, err := m.locks.LockConnector()
	if err != nil {
		return nil, newErrInternal("could not aquire lock: %v", err)
	}
	defer unlock()

	return m.getThingFromRepo(ctx, id)
}

// GetThings Class from the connected DB
func (m *Manager) GetThings(ctx context.Context, limit *int64) ([]*models.Thing, error) {
	unlock, err := m.locks.LockConnector()
	if err != nil {
		return nil, newErrInternal("could not aquire lock: %v", err)
	}
	defer unlock()

	return m.getThingsFromRepo(ctx, limit)
}

// GetAction Class from connected DB
func (m *Manager) GetAction(ctx context.Context, id strfmt.UUID) (*models.Action, error) {
	unlock, err := m.locks.LockConnector()
	if err != nil {
		return nil, newErrInternal("could not aquire lock: %v", err)
	}
	defer unlock()

	return m.getActionFromRepo(ctx, id)
}

// GetActions Class from connected DB
func (m *Manager) GetActions(ctx context.Context, limit *int64) ([]*models.Action, error) {
	unlock, err := m.locks.LockConnector()
	if err != nil {
		return nil, newErrInternal("could not aquire lock: %v", err)
	}
	defer unlock()

	return m.getActionsFromRepo(ctx, limit)
}

func (m *Manager) getThingFromRepo(ctx context.Context, id strfmt.UUID) (*models.Thing, error) {
	thing := models.Thing{}
	thing.Schema = map[string]models.JSONObject{}
	err := m.repo.GetThing(ctx, id, &thing)
	if err != nil {
		switch err {
		// TODO: Replace with structured error
		case errors.New("not found"):
			return nil, newErrNotFound(err.Error())
		default:
			return nil, newErrInternal("could not get thing from db: %v", err)
		}
	}

	return &thing, nil
}

func (m *Manager) getThingsFromRepo(ctx context.Context, limit *int64) ([]*models.Thing, error) {
	thingsResponse := models.ThingsListResponse{}
	thingsResponse.Things = []*models.Thing{}
	err := m.repo.ListThings(ctx, m.localLimitOrGlobalLimit(limit), &thingsResponse)
	if err != nil {
		return nil, newErrInternal("could not list things: %v", err)
	}

	return thingsResponse.Things, nil
}

func (m *Manager) getActionFromRepo(ctx context.Context, id strfmt.UUID) (*models.Action, error) {
	action := models.Action{}
	action.Schema = map[string]models.JSONObject{}
	err := m.repo.GetAction(ctx, id, &action)
	if err != nil {
		switch err {
		// TODO: Replace with structured error
		case errors.New("not found"):
			return nil, newErrNotFound(err.Error())
		default:
			return nil, newErrInternal("could not get action from db: %v", err)
		}
	}

	return &action, nil
}

func (m *Manager) getActionsFromRepo(ctx context.Context, limit *int64) ([]*models.Action, error) {
	actionsResponse := models.ActionsListResponse{}
	actionsResponse.Actions = []*models.Action{}
	err := m.repo.ListActions(ctx, m.localLimitOrGlobalLimit(limit), &actionsResponse)
	if err != nil {
		return nil, newErrInternal("could not list actions: %v", err)
	}

	return actionsResponse.Actions, nil
}

func (m *Manager) localLimitOrGlobalLimit(paramMaxResults *int64) int {
	maxResults := m.config.Config.QueryDefaults.Limit
	// Get the max results from params, if exists
	if paramMaxResults != nil {
		maxResults = *paramMaxResults
	}

	// Max results form URL, otherwise max = config.Limit.
	return int(math.Min(float64(maxResults), float64(m.config.Config.QueryDefaults.Limit)))
}

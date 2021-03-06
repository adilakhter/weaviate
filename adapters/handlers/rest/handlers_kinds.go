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

package rest

import (
	"encoding/json"

	jsonpatch "github.com/evanphx/json-patch"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/actions"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/things"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/usecases/kinds"
	"github.com/semi-technologies/weaviate/usecases/telemetry"
)

type kindHandlers struct {
	manager     *kinds.Manager
	requestsLog *telemetry.RequestsLog
}

func (h *kindHandlers) addThing(params things.WeaviateThingsCreateParams,
	principal *models.Principal) middleware.Responder {
	thing, err := h.manager.AddThing(params.HTTPRequest.Context(), params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrInvalidUserInput:
			return things.NewWeaviateThingsCreateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return things.NewWeaviateThingsCreateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalAdd)
	return things.NewWeaviateThingsCreateOK().WithPayload(thing)
}

func (h *kindHandlers) validateThing(params things.WeaviateThingsValidateParams,
	principal *models.Principal) middleware.Responder {

	err := h.manager.ValidateThing(params.HTTPRequest.Context(), params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrInvalidUserInput:
			return things.NewWeaviateThingsValidateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return things.NewWeaviateThingsValidateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalQueryMeta)
	return things.NewWeaviateThingsValidateOK()
}

func (h *kindHandlers) addAction(params actions.WeaviateActionsCreateParams,
	principal *models.Principal) middleware.Responder {
	action, err := h.manager.AddAction(params.HTTPRequest.Context(), params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrInvalidUserInput:
			return actions.NewWeaviateActionsCreateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return actions.NewWeaviateActionsCreateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalAdd)
	return actions.NewWeaviateActionsCreateOK().WithPayload(action)
}

func (h *kindHandlers) validateAction(params actions.WeaviateActionsValidateParams,
	principal *models.Principal) middleware.Responder {

	err := h.manager.ValidateAction(params.HTTPRequest.Context(), params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrInvalidUserInput:
			return actions.NewWeaviateActionsValidateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return actions.NewWeaviateActionsValidateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalQueryMeta)
	return actions.NewWeaviateActionsValidateOK()
}

func (h *kindHandlers) getThing(params things.WeaviateThingsGetParams,
	principal *models.Principal) middleware.Responder {
	thing, err := h.manager.GetThing(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound:
			return things.NewWeaviateThingsGetNotFound()
		default:
			return things.NewWeaviateThingsGetInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalQuery)
	return things.NewWeaviateThingsGetOK().WithPayload(thing)
}

func (h *kindHandlers) getAction(params actions.WeaviateActionsGetParams,
	principal *models.Principal) middleware.Responder {
	action, err := h.manager.GetAction(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound:
			return actions.NewWeaviateActionsGetNotFound()
		default:
			return actions.NewWeaviateActionsGetInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalQuery)
	return actions.NewWeaviateActionsGetOK().WithPayload(action)
}

func (h *kindHandlers) getThings(params things.WeaviateThingsListParams,
	principal *models.Principal) middleware.Responder {
	list, err := h.manager.GetThings(params.HTTPRequest.Context(), params.Limit)
	if err != nil {
		return things.NewWeaviateThingsListInternalServerError().
			WithPayload(errPayloadFromSingleErr(err))
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalQuery)
	// TODO: unify response with other verbs
	return things.NewWeaviateThingsListOK().
		WithPayload(&models.ThingsListResponse{
			Things:       list,
			TotalResults: int64(len(list)),
		})
}

func (h *kindHandlers) getActions(params actions.WeaviateActionsListParams,
	principal *models.Principal) middleware.Responder {
	list, err := h.manager.GetActions(params.HTTPRequest.Context(), params.Limit)
	if err != nil {
		return actions.NewWeaviateActionsListInternalServerError().
			WithPayload(errPayloadFromSingleErr(err))
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalQuery)
	// TODO: unify response with other verbs
	return actions.NewWeaviateActionsListOK().
		WithPayload(&models.ActionsListResponse{
			Actions:      list,
			TotalResults: int64(len(list)),
		})
}

func (h *kindHandlers) updateThing(params things.WeaviateThingsUpdateParams,
	principal *models.Principal) middleware.Responder {
	thing, err := h.manager.UpdateThing(params.HTTPRequest.Context(), params.ID, params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrInvalidUserInput:
			return things.NewWeaviateThingsUpdateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return things.NewWeaviateThingsUpdateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)
	return things.NewWeaviateThingsUpdateOK().WithPayload(thing)
}

func (h *kindHandlers) updateAction(params actions.WeaviateActionUpdateParams,
	principal *models.Principal) middleware.Responder {
	action, err := h.manager.UpdateAction(params.HTTPRequest.Context(), params.ID, params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrInvalidUserInput:
			return actions.NewWeaviateActionUpdateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return actions.NewWeaviateActionUpdateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)
	return actions.NewWeaviateActionUpdateOK().WithPayload(action)
}

func (h *kindHandlers) deleteThing(params things.WeaviateThingsDeleteParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.DeleteThing(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound:
			return things.NewWeaviateThingsDeleteNotFound()
		default:
			return things.NewWeaviateThingsDeleteInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)
	return things.NewWeaviateThingsDeleteNoContent()
}

func (h *kindHandlers) deleteAction(params actions.WeaviateActionsDeleteParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.DeleteAction(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound:
			return actions.NewWeaviateActionsDeleteNotFound()
		default:
			return actions.NewWeaviateActionsDeleteInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)
	return actions.NewWeaviateActionsDeleteNoContent()
}

// patchThing uses RFC 6902 semantics (https://tools.ietf.org/html/rfc6902) to allow
// a partial modificatiof the thing resource
//
// Internally, this means, we need to first run the Get UC, then apply the
// patch and then run the update UC
func (h *kindHandlers) patchThing(params things.WeaviateThingsPatchParams, principal *models.Principal) middleware.Responder {
	origThing, err := h.manager.GetThing(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound:
			return things.NewWeaviateThingsPatchNotFound()
		default:
			return things.NewWeaviateThingsPatchInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	patched := &models.Thing{}
	err = h.getPatchedKind(origThing, params.Body, patched)
	if err != nil {
		switch err.(type) {
		case kinds.ErrInvalidUserInput:
			return things.NewWeaviateThingsPatchUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return things.NewWeaviateThingsPatchInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	updated, err := h.manager.UpdateThing(params.HTTPRequest.Context(), params.ID, patched)
	if err != nil {
		switch err.(type) {
		case kinds.ErrInvalidUserInput:
			return things.NewWeaviateThingsUpdateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return things.NewWeaviateThingsUpdateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)

	return things.NewWeaviateThingsPatchOK().WithPayload(updated)
}

// patchAction uses RFC 6902 semantics (https://tools.ietf.org/html/rfc6902) to allow
// a partial modificatiof the action resource
//
// Internally, this means, we need to first run the Get UC, then apply the
// patch and then run the update UC
func (h *kindHandlers) patchAction(params actions.WeaviateActionsPatchParams, principal *models.Principal) middleware.Responder {
	origAction, err := h.manager.GetAction(params.HTTPRequest.Context(), params.ID)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound:
			return actions.NewWeaviateActionsPatchNotFound()
		default:
			return actions.NewWeaviateActionsPatchInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	patched := &models.Action{}
	err = h.getPatchedKind(origAction, params.Body, patched)
	if err != nil {
		switch err.(type) {
		case kinds.ErrInvalidUserInput:
			return actions.NewWeaviateActionsPatchUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return actions.NewWeaviateActionsPatchInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	updated, err := h.manager.UpdateAction(params.HTTPRequest.Context(), params.ID, patched)
	if err != nil {
		switch err.(type) {
		case kinds.ErrInvalidUserInput:
			return actions.NewWeaviateActionUpdateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return actions.NewWeaviateActionUpdateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)

	// Returns accepted so a Go routine can process in the background
	return actions.NewWeaviateActionsPatchOK().WithPayload(updated)
}

func (h *kindHandlers) getPatchedKind(orig interface{},
	patch interface{}, target interface{}) error {

	// Get PATCH params in format RFC 6902
	jsonBody, err := json.Marshal(patch)
	if err != nil {
		return kinds.ErrInternal(err)
	}

	patchObject, err := jsonpatch.DecodePatch([]byte(jsonBody))
	if err != nil {
		return kinds.ErrInvalidUserInput(err)
	}

	// Convert Kind to JSON
	origJSON, err := json.Marshal(orig)
	if err != nil {
		return kinds.ErrInternal(err)
	}

	// Apply the patch
	updatedJSON, err := patchObject.Apply(origJSON)
	if err != nil {
		return kinds.ErrInternal(err)
	}

	// Marshal back to original format
	err = json.Unmarshal([]byte(updatedJSON), &target)
	if err != nil {
		return kinds.ErrInternal(err)
	}

	return nil
}

func (h *kindHandlers) addThingReference(params things.WeaviateThingsReferencesCreateParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.AddThingReference(params.HTTPRequest.Context(), params.ID, params.PropertyName, params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound, kinds.ErrInvalidUserInput:
			return things.NewWeaviateThingsReferencesCreateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return things.NewWeaviateThingsReferencesCreateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)
	return things.NewWeaviateThingsReferencesCreateOK()
}

func (h *kindHandlers) addActionReference(params actions.WeaviateActionsReferencesCreateParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.AddActionReference(params.HTTPRequest.Context(), params.ID, params.PropertyName, params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound, kinds.ErrInvalidUserInput:
			return actions.NewWeaviateActionsReferencesCreateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return actions.NewWeaviateActionsReferencesCreateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)
	return actions.NewWeaviateActionsReferencesCreateOK()
}

func (h *kindHandlers) updateActionReferences(params actions.WeaviateActionsReferencesUpdateParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.UpdateActionReferences(params.HTTPRequest.Context(), params.ID, params.PropertyName, params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound, kinds.ErrInvalidUserInput:
			return actions.NewWeaviateActionsReferencesUpdateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return actions.NewWeaviateActionsReferencesUpdateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)
	return actions.NewWeaviateActionsReferencesUpdateOK()
}

func (h *kindHandlers) updateThingReferences(params things.WeaviateThingsReferencesUpdateParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.UpdateThingReferences(params.HTTPRequest.Context(), params.ID, params.PropertyName, params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound, kinds.ErrInvalidUserInput:
			return things.NewWeaviateThingsReferencesUpdateUnprocessableEntity().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return things.NewWeaviateThingsReferencesUpdateInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)
	return things.NewWeaviateThingsReferencesUpdateOK()
}

func (h *kindHandlers) deleteActionReference(params actions.WeaviateActionsReferencesDeleteParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.DeleteActionReference(params.HTTPRequest.Context(), params.ID, params.PropertyName, params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound, kinds.ErrInvalidUserInput:
			return actions.NewWeaviateActionsReferencesDeleteNotFound().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return actions.NewWeaviateActionsReferencesDeleteInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)
	return actions.NewWeaviateActionsReferencesDeleteNoContent()
}

func (h *kindHandlers) deleteThingReference(params things.WeaviateThingsReferencesDeleteParams,
	principal *models.Principal) middleware.Responder {
	err := h.manager.DeleteThingReference(params.HTTPRequest.Context(), params.ID, params.PropertyName, params.Body)
	if err != nil {
		switch err.(type) {
		case kinds.ErrNotFound, kinds.ErrInvalidUserInput:
			return things.NewWeaviateThingsReferencesDeleteNotFound().
				WithPayload(errPayloadFromSingleErr(err))
		default:
			return things.NewWeaviateThingsReferencesDeleteInternalServerError().
				WithPayload(errPayloadFromSingleErr(err))
		}
	}

	h.telemetryLogAsync(telemetry.TypeREST, telemetry.LocalManipulate)
	return things.NewWeaviateThingsReferencesDeleteNoContent()
}

func setupKindHandlers(api *operations.WeaviateAPI, requestsLog *telemetry.RequestsLog, manager *kinds.Manager) {
	h := &kindHandlers{manager, requestsLog}

	api.ThingsWeaviateThingsCreateHandler = things.
		WeaviateThingsCreateHandlerFunc(h.addThing)
	api.ThingsWeaviateThingsValidateHandler = things.
		WeaviateThingsValidateHandlerFunc(h.validateThing)
	api.ThingsWeaviateThingsGetHandler = things.
		WeaviateThingsGetHandlerFunc(h.getThing)
	api.ThingsWeaviateThingsDeleteHandler = things.
		WeaviateThingsDeleteHandlerFunc(h.deleteThing)
	api.ThingsWeaviateThingsListHandler = things.
		WeaviateThingsListHandlerFunc(h.getThings)
	api.ThingsWeaviateThingsUpdateHandler = things.
		WeaviateThingsUpdateHandlerFunc(h.updateThing)
	api.ThingsWeaviateThingsPatchHandler = things.
		WeaviateThingsPatchHandlerFunc(h.patchThing)
	api.ThingsWeaviateThingsReferencesCreateHandler = things.
		WeaviateThingsReferencesCreateHandlerFunc(h.addThingReference)
	api.ThingsWeaviateThingsReferencesDeleteHandler = things.
		WeaviateThingsReferencesDeleteHandlerFunc(h.deleteThingReference)
	api.ThingsWeaviateThingsReferencesUpdateHandler = things.
		WeaviateThingsReferencesUpdateHandlerFunc(h.updateThingReferences)

	api.ActionsWeaviateActionsCreateHandler = actions.
		WeaviateActionsCreateHandlerFunc(h.addAction)
	api.ActionsWeaviateActionsValidateHandler = actions.
		WeaviateActionsValidateHandlerFunc(h.validateAction)
	api.ActionsWeaviateActionsGetHandler = actions.
		WeaviateActionsGetHandlerFunc(h.getAction)
	api.ActionsWeaviateActionsDeleteHandler = actions.
		WeaviateActionsDeleteHandlerFunc(h.deleteAction)
	api.ActionsWeaviateActionsListHandler = actions.
		WeaviateActionsListHandlerFunc(h.getActions)
	api.ActionsWeaviateActionUpdateHandler = actions.
		WeaviateActionUpdateHandlerFunc(h.updateAction)
	api.ActionsWeaviateActionsPatchHandler = actions.
		WeaviateActionsPatchHandlerFunc(h.patchAction)
	api.ActionsWeaviateActionsReferencesCreateHandler = actions.
		WeaviateActionsReferencesCreateHandlerFunc(h.addActionReference)
	api.ActionsWeaviateActionsReferencesDeleteHandler = actions.
		WeaviateActionsReferencesDeleteHandlerFunc(h.deleteActionReference)
	api.ActionsWeaviateActionsReferencesUpdateHandler = actions.
		WeaviateActionsReferencesUpdateHandlerFunc(h.updateActionReferences)

}

func (h *kindHandlers) telemetryLogAsync(requestType, identifier string) {
	go func() {
		h.requestsLog.Register(requestType, identifier)
	}()
}

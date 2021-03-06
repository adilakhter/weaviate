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
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/meta"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/p2_p"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/semi-technologies/weaviate/usecases/network"
	"github.com/semi-technologies/weaviate/usecases/network/common/peers"
	"github.com/semi-technologies/weaviate/usecases/telemetry"
)

type schemaManager interface {
	GetSchema() schema.Schema
}

func setupMiscHandlers(api *operations.WeaviateAPI, requestsLog *telemetry.RequestsLog,
	serverConfig *config.WeaviateConfig, network network.Network, schemaManager schemaManager) {
	api.MetaWeaviateMetaGetHandler = meta.WeaviateMetaGetHandlerFunc(func(params meta.WeaviateMetaGetParams, principal *models.Principal) middleware.Responder {
		databaseSchema := schema.HackFromDatabaseSchema(schemaManager.GetSchema())
		// Create response object
		metaResponse := &models.Meta{}

		// Set the response object's values
		metaResponse.Hostname = serverConfig.GetHostAddress()
		metaResponse.ActionsSchema = databaseSchema.ActionSchema.Schema
		metaResponse.ThingsSchema = databaseSchema.ThingSchema.Schema

		// Register the request
		go func() {
			requestsLog.Register(telemetry.TypeREST, telemetry.LocalQueryMeta)
		}()
		return meta.NewWeaviateMetaGetOK().WithPayload(metaResponse)
	})

	api.P2PWeaviateP2pGenesisUpdateHandler = p2_p.WeaviateP2pGenesisUpdateHandlerFunc(func(params p2_p.WeaviateP2pGenesisUpdateParams) middleware.Responder {
		newPeers := make([]peers.Peer, 0)

		for _, genesisPeer := range params.Peers {
			peer := peers.Peer{
				ID:         genesisPeer.ID,
				Name:       genesisPeer.Name,
				URI:        genesisPeer.URI,
				SchemaHash: genesisPeer.SchemaHash,
			}

			newPeers = append(newPeers, peer)
		}

		err := network.UpdatePeers(newPeers)

		if err == nil {
			// Register the request
			go func() {
				requestsLog.Register(telemetry.TypeREST, telemetry.NetworkQueryMeta)
			}()
			return p2_p.NewWeaviateP2pGenesisUpdateOK()
		}
		return p2_p.NewWeaviateP2pGenesisUpdateInternalServerError()
	})

	api.P2PWeaviateP2pHealthHandler = p2_p.WeaviateP2pHealthHandlerFunc(func(params p2_p.WeaviateP2pHealthParams) middleware.Responder {
		// For now, always just return success.
		return middleware.NotImplemented("operation P2PWeaviateP2pHealth has not yet been implemented")
	})
}

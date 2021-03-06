package server

import (
	"net/http"

	"github.com/containers/podman/v2/pkg/api/handlers/compat"
	"github.com/containers/podman/v2/pkg/api/handlers/libpod"
	"github.com/gorilla/mux"
)

func (s *APIServer) registerSecretHandlers(r *mux.Router) error {
	// swagger:operation POST /libpod/secrets/create libpod libpodCreateSecret
	// ---
	// tags:
	//  - secrets
	// summary: Create a secret
	// parameters:
	//   - in: query
	//     name: name
	//     type: string
	//     description: User-defined name of the secret.
	//     required: true
	//   - in: query
	//     name: driver
	//     type: string
	//     description: Secret driver
	//     default: "file"
	//   - in: body
	//     name: request
	//     description: Secret
	//     schema:
	//       type: string
	// produces:
	// - application/json
	// responses:
	//   '201':
	//     $ref: "#/responses/SecretCreateResponse"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/libpod/secrets/create"), s.APIHandler(libpod.CreateSecret)).Methods(http.MethodPost)
	// swagger:operation GET /libpod/secrets/json libpod libpodListSecret
	// ---
	// tags:
	//  - secrets
	// summary: List secrets
	// description: Returns a list of secrets
	// produces:
	// - application/json
	// parameters:
	// responses:
	//   '200':
	//     "$ref": "#/responses/SecretListResponse"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/libpod/secrets/json"), s.APIHandler(compat.ListSecrets)).Methods(http.MethodGet)
	// swagger:operation GET /libpod/secrets/{name}/json libpod libpodInspectSecret
	// ---
	// tags:
	//  - secrets
	// summary: Inspect secret
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the secret
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     "$ref": "#/responses/SecretInspectResponse"
	//   '404':
	//     "$ref": "#/responses/NoSuchSecret"
	//   '500':
	//     "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/libpod/secrets/{name}/json"), s.APIHandler(compat.InspectSecret)).Methods(http.MethodGet)
	// swagger:operation DELETE /libpod/secrets/{name} libpod libpodRemoveSecret
	// ---
	// tags:
	//  - secrets
	// summary: Remove secret
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the secret
	//  - in: query
	//    name: all
	//    type: boolean
	//    description: Remove all secrets
	//    default: false
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '404':
	//     "$ref": "#/responses/NoSuchSecret"
	//   '500':
	//     "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/libpod/secrets/{name}"), s.APIHandler(compat.RemoveSecret)).Methods(http.MethodDelete)

	/*
	 * Docker compatibility endpoints
	 */
	// swagger:operation GET /secrets compat ListSecret
	// ---
	// tags:
	//  - secrets (compat)
	// summary: List secrets
	// description: Returns a list of secrets
	// produces:
	// - application/json
	// parameters:
	// responses:
	//   '200':
	//     "$ref": "#/responses/SecretListResponse"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/secrets"), s.APIHandler(compat.ListSecrets)).Methods(http.MethodGet)
	r.Handle("/secrets", s.APIHandler(compat.ListSecrets)).Methods(http.MethodGet)
	// swagger:operation POST /secrets/create compat CreateSecret
	// ---
	// tags:
	//  - secrets (compat)
	// summary: Create a secret
	// parameters:
	//  - in: body
	//    name: create
	//    description: |
	//      attributes for creating a secret
	//    schema:
	//      $ref: "#/definitions/SecretCreate"
	// produces:
	// - application/json
	// responses:
	//   '201':
	//     $ref: "#/responses/SecretCreateResponse"
	//   '409':
	//     "$ref": "#/responses/SecretInUse"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/secrets/create"), s.APIHandler(compat.CreateSecret)).Methods(http.MethodPost)
	r.Handle("/secrets/create", s.APIHandler(compat.CreateSecret)).Methods(http.MethodPost)
	// swagger:operation GET /secrets/{name} compat InspectSecret
	// ---
	// tags:
	//  - secrets (compat)
	// summary: Inspect secret
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the secret
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     "$ref": "#/responses/SecretInspectResponse"
	//   '404':
	//     "$ref": "#/responses/NoSuchSecret"
	//   '500':
	//     "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/secrets/{name}"), s.APIHandler(compat.InspectSecret)).Methods(http.MethodGet)
	r.Handle("/secrets/{name}", s.APIHandler(compat.InspectSecret)).Methods(http.MethodGet)
	// swagger:operation DELETE /secrets/{name} compat RemoveSecret
	// ---
	// tags:
	//  - secrets (compat)
	// summary: Remove secret
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the secret
	// produces:
	// - application/json
	// responses:
	//   '204':
	//     description: no error
	//   '404':
	//     "$ref": "#/responses/NoSuchSecret"
	//   '500':
	//     "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/secrets/{name}"), s.APIHandler(compat.RemoveSecret)).Methods(http.MethodDelete)
	r.Handle("/secret/{name}", s.APIHandler(compat.RemoveSecret)).Methods(http.MethodDelete)

	r.Handle(VersionedPath("/secrets/{name}/update"), s.APIHandler(compat.UpdateSecret)).Methods(http.MethodPost)
	r.Handle("/secrets/{name}/update", s.APIHandler(compat.UpdateSecret)).Methods(http.MethodPost)
	return nil
}

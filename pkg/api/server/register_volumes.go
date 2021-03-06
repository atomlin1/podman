package server

import (
	"net/http"

	"github.com/containers/podman/v3/pkg/api/handlers/compat"
	"github.com/containers/podman/v3/pkg/api/handlers/libpod"
	"github.com/gorilla/mux"
)

func (s *APIServer) registerVolumeHandlers(r *mux.Router) error {
	// swagger:operation POST /libpod/volumes/create libpod libpodCreateVolume
	// ---
	// tags:
	//  - volumes
	// summary: Create a volume
	// parameters:
	//  - in: body
	//    name: create
	//    description: attributes for creating a container
	//    schema:
	//      $ref: "#/definitions/VolumeCreate"
	// produces:
	// - application/json
	// responses:
	//   '201':
	//     $ref: "#/responses/VolumeCreateResponse"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/libpod/volumes/create"), s.APIHandler(libpod.CreateVolume)).Methods(http.MethodPost)
	// swagger:operation GET /libpod/volumes/{name}/exists libpod libpodExistsVolume
	// ---
	// tags:
	//  - volumes
	// summary: Volume exists
	// description: Check if a volume exists
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name of the volume
	// produces:
	// - application/json
	// responses:
	//   204:
	//     description: volume exists
	//   404:
	//     $ref: '#/responses/NoSuchVolume'
	//   500:
	//     $ref: '#/responses/InternalError'
	r.Handle(VersionedPath("/libpod/volumes/{name}/exists"), s.APIHandler(libpod.ExistsVolume)).Methods(http.MethodGet)
	// swagger:operation GET /libpod/volumes/json libpod libpodListVolumes
	// ---
	// tags:
	//  - volumes
	// summary: List volumes
	// description: Returns a list of volumes
	// produces:
	// - application/json
	// parameters:
	//  - in: query
	//    name: filters
	//    type: string
	//    description: |
	//      JSON encoded value of the filters (a map[string][]string) to process on the volumes list. Available filters:
	//        - driver=<volume-driver-name> Matches volumes based on their driver.
	//        - label=<key> or label=<key>:<value> Matches volumes based on the presence of a label alone or a label and a value.
	//        - name=<volume-name> Matches all of volume name.
	//        - opt=<driver-option> Matches a storage driver options
	// responses:
	//   '200':
	//     "$ref": "#/responses/VolumeList"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/libpod/volumes/json"), s.APIHandler(libpod.ListVolumes)).Methods(http.MethodGet)
	// swagger:operation POST /libpod/volumes/prune libpod libpodPruneVolumes
	// ---
	// tags:
	//  - volumes
	// summary: Prune volumes
	// produces:
	// - application/json
	// responses:
	//   '200':
	//      "$ref": "#/responses/VolumePruneResponse"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/libpod/volumes/prune"), s.APIHandler(libpod.PruneVolumes)).Methods(http.MethodPost)
	// swagger:operation GET /libpod/volumes/{name}/json libpod libpodInspectVolume
	// ---
	// tags:
	//  - volumes
	// summary: Inspect volume
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the volume
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     "$ref": "#/responses/VolumeCreateResponse"
	//   '404':
	//     "$ref": "#/responses/NoSuchVolume"
	//   '500':
	//     "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/libpod/volumes/{name}/json"), s.APIHandler(libpod.InspectVolume)).Methods(http.MethodGet)
	// swagger:operation DELETE /libpod/volumes/{name} libpod libpodRemoveVolume
	// ---
	// tags:
	//  - volumes
	// summary: Remove volume
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the volume
	//  - in: query
	//    name: force
	//    type: boolean
	//    description: force removal
	// produces:
	// - application/json
	// responses:
	//   204:
	//     description: no error
	//   404:
	//     $ref: "#/responses/NoSuchVolume"
	//   409:
	//     description: Volume is in use and cannot be removed
	//   500:
	//     $ref: "#/responses/InternalError"
	r.Handle(VersionedPath("/libpod/volumes/{name}"), s.APIHandler(libpod.RemoveVolume)).Methods(http.MethodDelete)

	/*
	 * Docker compatibility endpoints
	 */

	// swagger:operation GET /volumes compat listVolumes
	// ---
	// tags:
	//  - volumes (compat)
	// summary: List volumes
	// description: Returns a list of volume
	// produces:
	// - application/json
	// parameters:
	//  - in: query
	//    name: filters
	//    type: string
	//    description: |
	//      JSON encoded value of the filters (a map[string][]string) to process on the volumes list. Available filters:
	//        - driver=<volume-driver-name> Matches volumes based on their driver.
	//        - label=<key> or label=<key>:<value> Matches volumes based on the presence of a label alone or a label and a value.
	//        - name=<volume-name> Matches all of volume name.
	//
	//      Note:
	//        The boolean `dangling` filter is not yet implemented for this endpoint.
	// responses:
	//   '200':
	//     "$ref": "#/responses/VolumeListResponse"
	//   '500':
	//     "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/volumes"), s.APIHandler(compat.ListVolumes)).Methods(http.MethodGet)
	r.Handle("/volumes", s.APIHandler(compat.ListVolumes)).Methods(http.MethodGet)

	// swagger:operation POST /volumes/create compat createVolume
	// ---
	// tags:
	//  - volumes (compat)
	// summary: Create a volume
	// parameters:
	//  - in: body
	//    name: create
	//    description: |
	//      attributes for creating a container.
	//      Note: If a volume by the same name exists, a 201 response with that volume's information will be generated.
	//    schema:
	//      $ref: "#/definitions/DockerVolumeCreate"
	// produces:
	// - application/json
	// responses:
	//   '201':
	//     "$ref": "#/responses/DockerVolumeInfoResponse"
	//   '500':
	//     "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/volumes/create"), s.APIHandler(compat.CreateVolume)).Methods(http.MethodPost)
	r.Handle("/volumes/create", s.APIHandler(compat.CreateVolume)).Methods(http.MethodPost)

	// swagger:operation GET /volumes/{name} compat inspectVolume
	// ---
	// tags:
	//  - volumes (compat)
	// summary: Inspect volume
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the volume
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     "$ref": "#/responses/DockerVolumeInfoResponse"
	//   '404':
	//     "$ref": "#/responses/NoSuchVolume"
	//   '500':
	//     "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/volumes/{name}"), s.APIHandler(compat.InspectVolume)).Methods(http.MethodGet)
	r.Handle("/volumes/{name}", s.APIHandler(compat.InspectVolume)).Methods(http.MethodGet)

	// swagger:operation DELETE /volumes/{name} compat removeVolume
	// ---
	// tags:
	//  - volumes (compat)
	// summary: Remove volume
	// parameters:
	//  - in: path
	//    name: name
	//    type: string
	//    required: true
	//    description: the name or ID of the volume
	//  - in: query
	//    name: force
	//    type: boolean
	//    description: |
	//      Force removal of the volume. This actually only causes errors due
	//      to the names volume not being found to be suppressed, which is the
	//      behaviour Docker implements.
	// produces:
	// - application/json
	// responses:
	//   204:
	//     description: no error
	//   404:
	//     "$ref": "#/responses/NoSuchVolume"
	//   409:
	//     description: Volume is in use and cannot be removed
	//   500:
	//     "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/volumes/{name}"), s.APIHandler(compat.RemoveVolume)).Methods(http.MethodDelete)
	r.Handle("/volumes/{name}", s.APIHandler(compat.RemoveVolume)).Methods(http.MethodDelete)

	// swagger:operation POST /volumes/prune compat pruneVolumes
	// ---
	// tags:
	//  - volumes (compat)
	// summary: Prune volumes
	// produces:
	// - application/json
	// parameters:
	//  - in: query
	//    name: filters
	//    type: string
	//    description: |
	//      JSON encoded value of filters (a map[string][]string) to match volumes against before pruning.
	//
	//      Note: No filters are currently supported and any filters specified will cause an error response.
	// responses:
	//   '200':
	//      "$ref": "#/responses/DockerVolumePruneResponse"
	//   '500':
	//      "$ref": "#/responses/InternalError"
	r.Handle(VersionedPath("/volumes/prune"), s.APIHandler(compat.PruneVolumes)).Methods(http.MethodPost)
	r.Handle("/volumes/prune", s.APIHandler(compat.PruneVolumes)).Methods(http.MethodPost)

	return nil
}

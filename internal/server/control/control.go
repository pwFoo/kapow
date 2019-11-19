/*
 * Copyright 2019 Banco Bilbao Vizcaya Argentaria, S.A.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package control

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/BBVA/kapow/internal/server/model"
	"github.com/BBVA/kapow/internal/server/srverrors"
	"github.com/BBVA/kapow/internal/server/user"
)

// configRouter Populates the server mux with all the supported routes. The
// server exposes list, get, delete and add route endpoints.
func configRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/routes/{id}", removeRoute).
		Methods(http.MethodDelete)
	r.HandleFunc("/routes/{id}", getRoute).
		Methods(http.MethodGet)
	r.HandleFunc("/routes", listRoutes).
		Methods(http.MethodGet)
	r.HandleFunc("/routes", addRoute).
		Methods(http.MethodPost)
	return r
}

// funcRemove Method used to ask the route model module to delete a route
var funcRemove func(id string) error = user.Routes.Delete

// removeRoute Handler that removes the requested route.  If it doesn't exist,
// returns 404 and an error entity
func removeRoute(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	if err := funcRemove(id); err != nil {
		srverrors.WriteErrorResponse(http.StatusNotFound, "Route Not Found", res)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}

// funcList Method used to ask the route model module for the list of routes
var funcList func() []model.Route = user.Routes.List

// listRoutes Handler that retrieves a list of the existing routes. An empty
// list is returned when no routes exist
func listRoutes(res http.ResponseWriter, req *http.Request) {

	list := funcList()

	listBytes, _ := json.Marshal(list)
	res.Header().Set("Content-Type", "application/json")
	_, _ = res.Write(listBytes)
}

// funcAdd Method used to ask the route model module to append a new route
var funcAdd func(model.Route) model.Route = user.Routes.Append

// idGenerator UUID generator for new routes
var idGenerator = uuid.NewUUID

// pathValidator Validates that a path complies with the gorilla mux
// requirements
var pathValidator func(string) error = func(path string) error {
	return mux.NewRouter().NewRoute().BuildOnly().Path(path).GetError()
}

// addRoute Handler that adds a new route. Makes all parameter validation and
// creates the a new is for the route
func addRoute(res http.ResponseWriter, req *http.Request) {
	var route model.Route

	payload, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(payload, &route)
	if err != nil {
		srverrors.WriteErrorResponse(http.StatusBadRequest, "Malformed JSON", res)
		return
	}

	if route.Method == "" {
		srverrors.WriteErrorResponse(http.StatusUnprocessableEntity, "Invalid Route", res)
		return
	}

	if route.Pattern == "" {
		srverrors.WriteErrorResponse(http.StatusUnprocessableEntity, "Invalid Route", res)
		return
	}

	err = pathValidator(route.Pattern)
	if err != nil {
		srverrors.WriteErrorResponse(http.StatusUnprocessableEntity, "Invalid Route", res)
		return
	}

	id, err := idGenerator()
	if err != nil {
		srverrors.WriteErrorResponse(http.StatusInternalServerError, "Internal Server Error", res)
		return
	}

	route.ID = id.String()

	created := funcAdd(route)
	createdBytes, _ := json.Marshal(created)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	_, _ = res.Write(createdBytes)
}

// funcGet Method used to ask the route model module for the details of a route
var funcGet func(string) (model.Route, error) = user.Routes.Get

// getRoute Handler that retrieves the details of a route. If the route doesn't
// exists returns 404 and an error entity
func getRoute(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if r, err := funcGet(id); err != nil {
		srverrors.WriteErrorResponse(http.StatusNotFound, "Route Not Found", res)
	} else {
		res.Header().Set("Content-Type", "application/json")
		rBytes, _ := json.Marshal(r)
		_, _ = res.Write(rBytes)
	}
}

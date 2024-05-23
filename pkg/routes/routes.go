/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type UserData struct {
	Name	string	`json:"name"`
	Message	string	`json:"msg"`
}

type Response struct {
	Status 	string	`json:"status"`
}

func handleHelloEndpoint (rw http.ResponseWriter, r *http.Request) {

	fmt.Println("Query received in /hello")
	
	var userData UserData
	err := json.NewDecoder(r.Body).Decode(&userData)

	rw.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("Error in request")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(Response{Status: "error found: " + err.Error()})
	} else {
		fmt.Println(fmt.Sprintf("Received: {\"name\": \"%s\", \"msg\": \"%s\"}\n", userData.Name, userData.Message))

		if userData.Name != "" && userData.Message != "" {
			fmt.Println("Correct case")
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(Response{Status: "correct"})
		} else {
			fmt.Println("Not found case")
			rw.WriteHeader(http.StatusNotFound)
			json.NewEncoder(rw).Encode(Response{Status: "not found"})
		}
	}
}

func SetChiRouter() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	return r
}

func SetEndpoints(mx *chi.Mux) {

	// Hello endpoint
	mx.Get("/hello", handleHelloEndpoint)
}

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

package main

import (
	"log/slog"
	"net/http"

	"github.com/rampherher/dummy-webserver/pkg/routes"
)

func main() {

	// Create the Chi Router
	r := routes.SetChiRouter()

	// Set endpoints beforehand
	slog.Info("Set endpoints on webserver")
	routes.SetEndpoints(r)

	// Then launch the webserver
	slog.Info("Start webserver")
	http.ListenAndServe(":8080", r)
}

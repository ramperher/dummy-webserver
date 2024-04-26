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

// import the package we need to use
import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func setEndpoints() {

	log.Println("Configure webserver")

	// Hello endpoint
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Query received in /hello")
		rw.WriteHeader(200)
		rw.Write([]byte("ok"))
	})
}

func startWebServer() {

	log.Println("Start webserver")

	// Create server at localhost:8080 and using tcp
	listener, err := net.Listen("tcp", ":8080")

	// If recieving an error, just record it and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP connection for the listener of the server
	http.Serve(listener, nil)
}

func main() {

	// Set endpoints beforehand
	setEndpoints()

	// Then launch the webserver
	startWebServer()
}

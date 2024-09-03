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
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestEndpoints(t *testing.T) {
	testCases := []struct {
		name     string
		body     []byte
		result   int
	}{
		{
			name:     "empty input",
			body:     []byte{},
			result:   400,
		},
		{
			name:     "wrong name",
			body:     []byte(`{"nam": "bob"}`),
			result:   404,
		},
		{
			name:     "only correct msg",
			body:     []byte(`{"msg":"hi all"}`),
			result:   404,
		},
		{
			name:     "all correct",
			body:     []byte(`{"name":"bob","msg":"hi all"}`),
			result:   200,
		},
	}
	for _, tc := range testCases {

		jsonBody := tc.body
		bodyReader := bytes.NewReader(jsonBody)

		requestURL := "http://localhost:8080/hello"
		req, err := http.NewRequest(http.MethodGet, requestURL, bodyReader)
	
		if err != nil {
			   slog.Error(fmt.Sprintf("client: could not create request: %s\n", err))
			   os.Exit(1)
		}
		req.Header.Set("Content-Type", "application/json")
	  
		client := http.Client{
		   Timeout: 30 * time.Second,
		}
	  
		res, err := client.Do(req)
		if err != nil {
			slog.Error(fmt.Sprintf("client: error making http request: %s\n", err))
			os.Exit(1)
		} else {
			if (tc.result != res.StatusCode) {
				t.Errorf("test failed")
			} else {
				slog.Info(fmt.Sprintf("correct result: %d", tc.result))
			}
		}
	}
}

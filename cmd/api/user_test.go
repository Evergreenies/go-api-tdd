package main

import (
	"net/http"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		name         string
		expectedCode int
		body         string
	}{
		{
			name:         "OK",
			expectedCode: http.StatusOK,
			body:         `{ "name": "Suyo Shimpi", "email": "suyogshimpi@mail.com", "password": "password" }`,
		}, {
			name:         "Bad JSON",
			expectedCode: http.StatusBadRequest,
			body:         `{"name": "}`,
		}, {
			name:         "Validation Error",
			expectedCode: http.StatusBadRequest,
			body:         `{ "name": "", "email": "", "password": "" }`,
		},
	}
	serv := newServer(testStore, nil)
	ts := newTestServer(serv.routes())

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			resp, err := ts.Client().Post(
				ts.URL+"/api/v1/users/create",
				"application/json",
				strings.NewReader(test.body),
			)
			if err != nil {
				t.Fatal(err)
			}

			if test.expectedCode != resp.StatusCode {
				t.Errorf("expected status code %d; got %d\n", test.expectedCode, resp.StatusCode)
			}
		})
	}
}

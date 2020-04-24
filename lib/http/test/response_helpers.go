package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetResponseBody(t *testing.T, response *http.Response) string {
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	return string(b)
}

func AssertJSONMatches(t *testing.T, response *http.Response, expected interface{}) {
	body := GetResponseBody(t, response)

	// Convert our expected response to JSON
	b, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}
	expectedBody := string(b)

	assert.Equal(t, expectedBody, body)
}

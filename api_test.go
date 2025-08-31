package main

import (
	"errors"

	"reflect"

	"net/http"
	"testing"

	"github.com/TaitA2/learn-cicd-starter/internal/auth"
)

type response struct {
	apiKey string
	err    error
}

func TestGetApiKey(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected response
	}{
		"empty":           {input: "", expected: response{apiKey: "", err: errors.New("no authorization header included")}},
		"too few splits":  {input: "ApiKey", expected: response{apiKey: "", err: errors.New("malformed authorization header")}},
		"wrong split[0]":  {input: "NotApiKey ", expected: response{apiKey: "", err: errors.New("malformed authorization header")}},
		"too many splits": {input: "Not Api Key", expected: response{apiKey: "", err: errors.New("malformed authorization header")}},
		"working":         {input: "ApiKey 1245", expected: response{apiKey: "12345", err: nil}},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}
			headers.Add("Authorization", c.input)
			apiKey, err := auth.GetAPIKey(headers)
			if !reflect.DeepEqual(err, c.expected.err) {
				prettyEx, prettyErr := prettify(c.expected.err.Error(), err.Error())
				t.Fatalf("\n\033[32mExpected:\t\033[37m %v, \n\033[31mGot:\t %v", prettyEx, prettyErr)
			}
			if apiKey != c.expected.apiKey {
				prettyEx, prettyApi := prettify(c.expected.apiKey, apiKey)
				t.Fatalf("\n\033[32mexpected:\t\033[37m %v, \n\033[31mGot:\t %v", prettyEx, prettyApi)
			}
		})
	}
}

func prettify(ex, got string) (ex1, got1 string) {
	var i = 0
	for i < len(got) && i < len(ex) && got[i] == ex[i] {
		i++
	}
	if i < len(got) {
		got1 = "\033[32m" + string(got[:i]) + "\033[31m" + string(got[i:]) + "\033[37m"
		ex1 = "\033[32m" + string(ex[:i]) + "\033[33m" + string(ex[i:]) + "\033[37m"
	}

	return

}

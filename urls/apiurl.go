package urls

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/dmolesUC/gliq/options"
	"github.com/dmolesUC/gliq/util"
)

// ------------------------------------------------------------
// Exported

const ApiBaseUrlStr = "https://gitlab.com/api/v4"

type ApiUrl struct {
	url.URL
}

func (u *ApiUrl) Clone() *ApiUrl {
	u2 := *u
	if u.User != nil {
		user2 := *u.User
		u2.User = &user2
	}
	return &u2
}

func (u *ApiUrl) ReadInto(v any) {
	resp := u.get()
	defer util.QuietlyClose(resp.Body)

	body, err := io.ReadAll(resp.Body)
	util.QuietlyHandle(err)

	bodyReader := bytes.NewReader(body)

	err = json.NewDecoder(bodyReader).Decode(v)
	if err != nil {
		log.Printf("verbose error info: %#v", err)
		log.Printf("body was: %#v\n", string(body[:]))
		util.QuietlyHandle(err) // TODO: something smarter
	}
}

func (u *ApiUrl) WithParams() *ApiUrl {
	u2 := u.Clone()
	u2.RawQuery = queryFromOptions()
	return u2
}

func (u *ApiUrl) JoinPath(elem ...string) *ApiUrl {
	u2 := u.Clone()
	u2.URL = *u.URL.JoinPath(elem...)
	return u2
}

// ------------------------------------------------------------
// Exported functions

// ------------------------------------------------------------
// Package-local

var apiBaseUrl *ApiUrl

func (u *ApiUrl) toRequest() *http.Request {
	log.Printf("GET %v\n", u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	util.QuietlyHandle(err)
	return req
}

func (u *ApiUrl) get() *http.Response {
	req := u.toRequest()
	if token := options.AccessToken(); len(token) > 0 {
		// log.Printf("Adding auth token %v\n", token)
		req.Header.Add("Authorization", "Bearer "+token)
	}
	return doRequest(req)
}

// ------------------------------------------------------------
// Initializer

func init() {
	u, _ := url.Parse(ApiBaseUrlStr)
	apiBaseUrl = &ApiUrl{*u}
}

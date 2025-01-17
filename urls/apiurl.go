package urls

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/dmolesUC/gliq/config"
	"github.com/dmolesUC/gliq/params"
	"github.com/dmolesUC/gliq/util"
)

// ------------------------------------------------------------
// Constants

const ApiBaseUrlStr = "https://gitlab.com/api/v4"

// ------------------------------------------------------------
// ApiUrl

type ApiUrl struct {
	url.URL
}

// ------------------------------
// Public methods

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
	u2.RawQuery = params.ToRawQuery()
	return u2
}

func (u *ApiUrl) JoinPath(elem ...string) *ApiUrl {
	u2 := u.Clone()
	u2.URL = *u.URL.JoinPath(elem...)
	return u2
}

// ------------------------------
// Private methods

func (u *ApiUrl) toRequest() *http.Request {
	log.Printf("GET %v\n", u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	util.QuietlyHandle(err)
	return req
}

func (u *ApiUrl) get() *http.Response {
	req := u.toRequest()
	if token := config.Token; len(token) > 0 {
		// log.Printf("Adding auth token %v\n", token)
		req.Header.Add("Authorization", "Bearer "+token)
	}
	return doRequest(req)
}

// ------------------------------------------------------------
// Exported functions

func ReadAs[T any](u *ApiUrl) T {
	var v T
	u.ReadInto(&v)
	return v
}

// ------------------------------------------------------------
// Private functions

func doRequest(req *http.Request) *http.Response {
	client := &http.Client{}
	resp, err := client.Do(req)
	util.QuietlyHandle(err)
	return resp
}

// ------------------------------------------------------------
// Private state

var apiBaseUrl *ApiUrl

// ------------------------------------------------------------
// Initializer

func init() {
	u, _ := url.Parse(ApiBaseUrlStr)
	apiBaseUrl = &ApiUrl{*u}
}

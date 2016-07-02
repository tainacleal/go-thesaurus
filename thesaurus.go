package thesaurus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	restAPI  = "http://words.bighugelabs.com/api/2"
	errNoKey = "Key not set"
)

//Thesa holds the API Key for word lookup
type Thesa struct {
	Key string
}

//Response is the type for a Thesaurus response
type Response struct {
	Word   string `json:"-"`
	Format string `json:"-"`
	Code   int    `json:"-"`
	Status string `json:"-"`
	Noun   *Noun  `json:"noun"`
	Verb   *Verb  `json:"verb"`
}

//Noun is the type that holds the definitions for a noun word
type Noun struct {
	Syn []string `json:"syn"`
	Ant []string `json:"ant"`
	Rel []string `json:"rel"`
	Sim []string `json:"sim"`
	Usr []string `json:"usr"`
}

//Verb is the type that holds the definitions for a verb word
type Verb struct {
	Syn []string `json:"syn"`
	Ant []string `json:"ant"`
	Rel []string `json:"rel"`
	Sim []string `json:"sim"`
	Usr []string `json:"usr"`
}

//NewResponse is a constructor for a Response type
func NewResponse(word string) *Response {
	return &Response{Word: word}
}

//Configure sets the API Key
func Configure(key string) *Thesa {
	return &Thesa{Key: key}
}

//LookUp receives a word for look up and returns a thesaurus.Response object
func (t *Thesa) LookUp(word string) (*Response, error) {
	if t.Key != "" {
		resp, err := http.Get(fmt.Sprintf("%s/%s/%s/json", restAPI, t.Key, word))
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		r := NewResponse(word)
		r.Code = resp.StatusCode
		r.Status = resp.Status
		err = json.NewDecoder(resp.Body).Decode(&r)
		return r, err
	}
	return nil, errors.New(errNoKey)
}

//LookUpResponse receives a word for look up and the response format needed. Eg. xml, json
func (t *Thesa) LookUpResponse(word string, respType string) ([]byte, int, error) {
	if t.Key != "" {
		resp, err := http.Get(fmt.Sprintf("%s/%s/%s/%s", restAPI, t.Key, word, respType))
		if err != nil {
			return nil, 0, err
		}
		defer resp.Body.Close()
		code := resp.StatusCode
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, 0, err
		}
		return response, code, err
	}
	return nil, 0, errors.New(errNoKey)
}

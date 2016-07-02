package thesaurus

import (
	"io"
	"net/http"
	"testing"

	"github.com/tainacleal/go-thesaurus/config"
)

var TestThesa = Configure(config.Key)

func TestLookUp(t *testing.T) {
	resp, err := TestThesa.LookUp("go")
	if err != nil {
		t.Error(err)
	}

	if resp.Code != http.StatusOK {
		t.Errorf("Expected code %d, but got %d", http.StatusOK, resp.Code)
	}
}

func TestLookUpNoKey(t *testing.T) {
	thesa := &Thesa{}
	_, err := thesa.LookUp("go")
	if err.Error() != errNoKey {
		t.Fatalf("Expecting error %s, but got %s", errNoKey, err.Error())
	}
}

func TestLookUpResponse(t *testing.T) {
	_, code, err := TestThesa.LookUpResponse("programming", "xml")
	if err != nil {
		t.Errorf("Expecting err to be nil, but got %s", err)
	}
	if code != http.StatusOK {
		t.Errorf("Expecting code to be %d, but got %d", http.StatusOK, code)
	}
}

func TestWrongKey(t *testing.T) {
	thesa := Configure("wrongtestkey")
	resp, err := thesa.LookUp("wrong")
	if err != io.EOF {
		t.Errorf("Expecting err to be nil, but got %s", err.Error())
	}
	if resp.Code != http.StatusInternalServerError {
		t.Errorf("Expecting code %d, but got %d", http.StatusInternalServerError, resp.Code)
	}
}

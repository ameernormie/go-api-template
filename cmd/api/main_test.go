package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/ameernormie/go-api-template/internal/routes"
	"github.com/ameernormie/go-api-template/internal/services/encryptionservice"
	"github.com/ameernormie/go-api-template/internal/testutil"
)

func TestCpsRoute(t *testing.T) {
	testutil.InitDb()
	currentTime := time.Now()
	parsedTime := currentTime.UTC().Format(routes.INPUTFORMAT)
	token, err := encryptionservice.Encrypt(parsedTime)

	if err != nil {
		t.Errorf("Error while encryption: %d", err)
	}
	encodedURL := url.QueryEscape(token)
	router := routes.GetRouter()

	w := httptest.NewRecorder()
	getTestURL := fmt.Sprintf("/api/ping?token=%s", encodedURL)
	req, _ := http.NewRequest("GET", getTestURL, nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("response code is not OK: %d", w.Code)
	}

	if w.Body.String() != "{\"message\":\"pong\"}" {
		t.Errorf("response body is not correct: %s", w.Body.String())
	}
}

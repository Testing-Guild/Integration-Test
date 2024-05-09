package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type Application struct {
	router *mux.Router
}

func (a *Application) Post(path string, handler http.HandlerFunc) {
	a.router.HandleFunc(path, handler)
}

func (a *Application) Listen(port int) (net.Listener, error) {
	return nil, fmt.Errorf("Listen not implemented (testing purposes)")
}

func newApp() *Application {
	return &Application{router: mux.NewRouter()}
}

func sayNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"name": name})
}

func TestSayName(t *testing.T) {
	app := newApp()
	app.Post("/say-name/:name", sayNameHandler)
	recorder := httptest.NewRecorder()
	requestBody := map[string]string{"name": "mahdi"}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/say-name/mahdi", bytes.NewReader(requestBodyBytes))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	app.router.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", recorder.Code)
	}
	var responseBody map[string]string
	if err := json.NewDecoder(recorder.Body).Decode(&responseBody); err != nil {
		t.Fatal(err)
	}

	if responseBody["name"] != "mahdi" {
		t.Errorf("Expected response body to contain name: 'mahdi', got: %v", responseBody)
	}
}

func main() {
	testing.RunTests("", func(t *testing.T) {
		TestSayName(t)
	})
}

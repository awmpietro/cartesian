package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoadData(t *testing.T) {
	_, err := ioutil.ReadFile("./data/points.json")
	if err != nil {
		t.Errorf("Expected to open the points.json file")
	}
}

func TestApiPoints(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()
	first, err := http.Get(fmt.Sprintf("%s/api/points", ts.URL))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if first.StatusCode != 400 {
		t.Errorf("Expected status code 400, got %v", first.StatusCode)
	}
	second, err := http.Get(fmt.Sprintf("%s/api/points?x=1&y=2&distance=100", ts.URL))
	if second.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %v", second.StatusCode)
	}
}

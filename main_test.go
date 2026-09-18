package main

import (
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	if w.Code != 200 {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	healthHandler(w, req)
	if w.Code != 200 {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

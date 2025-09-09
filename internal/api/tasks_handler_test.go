package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCreateTask(t *testing.T) {
	api := Application{
		TaskService: nil,
	}

	payload := map[string]any{
		"title":       "Learn TDD",
		"description": "hands on tdd",
		"priority":    1,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatal("failed parse request payload")
	}

	req := httptest.NewRequest("POST", "/api/v1/tasks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(api.handleCreateTask)
	handler.ServeHTTP(rec, req)

	t.Logf("Req.body %s\n", rec.Body.Bytes())
	if rec.Code != http.StatusCreated {
		t.Errorf("Status code diff got: %d, want: %d", rec.Code, http.StatusCreated)
	}

	var resBody map[string]any

	r := json.Unmarshal(rec.Body.Bytes(), &resBody)
	if r != nil {
		t.Fatalf("failed parse response %s", r.Error())
	}

	if resBody["title"] != payload["title"] {
		t.Errorf("title diff got: %s, wants: %s", resBody["title"], payload["title"])
	}
}

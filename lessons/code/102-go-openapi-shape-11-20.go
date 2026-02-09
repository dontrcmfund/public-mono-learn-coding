package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
GO OPENAPI SHAPE (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/102-go-openapi-shape-11-20.go
2) GET /openapi.json and inspect sections

Extra context:
- lessons/notes/189-openapi-first-principles.md
*/

// LESSON 11: OpenAPI-like schema builder
// Why this matters: machine-readable contract enables tooling.
func buildOpenAPIDoc() map[string]any {
	return map[string]any{
		"openapi": "3.0.3",
		"info": map[string]any{
			"title":   "User API",
			"version": "1.0.0",
		},
		"paths": map[string]any{
			"/users": map[string]any{
				"get": map[string]any{
					"summary": "List users",
					"responses": map[string]any{
						"200": map[string]any{"description": "OK"},
					},
				},
				"post": map[string]any{
					"summary": "Create user",
					"requestBody": map[string]any{
						"required": true,
					},
					"responses": map[string]any{
						"201": map[string]any{"description": "Created"},
						"400": map[string]any{"description": "Validation error"},
					},
				},
			},
		},
		"components": map[string]any{
			"schemas": map[string]any{
				"User": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"id":    map[string]any{"type": "integer"},
						"email": map[string]any{"type": "string"},
						"name":  map[string]any{"type": "string"},
					},
					"required": []string{"id", "email", "name"},
				},
				"APIError": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"code":    map[string]any{"type": "string"},
						"message": map[string]any{"type": "string"},
					},
					"required": []string{"code", "message"},
				},
			},
		},
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// LESSON 12: OpenAPI endpoint
// Why this matters: serving spec from API reduces doc drift risk.
func openAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	writeJSON(w, http.StatusOK, buildOpenAPIDoc())
}

// LESSON 13: Docs entry endpoint
// Why this matters: discoverability for API consumers.
func docsIndexHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"contract": "/openapi.json",
		"version":  "1.0.0",
	})
}

// LESSON 14: Health endpoint
// Why this matters: operational baseline remains explicit.
func healthHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// LESSON 15: Router composition
// Why this matters: docs endpoints are part of service interface.
func buildMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/openapi.json", openAPIHandler)
	mux.HandleFunc("/docs", docsIndexHandler)
	mux.HandleFunc("/health", healthHandler)
	return mux
}

// LESSON 16: Version declaration habit
// Why this matters: versioning supports safe consumer upgrades.

// LESSON 17: Component schema reuse habit
// Why this matters: shared schemas reduce contract inconsistency.

// LESSON 18: Explicit error model habit
// Why this matters: clients handle failures predictably.

// LESSON 19: OpenAPI as source of truth habit
// Why this matters: contract-first thinking prevents accidental drift.

// LESSON 20: End-to-end OpenAPI server demo
// Why this matters: reinforces docs as runtime artifact, not static afterthought.
func main() {
	mux := buildMux()
	addr := ":8094"
	fmt.Println("Go OpenAPI shape lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go OpenAPI Shape 11-20

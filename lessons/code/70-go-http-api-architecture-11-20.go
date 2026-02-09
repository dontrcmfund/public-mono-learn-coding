package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

/*
GO HTTP API ARCHITECTURE (Lessons 11-20)

Suggested use:
1) Run: go run lessons/code/70-go-http-api-architecture-11-20.go
2) Call endpoints on localhost:8081

Extra context:
- lessons/notes/158-go-api-architecture-principles.md
- lessons/notes/156-go-api-principles.md
*/

type Subscriber struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

type SubscriberRepo interface {
	List() []Subscriber
	Add(email string) Subscriber
	Deactivate(email string) (Subscriber, error)
}

type InMemorySubscriberRepo struct {
	items []Subscriber
}

func (r *InMemorySubscriberRepo) List() []Subscriber {
	out := make([]Subscriber, len(r.items))
	copy(out, r.items)
	return out
}

func (r *InMemorySubscriberRepo) Add(email string) Subscriber {
	nextID := 1
	if len(r.items) > 0 {
		nextID = r.items[len(r.items)-1].ID + 1
	}
	s := Subscriber{ID: nextID, Email: email, Active: true}
	r.items = append(r.items, s)
	return s
}

func (r *InMemorySubscriberRepo) Deactivate(email string) (Subscriber, error) {
	for i, s := range r.items {
		if s.Email == email {
			if !s.Active {
				return s, errors.New("already inactive")
			}
			s.Active = false
			r.items[i] = s
			return s, nil
		}
	}
	return Subscriber{}, errors.New("not found")
}

type SubscriberService struct {
	repo SubscriberRepo
}

func normalizeEmail(raw string) string {
	return strings.TrimSpace(strings.ToLower(raw))
}

func validEmail(email string) bool {
	return strings.Contains(email, "@") && !strings.HasPrefix(email, "@") && !strings.HasSuffix(email, "@")
}

func (s *SubscriberService) Create(emailRaw string) (Subscriber, error) {
	email := normalizeEmail(emailRaw)
	if !validEmail(email) {
		return Subscriber{}, errors.New("invalid email")
	}
	for _, existing := range s.repo.List() {
		if existing.Email == email {
			return Subscriber{}, errors.New("duplicate email")
		}
	}
	return s.repo.Add(email), nil
}

func (s *SubscriberService) Deactivate(emailRaw string) (Subscriber, error) {
	email := normalizeEmail(emailRaw)
	if !validEmail(email) {
		return Subscriber{}, errors.New("invalid email")
	}
	return s.repo.Deactivate(email)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

type createRequest struct {
	Email string `json:"email"`
}

func buildMux(service *SubscriberService) *http.ServeMux {
	mux := http.NewServeMux()

	// LESSON 11-14: thin handlers with service delegation
	mux.HandleFunc("/subscribers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, http.StatusOK, service.repo.List())
		case http.MethodPost:
			var req createRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
				return
			}
			sub, err := service.Create(req.Email)
			if err != nil {
				statusCode := http.StatusBadRequest
				if err.Error() == "duplicate email" {
					statusCode = http.StatusConflict
				}
				writeJSON(w, statusCode, map[string]string{"error": err.Error()})
				return
			}
			writeJSON(w, http.StatusCreated, sub)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}
	})

	// LESSON 15-19: action endpoint
	mux.HandleFunc("/subscribers/deactivate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		var req createRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}
		sub, err := service.Deactivate(req.Email)
		if err != nil {
			mapping := map[string]int{
				"invalid email":     http.StatusBadRequest,
				"not found":         http.StatusNotFound,
				"already inactive":  http.StatusConflict,
			}
			statusCode, ok := mapping[err.Error()]
			if !ok {
				statusCode = http.StatusBadRequest
			}
			writeJSON(w, statusCode, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, sub)
	})

	// LESSON 20: architecture marker
	mux.HandleFunc("/architecture", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"domain":   "email rules",
			"service":  "create/deactivate",
			"repo":     "in-memory adapter",
			"transport": "net/http handlers",
		})
	})

	return mux
}

func main() {
	repo := &InMemorySubscriberRepo{}
	service := &SubscriberService{repo: repo}
	mux := buildMux(service)
	addr := ":8081"
	fmt.Println("Go API architecture lessons server on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("server error:", err)
	}
}

// End of Go HTTP API Architecture 11-20

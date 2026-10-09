package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"fullstack-ai-feedback/internal/ai"
	"fullstack-ai-feedback/internal/model"
)

type Server struct {
	aiService *ai.Service
}

func main() {
	s := &Server{aiService: ai.NewService()}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/feedback/process", s.handleProcessFeedback)

	// CORS Header für Frontend-Kommunikation
	handler := corsMiddleware(mux)

	log.Println("Go Backend running on :8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Server crash: %v", err)
	}
}

func (s *Server) handleProcessFeedback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.ProcessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	analysis, err := s.aiService.AnalyzeFeedback(r.Context(), req.Content)
	if err != nil {
		http.Error(w, "AI Processing failed", http.StatusInternalServerError)
		return
	}

	item := model.FeedbackItem{
		ID:           "fb_102",
		Customer:     req.Customer,
		Content:      req.Content,
		Category:     analysis.Category,
		Sentiment:    analysis.Sentiment,
		AISuggestion: analysis.DraftReply,
		CreatedAt:    time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(item)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}
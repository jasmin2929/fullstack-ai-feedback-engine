package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"fullstack-ai-feedback/internal/ai"
	"fullstack-ai-feedback/internal/model"
)

type FeedbackHandler struct {
	aiService *ai.Service
}

func NewFeedbackHandler(aiService *ai.Service) *FeedbackHandler {
	return &FeedbackHandler{aiService: aiService}
}

func (h *FeedbackHandler) ProcessFeedback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.ProcessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	analysis, err := h.aiService.AnalyzeFeedback(r.Context(), req.Content)
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
		CreatedAt:    time.Now().UTC(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(item)
}
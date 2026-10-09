package model

import "time"

type FeedbackItem struct {
	ID        string    `json:"id"`
	Customer  string    `json:"customer"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`  // Generiert von KI
	Sentiment string    `json:"sentiment"` // Generiert von KI: POSITIVE, NEUTRAL, NEGATIVE
	AISuggestion string `json:"ai_suggestion"`
	CreatedAt time.Time `json:"created_at"`
}

type ProcessRequest struct {
	Customer string `json:"customer"`
	Content  string `json:"content"`
}
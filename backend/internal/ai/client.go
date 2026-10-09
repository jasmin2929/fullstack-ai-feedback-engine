package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type AIAnalysis struct {
	Category     string `json:"category"`
	Sentiment    string `json:"sentiment"`
	DraftReply   string `json:"draft_reply"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// AnalyzeFeedback simuliert/integriert eine LLM-Auswertung mit Timeout-Schutz
func (s *Service) AnalyzeFeedback(ctx context.Context, content string) (*AIAnalysis, error) {
	// Context Timeout verhindert Hängenbleiben bei API-Calls
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Simulierte KI-Pipeline (Hier OpenAI Client einbinden)
	select {
	case <-time.After(300 * time.Millisecond): // Schnelle Verarbeitungszeit
		return &AIAnalysis{
			Category:   "Feature Request",
			Sentiment:  "POSITIVE",
			DraftReply: fmt.Sprintf("Thanks for the feedback regarding '%s'! We added this to our roadmap.", content),
		}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
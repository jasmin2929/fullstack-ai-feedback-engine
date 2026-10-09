package ai

import (
	"context"
	"testing"
	"time"
)

func TestAnalyzeFeedback(t *testing.T) {
	service := NewService()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := service.AnalyzeFeedback(ctx, "Great service!")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if res.Sentiment != "POSITIVE" {
		t.Errorf("expected POSITIVE, got %s", res.Sentiment)
	}
}
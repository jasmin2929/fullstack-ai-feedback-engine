export interface FeedbackItem {
  id: string;
  customer: string;
  content: string;
  category: string;
  sentiment: 'POSITIVE' | 'NEUTRAL' | 'NEGATIVE';
  ai_suggestion: string;
  created_at: string;
}
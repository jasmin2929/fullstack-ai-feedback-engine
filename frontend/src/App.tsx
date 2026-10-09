import React, { useState } from 'react';
import { FeedbackItem } from './types';

export default function App() {
  const [content, setContent] = useState('');
  const [customer, setCustomer] = useState('');
  const [loading, setLoading] = useState(false);
  const [feedbacks, setFeedbacks] = useState<FeedbackItem[]>([]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!content || !customer) return;

    setLoading(true);
    try {
      const res = await fetch('http://localhost:8080/api/feedback/process', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ customer, content }),
      });
      const data: FeedbackItem = await res.json();
      setFeedbacks((prev) => [data, ...prev]);
      setContent('');
      setCustomer('');
    } catch (err) {
      console.error('API Error:', err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ maxWidth: '700px', margin: '40px auto', fontFamily: 'sans-serif' }}>
      <h2>⚡ AI Feedback Intelligence Engine</h2>
      <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '10px' }}>
        <input
          type="text"
          placeholder="Customer Name / Email"
          value={customer}
          onChange={(e) => setCustomer(e.target.value)}
          style={{ padding: '8px', fontSize: '14px' }}
        />
        <textarea
          placeholder="Enter feedback or support ticket content..."
          value={content}
          onChange={(e) => setContent(e.target.value)}
          rows={3}
          style={{ padding: '8px', fontSize: '14px' }}
        />
        <button type="submit" disabled={loading} style={{ padding: '10px', fontWeight: 'bold' }}>
          {loading ? 'Processing via AI...' : 'Analyze & Save Ticket'}
        </button>
      </form>

      <h3 style={{ marginTop: '30px' }}>Processed Tickets ({feedbacks.length})</h3>
      <div style={{ display: 'flex', flexDirection: 'column', gap: '15px' }}>
        {feedbacks.map((fb) => (
          <div key={fb.id} style={{ border: '1px solid #ccc', padding: '15px', borderRadius: '6px' }}>
            <strong>{fb.customer}</strong> — <span>{fb.category}</span>
            <p>{fb.content}</p>
            <div style={{ background: '#f4f4f4', padding: '10px', borderRadius: '4px' }}>
              🤖 <strong>AI Suggested Reply:</strong> {fb.ai_suggestion}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
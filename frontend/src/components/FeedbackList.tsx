import React from 'react';
import { FeedbackItem } from '../types';

interface FeedbackListProps {
  items: FeedbackItem[];
}

export const FeedbackList: React.FC<FeedbackListProps> = ({ items }) => {
  if (items.length === 0) {
    return <p style={{ color: '#666', marginTop: '20px' }}>No tickets processed yet.</p>;
  }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '15px', marginTop: '20px' }}>
      {items.map((item) => (
        <div
          key={item.id}
          style={{
            border: '1px solid #e2e8f0',
            borderRadius: '8px',
            padding: '16px',
            backgroundColor: '#ffffff',
            boxShadow: '0 1px 3px rgba(0,0,0,0.1)',
          }}
        >
          <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '8px' }}>
            <strong>{item.customer}</strong>
            <span
              style={{
                fontSize: '12px',
                padding: '2px 8px',
                borderRadius: '12px',
                backgroundColor: item.sentiment === 'POSITIVE' ? '#dcfce7' : '#fee2e2',
                color: item.sentiment === 'POSITIVE' ? '#166534' : '#991b1b',
              }}
            >
              {item.category} • {item.sentiment}
            </span>
          </div>
          <p style={{ margin: '8px 0', color: '#334155' }}>{item.content}</p>
          <div
            style={{
              backgroundColor: '#f8fafc',
              borderLeft: '3px solid #3b82f6',
              padding: '10px',
              borderRadius: '0 4px 4px 0',
              marginTop: '10px',
            }}
          >
            <div style={{ fontSize: '12px', color: '#64748b', marginBottom: '4px' }}>
              🤖 AI Suggested Reply
            </div>
            <div style={{ fontSize: '14px', color: '#1e293b' }}>{item.ai_suggestion}</div>
          </div>
        </div>
      ))}
    </div>
  );
};
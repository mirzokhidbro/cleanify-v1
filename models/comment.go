package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	ModelType string    `json:"model_type"`
	ModelID   int       `json:"model_id"`
	Type      string    `json:"type"` // "text" yoki "voice"
	Message   string    `json:"message"`
	VoiceURL  string    `json:"voice_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

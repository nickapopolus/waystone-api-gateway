package data

import "time"

type URL struct {
	ID             string    `json:"id,omitempty"`
	OriginalURL    string    `json:"original_url"`
	Slug           string    `json:"slug"`
	CustomSlug     string    `json:"custom_slug,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	ShortURL       string    `json:"short_url,omitempty"`
	ClickCount     int64     `json:"click_count,omitempty"`
	MaxClicks      int64     `json:"max_clicks,omitempty"`
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	IsActive       bool      `json:"is_active,omitempty"`
	UserID         string    `json:"user_id,omitempty"`
}

type URLCreateRequest struct {
	OriginalURL    string    `json:"original_url"`
	CustomSlug     string    `json:"custom_slug,omitempty"`
	MaxClicks      int64     `json:"max_clicks,omitempty"`
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	IsActive       bool      `json:"is_active,omitempty"`
}

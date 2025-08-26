package types

import "time"

// User is a placeholder model.
type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Thread struct {
	ID             string     `json:"id"`
	OrganizationID string     `json:"organization_id"`
	From           string     `json:"from"`
	LastRead       *time.Time `json:"last_read_at"`
	CreatedAt      time.Time  `json:"created_at"`
	HasUnread      bool       `json:"has_unread"`
}

type Message struct {
	ID                string    `json:"id"`
	ThreadID          string    `json:"thread_id"`
	Sender            string    `json:"sender"`
	SenderDisplayName string    `json:"sender_display_name"`
	Direction         string    `json:"direction"`
	Type              string    `json:"type"`
	Body              string    `json:"body"`
	MediaURL          string    `json:"media_url,omitempty"`
	MediaCaption      string    `json:"media_caption,omitempty"`
	MediaMimeType     string    `json:"media_mime_type,omitempty"`
	Voice             bool      `json:"voice"`
	Animated          bool      `json:"animated"`
	CreateAt          time.Time `json:"created_at"`
	Timestamp         string    `json:"timestamp"`
}

type OtpRequest struct {
	Recipient string `json:"recipient"`
	Code      string `json:"code,omitempty"`
}

type OtpResponse struct {
	Status string            `json:"status"`
	Data   map[string]string `json:"data"`
}

type ThreadsResponse struct {
	Status string    `json:"status"`
	Data   []Thread  `json:"data"`
	Meta   Paginator `json:"meta"`
}

type MessagesResponse struct {
	Status string    `json:"status"`
	Data   []Message `json:"data"`
	Meta   Paginator `json:"meta"`
}

type Paginator struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Query struct {
	Limit     int
	Offset    int
	OrderBy   string
	Direction string
}

type WelcomeMessage struct {
	Status string `json:"status"`
	Data   struct {
		Status string `json:"status"`
	} `json:"data"`
}

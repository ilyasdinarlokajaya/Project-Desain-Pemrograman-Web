package model

// tags
type Article struct {
	Article_name       string `json:"Article"`
	Broadcast_schedule string `json:"Broadcast"`
	Valid_from         string `json:"From"`
	Valid_until        string `json:"Until"`
	Article_status     string `json:"Status"`
}

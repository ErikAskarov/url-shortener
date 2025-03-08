package models

type Url struct {
	ID        int    `json:"id"`
	Url       string `json:"url"`
	ShortCode string `json:"shortCode"`
	Expires   string `json:"expires"`
}

type ShortUrlResponse struct {
	Url string `json:"url"`
}

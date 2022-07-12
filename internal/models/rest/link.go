package rest

// Link defines urls to external resources
type Link struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}

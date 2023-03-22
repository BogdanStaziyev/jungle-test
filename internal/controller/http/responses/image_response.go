package response

import "time"

type Image struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Path   string `json:"image_path"`
	URL    string `json:"image_url"`
}

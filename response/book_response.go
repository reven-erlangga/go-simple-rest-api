package response

type BookResponse struct {
	Id             int64  `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Author         string `json:"author"`
	PublishDate    string `json:"publish_date"`
	ImageCoverPath string `json:"image_cover_path"`
}
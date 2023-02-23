package request

type CreateBookRequest struct {
	Title          string `validate:"required, min=5, max=25" json:"title"`
	Description    string `validate:"required, min=10, max=50" json:"description"`
	Author         string `validate:"required" json:"author"`
	PublishDate    string `validate:"required" json:"publish_date"`
	ImageCoverPath string `validate:"required" json:"image_cover_path"`
}

type UpdateBookRequest struct {
	Id             int64  `validate:"required"`
	Title          string `validate:"required, min=5, max=25" json:"title"`
	Description    string `validate:"required, min=10, max=50" json:"description"`
	Author         string `validate:"required" json:"author"`
	PublishDate    string `validate:"required" json:"publish_date"`
	ImageCoverPath string `validate:"required" json:"image_cover_path"`
}
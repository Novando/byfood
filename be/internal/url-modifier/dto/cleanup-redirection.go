package dto

type CleanupRequest struct {
	Url       string `json:"url" validate:"http_url"`
	Operation string `json:"operation" validate:"oneof=all canonical redirection"`
}

type CleanupResponse struct {
	ProcessedUrl string `json:"processed_url"`
}

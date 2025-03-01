package models

import "mime/multipart"

type ValidationRequest struct {
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	ID        string `json:"id"`
}

type ValidationResponse struct {
	Valid      bool           `json:"valid"`
	Extracted  *ExtractedInfo `json:"extracted,omitempty"`
	Confidence float64        `json:"confidence,omitempty"`
	Message    string         `json:"message,omitempty"`
}

type ProcessingRequest struct {
	Files []*multipart.FileHeader `form:"files"`
}

type ExtractedInfo struct {
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	ID        string `json:"id"`
}

type ExternalResponse struct {
	Valid bool   `json:"valid"`
	Error string `json:"error,omitempty"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

type DocumentType string

const (
	DocumentTypePDF  DocumentType = "pdf"
	DocumentTypeJPG  DocumentType = "jpg"
	DocumentTypePNG  DocumentType = "png"
	DocumentTypeDOCX DocumentType = "docx"
)

type DocumentInfo struct {
	Filename    string       `json:"filename"`
	Type        DocumentType `json:"type"`
	Size        int64        `json:"size"`
	ContentType string       `json:"contentType"`
}

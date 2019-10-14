package models

import (
	"time"

	"github.com/ashlinchak/bookmarks/lib/validators"
)

// Bookmark type.
type Bookmark struct {
	ID           int                            `json:"id"         gorm:"auto_increment;primary_key;column:id"`
	URL          string                         `json:"url"        gorm:"column:url;unique;not null"             validate:"required,url"`
	Title        string                         `json:"title"      gorm:"column:title"`
	CreateatedAt time.Time                      `json:"createdAt"  gorm:"column:created_at;not null"`
	Tags         []Tag                          `json:"tags"       gorm:"many2many:bookmarks_tags;"`
	Errors       []validators.ValidationMessage `json:"-"          gorm:"-"`
}

// Validate method implements validator.Validatable interface
func (b *Bookmark) Validate() map[string]map[string]string {
	return map[string]map[string]string{
		"URL": map[string]string{
			"required": "URL is required",
			"url":      "URL should have valid format. It must include protocol (HTTP, HTTPS, etc.)",
		},
	}
}

// IsValid validates bookmark model
func (b *Bookmark) IsValid() bool {
	isValid := true
	b.Errors = nil

	messages, _ := validators.Validate(b)

	if len(messages) > 0 {
		isValid = false
		b.Errors = messages
	}

	return isValid
}

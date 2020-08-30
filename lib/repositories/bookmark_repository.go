package repositories

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ashlinchak/bookmarks/lib/models"
	"gorm.io/gorm"
)

var bookmark models.Bookmark

// BookmarkRepository represents the tag's repository contract.
type BookmarkRepository struct {
	Conn *gorm.DB
}

// List returns an array of tags.
func (r *BookmarkRepository) List(tagNames []string) ([]models.Bookmark, error) {
	var data []models.Bookmark

	query := r.Conn.Preload("Tags")
	countTags := len(tagNames)

	if countTags > 0 {
		query = query.
			Joins("INNER JOIN bookmarks_tags on bookmarks_tags.bookmark_id = bookmarks.id").
			Joins("INNER JOIN tags ON tags.id = bookmarks_tags.tag_id").
			Where("tags.name IN (?)", tagNames).
			Group("bookmarks.id").
			Having("count(distinct tags.name) = ?", countTags)
	}

	query.Find(&data)

	return data, nil
}

// Add bookmark
func (r *BookmarkRepository) Add(url string, title string, tags []string, notes string) (*models.Bookmark, error) {
	url = strings.TrimSpace(url)
	title = strings.TrimSpace(title)
	notes = strings.TrimSpace(notes)

	if len(title) == 0 {
		title = url
	}

	bookmark := models.Bookmark{
		URL:          url,
		Title:        title,
		Notes:        notes,
		CreateatedAt: time.Now(),
	}

	tx := r.Conn.Begin()

	if len(tags) > 0 {
		tagModels, err := findOrCreateTags(tags, tx)

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		bookmark.Tags = tagModels
	}

	if !bookmark.IsValid() {
		tx.Rollback()
		return &bookmark, errors.New("Bookmark is invalid")
	}

	if err := tx.Create(&bookmark).Error; err != nil {
		tx.Rollback()
		return &bookmark, err
	}

	tx.Commit()

	return &bookmark, nil
}

// Update bookmark
func (r *BookmarkRepository) Update(url string, newURL string, title string, tags []string, notes string) (*models.Bookmark, error) {
	newURL = strings.TrimSpace(newURL)
	title = strings.TrimSpace(title)
	notes = strings.TrimSpace(notes)

	tx := r.Conn.Begin()

	tx.Where("url = ?", url).First(&bookmark)

	if len(newURL) > 0 {
		bookmark.URL = newURL
	}

	if len(title) > 0 {
		bookmark.Title = title
	}

	if len(notes) > 0 {
		bookmark.Notes = notes
	}

	if len(tags) > 0 {
		tagModels, err := findOrCreateTags(tags, tx)

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Model(&bookmark).Association("Tags").Replace(tagModels); err != nil {
			tx.Rollback()
			return &bookmark, err
		}
	}

	if !bookmark.IsValid() {
		tx.Rollback()
		return &bookmark, errors.New("Bookmark is invalid")
	}

	if err := tx.Save(&bookmark).Error; err != nil {
		tx.Rollback()
		return &bookmark, err
	}

	tx.Commit()

	return &bookmark, nil
}

// DeleteByURL removes bookmark and clear not active tags
func (r *BookmarkRepository) DeleteByURL(url string) (err error) {
	tx := r.Conn.Begin()

	tx.Where("url = ?", url).First(&bookmark)

	if bookmark.ID == 0 {
		message := fmt.Sprintf("Bookmark with \"%s\" URL not found", url)
		return errors.New(message)
	}

	if err = tx.Delete(&bookmark).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return
}

func normalizedTagName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)

	return name
}

func findOrCreateTags(tags []string, tx *gorm.DB) (tagModels []models.Tag, err error) {
	for _, tag := range tags {
		var tagModel models.Tag

		tagName := normalizedTagName(tag)

		if err = tx.FirstOrCreate(&tagModel, models.Tag{Name: tagName}).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		tagModels = append(tagModels, tagModel)
	}

	return
}

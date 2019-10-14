package repositories

import (
	"github.com/ashlinchak/bookmarks/lib/models"
	"github.com/jinzhu/gorm"
)

// TagRepository represents the tag's repository contract.
type TagRepository struct {
	Conn *gorm.DB
}

// List returns an array of tags.
func (r *TagRepository) List(onlyActive bool) []models.Tag {
	var data []models.Tag

	query := r.Conn.Order("name ASC")
	if onlyActive {
		query = query.
			Joins("INNER JOIN bookmarks_tags ON bookmarks_tags.tag_id = tags.id").
			Joins("INNER JOIN bookmarks ON bookmarks.id = bookmarks_tags.bookmark_id")
	}

	query.Select("DISTINCT tags.*").Find(&data)

	return data
}

// DeleteNotActive removes not used tags
func (r *TagRepository) DeleteNotActive() (err error) {
	sql := `
		DELETE
		FROM tags
		WHERE tags.id IN (
			SELECT tags.id
			FROM tags
			LEFT JOIN bookmarks_tags ON bookmarks_tags.tag_id = tags.id
			LEFT JOIN bookmarks ON bookmarks_tags.bookmark_id = bookmarks.id
			GROUP BY tags.id
			HAVING COUNT(bookmarks.id) = 0
		)
	`
	return r.Conn.Exec(sql).Error
}

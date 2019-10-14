package models

// Tag type.
type Tag struct {
	ID   int    `json:"id"    gorm:"auto_increment;primary_key;column:id"`
	Name string `json:"name"  gorm:"column:name;not null;unique_index;"`
}

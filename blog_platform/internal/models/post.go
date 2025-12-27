package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	AuthorID  uuid.UUID `gorm:"type:uuid;not null" json:"author_id"`
	Title     string    `gorm:"not null" json:"title"`
	Slug      string    `gorm:"unique;not null" json:"slug"`
	Content   string    `gorm:"type:text" json:"content"`
	Status    string    `gorm:"default:'published'" json:"status"`
	Visibilty string    `gorm:"default:'public'" json:"visibility"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Tags      []Tag      `gorm:"many2many:post_tags;" json:"tags"`
	Comments  []Comment  `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"comments,omitempty"`
	Reactions []Reaction `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"reactions,omitempty"`
}

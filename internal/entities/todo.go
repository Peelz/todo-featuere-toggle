package entities

import "time"

type Todo struct {
	Title       string      `json:"title,omitempty"`
	Note        interface{} `json:"note,omitempty"`
	CreatorName string      `json:"creatorName,omitempty"`
	CreatedAt   *time.Time  `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time  `json:"updatedAt,omitempty"`
	DeletedAt   *time.Time  `json:"deletedAt,omitempty"`
}

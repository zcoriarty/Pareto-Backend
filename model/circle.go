package model

import (
	"time"
)

func init() {
	Register(&Circle{})
}

// Circle represents user domain model
type Circle struct {
	Base
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	ID           int        `json:"id,omitempty"`
	AccountID    string     `json:"account_id"`
	CircleSymbol string     `json:"circle_symbol,omitempty"`
	CircleName   string     `json:"circle_name,omitempty"`
	CircleBio    string     `json:"circle_bio,omitempty"`
}

// Delete updates the deleted_at field
func (u *Circle) Delete() {
	t := time.Now()
	u.DeletedAt = &t
}

// // Update updates the updated_at field
// func (u *Circle) CreateOrUpdate() {
// 	t := time.Now()
// 	u.UpdatedAt = t
// }

// UserRepo represents user database interface (the repository)
type CircleRepo interface {
	CreateOrUpdate(*Circle) (*Circle, error)
	View(int) (*Circle, error)
	List(*ListQuery, *Pagination) ([]Circle, error)
	Delete(*Circle) error
	

}


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
	CreatedAt                         *time.Time `json:"created_at,omitempty"`
	ID                          int        `json:"id"`
	AccountID                         string     `json:"account_id"`
	CircleSymbol					  string	 `json:"circle_symbol"`
	CircleName                        string     `json:"circle_name"`
	CircleBIO                         string     `json:"circle_bio"`
	
}


// Delete updates the deleted_at field
func (u *Circle) Delete() {
	t := time.Now()
	u.DeletedAt = &t
}

// Update updates the updated_at field
func (u *Circle) Update() {
	t := time.Now()
	u.UpdatedAt = t
}

// CircleRepo represents user database interface (the repository)
type CircleRepo interface {
	View(int) (*Circle, error)
	List(*ListQuery, *Pagination) ([]Circle, error)
	CreateOrUpdate(*Circle) (*Circle, error)
	Delete(*Circle) error

}


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
	CircleID                          int        `json:"circle_id"`
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
	FindByCircleName(string) (*Circle, error)
	List(*ListQuery, *Pagination) ([]Circle, error)
	Update(*Circle) (*Circle, error)
	Delete(*Circle) error
}

// // AccountRepo represents account database interface (the repository)
// type AccountRepo interface {
// 	Create(*Circle) (*Circle, error)
// 	CreateAndVerify(*Circle) (*Verification, error)
// 	CreateForgotToken(*Circle) (*Verification, error)
// 	CreateNewOTP(*Circle) (*Verification, error)
// 	CreateWithMobile(*Circle) error
// 	CreateWithMagic(*Circle) (int, error)
// 	ResetPassword(*Circle) error
// 	ChangePassword(*Circle) error
// 	UpdateAvatar(*Circle) error
// 	Activate(*Circle) error
// 	FindVerificationToken(string) (*Verification, error)
// 	FindVerificationTokenByCircle(*Circle) (*Verification, error)
// 	DeleteVerificationToken(*Verification) error
// }

// AuthCircle represents data stored in JWT token for user
type AuthCircle struct {
	ID       int
	CircleName string
	Role     AccessRole
}

package circle

import (
	"fmt"
	"net/http"

	"github.com/zcoriarty/Pareto-Backend/apperr"
	"github.com/zcoriarty/Pareto-Backend/model"
	"github.com/zcoriarty/Pareto-Backend/request"
	"github.com/zcoriarty/Pareto-Backend/secret"

	"github.com/zcoriarty/Pareto-Backend/repository/platform/query"
	"github.com/zcoriarty/Pareto-Backend/repository/platform/structs"

	"github.com/gin-gonic/gin"
)

// NewCircleService create a new circle application service
func NewCircleService(userRepo model.UserRepo, circleRepo model.CircleRepo, auth model.AuthService, rbac model.RBACService, secret secret.Service) *Service {
	fmt.Println("HERE8")
	return &Service{
		userRepo:   userRepo,
		circleRepo: circleRepo,
		auth:       auth,
		rbac:       rbac,
		secret:     secret,
	}
}

// Service represents the circle application service
type Service struct {
	userRepo   model.UserRepo
	circleRepo model.CircleRepo
	auth       model.AuthService
	rbac       model.RBACService
	secret      secret.Service
}

// List returns list of circles
func (s *Service) List(c *gin.Context, p *model.Pagination) ([]model.Circle, error) {
	u := s.auth.User(c)
	q, err := query.List(u)
	if err != nil {
		return nil, err
	}
	return s.circleRepo.List(q, p)
}

// View returns single circle
func (s *Service) View(c *gin.Context, id int) (*model.Circle, error) {
	if !s.rbac.EnforceUser(c, id) {
		return nil, apperr.New(http.StatusForbidden, "Forbidden")
	}
	return s.circleRepo.View(id)
}

// Update contains circle's information used for updating
type Update struct {
	ID           int
	CircleSymbol *string
	CircleName   *string
	CircleBio    *string
}

// Update updates circle's contact information
func (s *Service) CreateOrUpdate(c *gin.Context, update *Update) (*model.Circle, error) {
	fmt.Println("HERE7")
	if !s.rbac.EnforceUser(c, update.ID) {
		return nil, apperr.New(http.StatusForbidden, "Forbidden")
	}
	u, err := s.circleRepo.View(update.ID)
	if err != nil {
		return nil, err
	}
	structs.Merge(u, update)
	return s.circleRepo.CreateOrUpdate(u)
}

// Delete deletes a circle
func (s *Service) Delete(c *gin.Context, id int) error {
	fmt.Println("HERE6")
	u, err := s.circleRepo.View(id)
	if err != nil {
		return err
	}
	// if !s.rbac.IsLowerRole(c, u.Role.AccessLevel) {
	// 	return apperr.New(http.StatusForbidden, "Forbidden")
	// }
	u.Delete()
	return s.circleRepo.Delete(u)
}

// UpdateCircle updated user's circle
func (s *Service) UpdateCircle(c *gin.Context, update *request.UpdateC) (*model.Circle, error) {
	fmt.Println("HERE5")
	if !s.rbac.EnforceUser(c, update.ID) {
		return nil, apperr.New(http.StatusForbidden, "Forbidden")
	}
	u, err := s.circleRepo.View(update.ID)
	if err != nil {
		return nil, err
	}
	structs.Merge(u, update)
	return s.circleRepo.CreateOrUpdate(u)
}



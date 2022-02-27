package circle

import (
	"net/http"

	"github.com/zcoriarty/Pareto-Backend/apperr"
	"github.com/zcoriarty/Pareto-Backend/model"
	"github.com/zcoriarty/Pareto-Backend/repository/platform/query"
	"github.com/zcoriarty/Pareto-Backend/repository/platform/structs"

	"github.com/gin-gonic/gin"
)

// NewCircleService create a new circle application service
func NewCircleService(circleRepo model.CircleRepo, auth model.AuthService, rbac model.RBACService) *Service {
	return &Service{
		circleRepo: circleRepo,
		auth:     auth,
		rbac:     rbac,
	}
}

// Service represents the circle application service
type Service struct {
	circleRepo model.CircleRepo
	auth     model.AuthService
	rbac     model.RBACService
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
	ID        int
	CircleBio *string

}

// Update updates circle's contact information
func (s *Service) Update(c *gin.Context, update *Update) (*model.Circle, error) {
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

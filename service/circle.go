package service

import (
	"net/http"

	"github.com/zcoriarty/Pareto-Backend/apperr"
	"github.com/zcoriarty/Pareto-Backend/model"
	"github.com/zcoriarty/Pareto-Backend/repository/circle"
	"github.com/zcoriarty/Pareto-Backend/request"

	"github.com/gin-gonic/gin"
)

// Auth represents auth http service
type Circle struct {
	svc *circle.Service
}

func CirclesRouter(svc *circle.Service, r *gin.RouterGroup) {
	c := Circle{
		svc: svc,
	}
	cr := r.Group("/circles")
	cr.GET("", c.list)
	cr.GET("/:id", c.view)
	cr.PATCH("/:id", c.update)
	cr.DELETE("/:id", c.delete)

}

type CircleListResponse struct {
	Circles []model.Circle `json:"circles"`
	Page  int              `json:"page"`
}

func (u *Circle) list(c *gin.Context) {
	p, err := request.Paginate(c)
	if err != nil {
		return
	}
	result, err := u.svc.List(c, &model.Pagination{
		Limit: p.Limit, Offset: p.Offset,
	})
	if err != nil {
		apperr.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, CircleListResponse{
		Circles: result,
		Page:  p.Page,
	})
}

func (u *Circle) view(c *gin.Context) {
	id, err := request.ID(c)
	if err != nil {
		return
	}
	result, err := u.svc.View(c, id)
	if err != nil {
		apperr.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (u *Circle) update(c *gin.Context) {
	updateCircle, err := request.UpdateCircle(c)
	if err != nil {
		return
	}
	circleUpdate, err := u.svc.Update(c, &circle.Update{
		ID:        updateCircle.ID,
		CircleBio: updateCircle.CircleBIO,

	})
	if err != nil {
		apperr.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, circleUpdate)
}

func (u *Circle) delete(c *gin.Context) {
	id, err := request.ID(c)
	if err != nil {
		return
	}
	if err := u.svc.Delete(c, id); err != nil {
		apperr.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}


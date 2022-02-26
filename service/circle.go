package service

import (
	"net/http"

	"github.com/zcoriarty/Pareto-Backend/apperr"
	"github.com/zcoriarty/Pareto-Backend/model"
	"github.com/zcoriarty/Pareto-Backend/repository/circle"
	"github.com/zcoriarty/Pareto-Backend/request"

	"github.com/gin-gonic/gin"
)

// Circle represents the circle http service
type Circle struct {
	svc *circle.Service
}

// CircleRouter declares the routes for circle router group
func CircleRouter(svc *circle.Service, r *gin.RouterGroup) {
	u := Circle{
		svc: svc,
	}
	ur := r.Group("/circles")
	ur.GET("", u.list)
	ur.GET("/:id", u.view)
	ur.PATCH("/:id", u.update)
	ur.DELETE("/:id", u.delete)
}

type CircleListResponse struct {
	Circles []model.Circle `json:"circles"`
	Page  int          `json:"page"`
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
	updateCircle, err := request.CircleUpdate(c)
	if err != nil {
		return
	}
	circleUpdate, err := u.svc.Update(c, &circle.Update{
		CircleID:        updateCircle.CircleID,
		CircleName: updateCircle.CircleName,

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

package service

import (
	// "fmt"
	"net/http"

	"github.com/zcoriarty/Pareto-Backend/apperr"
	"github.com/zcoriarty/Pareto-Backend/model"
	"github.com/zcoriarty/Pareto-Backend/repository/circle"
	"github.com/zcoriarty/Pareto-Backend/request"

	"github.com/gin-gonic/gin"
)

// Auth represents auth http service
type CircleService struct {
	svc *circle.Service
}

func CircleRouter(svc *circle.Service, r *gin.RouterGroup) {
	c := CircleService{
		svc: svc,
	}
	cir := r.Group("/circles")
	// cir.GET("", c.list)
	cir.POST("/create_circle", c.createCircles)
	// cir.PATCH("", c.updateCircles)
	// cir.GET("/:id", c.view)
	// // cr.PATCH("/:id", c.update)
	// cir.DELETE("/:id", c.delete)

}

type CircleListResponse struct {
	Circle []model.Circle  `json:"circles"`
	Page  int              `json:"page"`
}


// func (a *CircleService) updateCircles(c *gin.Context) {
// 	p, err := request.UpdateCircle(c)
// 	if err != nil {
// 		return
// 	}
// 	circle, err := a.svc.UpdateCircle(c, p)
// 	if err != nil {
// 		apperr.Response(c, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, circle)
// }
type NewCircle struct {

	AccountID                         string `json:"account_id"`
	CircleSymbol                      string `json:"circle_symbol"`
	CircleName						  string `json:"circle_name"`
	CircleBio						  string `json:"circle_bio"`

}
func (a *CircleService) createCircles(c *gin.Context) {
	// r, err := request.AccountCreate(c)
	// if err != nil {
	// 	return
	// }
	var n NewCircle
	
	if err := c.ShouldBindJSON(&n); err != nil {
		apperr.Response(c, err)
	}
	circle := &model.Circle{
		
		AccountID:               n.AccountID,
		CircleSymbol:            n.CircleSymbol,
		CircleName:              n.CircleName,
		CircleBio:               n.CircleBio,
	}
	
	if err := a.svc.CreateCircle(c, circle); err != nil {
		apperr.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, circle)
}

func (u *CircleService) list(c *gin.Context) {
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
		Circle: result,
		Page:  p.Page,
	})
}

func (u *CircleService) view(c *gin.Context) {
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

// func (u *CircleService) update(c *gin.Context) {
// 	updateCircle, err := request.UpdateCircle(c)
// 	if err != nil {
// 		return
// 	}
// 	circleUpdate, err := u.svc.CreateOrUpdate(c, &circle.Update{
// 		CircleSymbol: updateCircle.CircleSymbol,
// 		CircleName:   updateCircle.CircleName,
// 		CircleBio:    updateCircle.CircleBio,

// 	})
// 	if err != nil {
// 		apperr.Response(c, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, circleUpdate)
// }

func (u *CircleService) delete(c *gin.Context) {
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




package service

import (
	"fmt"
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
	fmt.Println("HERE3")
	c := CircleService{
		svc: svc,
	}
	cir := r.Group("/circles")
	cir.GET("", c.list)
	// cir.POST("", c.createCircles)
	cir.PATCH("", c.updateCircles)
	cir.GET("/:id", c.view)
	// cr.PATCH("/:id", c.update)
	cir.DELETE("/:id", c.delete)

}

type CircleListResponse struct {
	Circle []model.Circle  `json:"circles"`
	Page  int              `json:"page"`
}


func (a *CircleService) updateCircles(c *gin.Context) {
	fmt.Println("HERE2")
	p, err := request.UpdateCircle(c)
	if err != nil {
		return
	}
	circle, err := a.svc.UpdateCircle(c, p)
	if err != nil {
		apperr.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, circle)
}

// func (a *CircleService) createCircles(c *gin.Context) {
// 	fmt.Println("HERE1")
// 	r, err := request.AccountCreate(c)
// 	if err != nil {
// 		return
// 	}
// 	circle := &model.Circle{
// 		CircleSymbol:            r.CircleSymbol,
// 		CircleName:              r.CircleName,
// 		CircleBio:               r.CircleBio,
// 	}
// 	user := &model.User{
// 		Username:            r.Username,
// 		Password:            r.Password,
// 		Email:               r.Email,
// 		FirstName:           r.FirstName,
// 		LastName:            r.LastName,
// 		RoleID:              r.RoleID,
// 		AccountID:           r.AccountID,
// 		AccountNumber:       r.AccountNumber,
// 		AccountCurrency:     r.AccountCurrency,
// 		AccountStatus:       r.AccountStatus,
// 		DOB:                 r.DOB,
// 		City:                r.City,
// 		State:               r.State,
// 		Country:             r.Country,
// 		TaxIDType:           r.TaxIDType,
// 		TaxID:               r.TaxID,
// 		FundingSource:       r.FundingSource,
// 		EmploymentStatus:    r.EmploymentStatus,
// 		InvestingExperience: r.InvestingExperience,
// 		PublicShareholder:   r.PublicShareholder,
// 		AnotherBrokerage:    r.AnotherBrokerage,
// 		DeviceID:            r.DeviceID,
// 		ProfileCompletion:   r.ProfileCompletion,
// 		ReferralCode:        r.ReferralCode,
// 	}
// 	if err := a.svc.CreateCircle(c, user, circle); err != nil {
// 		apperr.Response(c, err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, circle)
// }

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

func (u *CircleService) update(c *gin.Context) {
	updateCircle, err := request.UpdateCircle(c)
	if err != nil {
		return
	}
	circleUpdate, err := u.svc.CreateOrUpdate(c, &circle.Update{
		CircleSymbol: updateCircle.CircleSymbol,
		CircleName:   updateCircle.CircleName,
		CircleBio:    updateCircle.CircleBio,

	})
	if err != nil {
		apperr.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, circleUpdate)
}

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




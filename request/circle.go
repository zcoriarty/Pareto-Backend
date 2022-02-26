package request

import (
	"github.com/zcoriarty/Pareto-Backend/apperr"

	"github.com/gin-gonic/gin"
)

// UpdateUser contains user update data from json request
type UpdateCircle struct {

	CircleID                          int        `json:"circle_id"`
	AccountID                         string     `json:"account_id"`
	CircleSymbol					  string	 `json:"circle_symbol"`
	CircleName                        string     `json:"circle_name"`
	CircleBIO                         string     `json:"circle_bio"`
}

// UserUpdate validates user update request
func CircleUpdate(c *gin.Context) (*UpdateCircle, error) {
	var u UpdateCircle
	id, err := ID(c)
	if err != nil {
		return nil, err
	}
	if err := c.ShouldBindJSON(&u); err != nil {
		apperr.Response(c, err)
		return nil, err
	}
	u.CircleID = id
	return &u, nil
}

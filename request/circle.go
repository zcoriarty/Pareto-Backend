package request

import (
	"github.com/zcoriarty/Pareto-Backend/apperr"

	"github.com/gin-gonic/gin"
)

// UpdateC contains user update data from json request
type UpdateC struct {
	ID                                int        `json:"id"`
	AccountID                         *string     `json:"account_id"`
	CircleSymbol				  	  *string	 `json:"circle_symbol"`
	CircleName                        *string     `json:"circle_name"`
	CircleBIO                         *string     `json:"circle_bio"`
	
}

// UpdateCircle validates user update request
func UpdateCircle(c *gin.Context) (*UpdateC, error) {
	var u UpdateC
	id, err := ID(c)
	if err != nil {
		return nil, err
	}
	if err := c.ShouldBindJSON(&u); err != nil {
		apperr.Response(c, err)
		return nil, err
	}
	u.ID = id
	return &u, nil
}

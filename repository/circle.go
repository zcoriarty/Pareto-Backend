package repository

import (
	"fmt"
	"net/http"

	"github.com/go-pg/pg/v9/orm"
	"go.uber.org/zap"

	"github.com/zcoriarty/Pareto-Backend/apperr"
	"github.com/zcoriarty/Pareto-Backend/model"
)

// NewUserRepo returns a new UserRepo instance
func NewCircleRepo(db orm.DB, log *zap.Logger) *CircleRepo {
	fmt.Println("HERE9")
	return &CircleRepo{db, log}
}

// UserRepo is the client for our circle model
type CircleRepo struct {
	db         orm.DB
	log        *zap.Logger
	
}

// View returns single circle by ID
func (u *CircleRepo) View(id int) (*model.Circle, error) {
	var circle = new(model.Circle)
	sql := `SELECT "circle".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name"
	FROM "circles" AS "circle" LEFT JOIN "roles" AS "role" ON "role"."id" = "circle"."role_id"
	WHERE ("circle"."id" = ? and deleted_at is null)`
	_, err := u.db.QueryOne(circle, sql, id)
	if err != nil {
		u.log.Warn("circleRepo Error", zap.Error(err))
		return nil, apperr.New(http.StatusNotFound, "400 not found")
	}
	return circle, nil
}

// List returns list of all circles retreivable for the current circle, depending on role
func (u *CircleRepo) List(qp *model.ListQuery, p *model.Pagination) ([]model.Circle, error) {
	var circles []model.Circle
	q := u.db.Model(&circles).Column("circle.*", "Role").Limit(p.Limit).Offset(p.Offset).Where(notDeleted).Order("circle.id desc")
	if qp != nil {
		q.Where(qp.Query, qp.ID)
	}
	if err := q.Select(); err != nil {
		u.log.Warn("CircleDB Error", zap.Error(err))
		return nil, err
	}
	return circles, nil
}

// Create creates a new circle in our database.
func (a *CircleRepo) CreateOrUpdate(cir *model.Circle) (*model.Circle, error) {
	fmt.Println("HERE10")
	_circle := new(model.Circle)
	sql := `SELECT id FROM circles WHERE circle_symbol = ?`
	res, err := a.db.Query(_circle, sql, cir.CircleSymbol)
	if err == apperr.DB {
		a.log.Error("CircleRepo Error: ", zap.Error(err))
		return nil, apperr.DB
	}
	if res.RowsReturned() != 0 {
		// update..
		fmt.Println("updating...")
		_, err := a.db.Model(cir).Column(
			"created_at",
			"id",
			"account_id",
			"circle_symbol",
			"circle_name",
			"circle_bio",
		).WherePK().Update()
		if err != nil {
			a.log.Warn("CircleRepo Error: ", zap.Error(err))
			return nil, err
		}
		return cir, nil
	} else {
		// create
		fmt.Println("creating...")
		if err := a.db.Insert(cir); err != nil {
			a.log.Warn("CircleRepo error: ", zap.Error(err))
			return nil, apperr.DB
		}
	}
	return cir, nil
}

// Delete sets deleted_at for a circle
func (u *CircleRepo) Delete(circle *model.Circle) error {
	circle.Delete()
	_, err := u.db.Model(circle).Column("deleted_at").WherePK().Update()
	if err != nil {
		u.log.Warn("UserRepo Error", zap.Error(err))
	}
	return err
}

// // Create creates a new user in our database
// func (a *CircleRepo) CreateCircle(u *model.Circle) (*model.Circle, error) {
// 	fmt.Println("HERE11")
// 	circle := new(model.Circle)
// 	sql := `SELECT id FROM circles WHERE circle_symbol = ? AND deleted_at IS NULL`
// 	res, err := a.db.Query(circle, sql, u.CircleSymbol, u.CircleName, u.CircleBio)
// 	if err != nil {
// 		a.log.Error("AccountRepo Error: ", zap.Error(err))
// 		return nil, apperr.DB
// 	}
// 	if res.RowsReturned() != 0 {
// 		return nil, apperr.New(http.StatusBadRequest, "Circle already exists.")
// 	}
// 	if err := a.db.Insert(u); err != nil {
// 		a.log.Warn("AccountRepo error: ", zap.Error(err))
// 		return nil, apperr.DB
// 	}
// 	return u, nil
// }

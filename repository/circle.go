package repository

import (
	"net/http"

	"github.com/go-pg/pg/v9/orm"
	"go.uber.org/zap"

	"github.com/zcoriarty/Pareto-Backend/apperr"
	"github.com/zcoriarty/Pareto-Backend/model"
)


// NewCircleRepo returns a new CircleRepo instance
func NewCircleRepo(db orm.DB, log *zap.Logger) *CircleRepo {
	return &CircleRepo{db, log}
}

// CircleRepo is the client for our circle model
type CircleRepo struct {
	db  orm.DB
	log *zap.Logger
}

// View returns single circle by ID
func (u *CircleRepo) View(id int) (*model.Circle, error) {
	var circle = new(model.Circle)
	sql := `SELECT "circle".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name" 
	FROM "circles" AS "circle" LEFT JOIN "roles" AS "role" ON "role"."id" = "circle"."role_id" 
	WHERE ("circle"."id" = ? and deleted_at is null)`
	_, err := u.db.QueryOne(circle, sql, id)
	if err != nil {
		u.log.Warn("CircleRepo Error", zap.Error(err))
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

// Update updates circle's contact info
func (u *CircleRepo) Update(circle *model.Circle) (*model.Circle, error) {
	_, err := u.db.Model(circle).Column(
		"created_at",
		"circle_id",
		"account_id",
		"circle_symbol",
		"circle_name",
		"circle_bio",

	).WherePK().Update()
	if err != nil {
		u.log.Warn("CircleDB Error", zap.Error(err))
	}
	return circle, err
}

// Delete sets deleted_at for a circle
func (u *CircleRepo) Delete(circle *model.Circle) error {
	circle.Delete()
	_, err := u.db.Model(circle).Column("deleted_at").WherePK().Update()
	if err != nil {
		u.log.Warn("CircleRepo Error", zap.Error(err))
	}
	return err
}

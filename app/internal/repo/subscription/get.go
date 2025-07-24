package subscription

import (
	"app/internal/apperr"
	"app/internal/model"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"time"
)

func (s *SubscriptionRepo) GetByID(id string) (model.Subscription, error) {

	query := `SELECT id, name, price, user_id, start_date, end_date, created_at from services WHERE id = $1 `

	var tmp DTO
	err := s.db.Get(&tmp, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Println("GetByID error:", err)
			return model.Subscription{}, apperr.ErrNotFound
		}
		s.logger.Println("GetByID error:", err)
		return model.Subscription{}, err
	}
	res := model.Subscription{
		ID:        tmp.ID,
		Name:      tmp.Name,
		Price:     tmp.Price,
		UserID:    tmp.UserID,
		StartDate: tmp.StartDate,
		EndDate:   tmp.EndDate,
		CreatedAt: tmp.CreatedAt,
	}
	return res, nil
}

func (s *SubscriptionRepo) GetAll(limit, offset int) (model.GetAllResponse, error) {

	query := `SELECT id, name, price, user_id, start_date, end_date, created_at from services LIMIT $1 OFFSET $2;`

	var tmp []DTO

	err := s.db.Select(&tmp, query, limit, offset)

	if err != nil {
		s.logger.Println("GetAll error:", err)
		return model.GetAllResponse{Data: make([]model.Subscription, 0)}, err
	}

	if len(tmp) == 0 {
		s.logger.Println("Not found subscription")
		return model.GetAllResponse{Data: make([]model.Subscription, 0)}, apperr.ErrNotFound
	}

	res := model.GetAllResponse{
		Data: make([]model.Subscription, 0, len(tmp)),
	}

	for _, subscription := range tmp {
		res.Data = append(res.Data, model.Subscription{
			ID:        subscription.ID,
			Name:      subscription.Name,
			Price:     subscription.Price,
			UserID:    subscription.UserID,
			StartDate: subscription.StartDate,
			EndDate:   subscription.EndDate,
			CreatedAt: subscription.CreatedAt,
		})
	}
	return res, nil
}

func (s *SubscriptionRepo) GetTotalPrice(userID, subscriptionName string, from, to time.Time) (int, error) {
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query := builder.
		Select("COALESCE(SUM(price), 0)").
		From("services").
		Where(squirrel.And{
			squirrel.GtOrEq{"start_date": from},
			squirrel.LtOrEq{"start_date": to},
		})

	if userID != "" {
		query = query.Where(squirrel.Eq{"user_id": userID})
	}
	if subscriptionName != "" {
		query = query.Where(squirrel.Eq{"name": subscriptionName})
	}

	sqlStr, args, err := query.ToSql()
	if err != nil {

		s.logger.Println("GetTotalPrice build error:", err)
		return 0, err
	}

	var total int
	err = s.db.Get(&total, sqlStr, args...)
	if err != nil {

		s.logger.Println("GetTotalPrice exec error:", err)
		return 0, err
	}

	if total == 0 {
		s.logger.Println("Not found")
		return 0, apperr.ErrNotFound
	}

	return total, nil
}

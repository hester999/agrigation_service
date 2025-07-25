package subscription

import (
	"app/internal/apperr"
	"app/internal/model"
	"database/sql"
	"errors"
)

func (s *SubscriptionRepo) Update(subscription model.Subscription) (model.UpdateResponse, error) {

	query := `UPDATE services SET id = $1, name = $2, price = $3, user_id = $4, start_date = $5, end_date = $6, created_at = $7 WHERE id = $1 RETURNING id, name, price, user_id, start_date, end_date, created_at`

	var tmp DTO
	err := s.db.Get(&tmp, query, subscription.ID, subscription.Name, subscription.Price, subscription.UserID, subscription.StartDate, subscription.EndDate, subscription.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Println("subscription not found")
			return model.UpdateResponse{}, apperr.ErrNotFound
		}
		s.logger.Println("update err:", err)
		return model.UpdateResponse{}, err
	}

	res := model.UpdateResponse{
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

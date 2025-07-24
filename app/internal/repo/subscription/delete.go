package subscription

import (
	"app/internal/apperr"
)

func (s *SubscriptionRepo) DeleteByID(id string) error {
	query := `DELETE FROM services WHERE id = $1`
	res, err := s.db.Exec(query, id)
	if err != nil {
		s.logger.Println("Delete error:", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		s.logger.Println("RowsAffected error:", err)
		return err
	}

	if rowsAffected == 0 {
		s.logger.Println("Delete error: subscription not found")
		return apperr.ErrNotFound
	}

	return nil
}

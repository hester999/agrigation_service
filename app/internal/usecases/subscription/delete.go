package subscription

func (s *SubscriptionUsecases) DeleteByID(id string) error {

	err := s.repo.DeleteByID(id)
	if err != nil {
		s.logger.Println("delete subscription error: ", err)
		return err
	}
	return nil
}

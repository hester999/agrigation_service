package service_usecases

func (s *ServiceUsecases) DeleteByID(id string) error {

	err := s.repo.DeleteByID(id)
	if err != nil {
		s.logger.Println("delete service error: ", err)
		return err
	}
	return nil
}

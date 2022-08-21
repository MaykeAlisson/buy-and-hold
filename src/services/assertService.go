package services

type accertService struct{}

func AssertService() *accertService {
	return &accertService{}
}

func (service *accertService) CreateAssert() {}

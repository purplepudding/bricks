package service

type Service struct {
}

func NewService() *Service {
	//TODO wiring
	return &Service{}
}

func (service *Service) Run() {
	//TODO start service running
	//TODO handle exit signals
}

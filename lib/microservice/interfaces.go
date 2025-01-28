package microservice

type Service interface {
	Run() error
	Wire() error
}

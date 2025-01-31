package microservice

type Service[T any] interface {
	Run() error
	Wire(cfg T) error
}

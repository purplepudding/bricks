package microservice

type Service[T any] interface {
	Runnable
	Wireable[T]
}

type Runnable interface {
	Run() error
}

type Wireable[T any] interface {
	Wire(cfg T) error
}

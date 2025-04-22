package v1

//go:generate go tool mockgen -source auth/login_grpc.pb.go -destination auth/mock/login_grpc.pb.go -package mockauth
//go:generate go tool mockgen -source settings/global_grpc.pb.go -destination settings/mock/global_grpc.pb.go -package mocksettings
//go:generate go tool mockgen -source settings/service_grpc.pb.go -destination settings/mock/service_grpc.pb.go -package mocksettings

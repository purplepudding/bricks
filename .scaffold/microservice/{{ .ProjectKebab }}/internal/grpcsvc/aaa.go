package grpcsvc

import (
	{{.ProjectKebab}}v1 "github.com/purplepudding/foundation/api/pkg/pb/foundation/v1/{{.ProjectKebab}}"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ {{.ProjectKebab}}v1.AAAServiceServer = (*AAAService)(nil)

type AAAService struct {
  {{.ProjectKebab}}v1.UnimplementedAAAServiceServer
}
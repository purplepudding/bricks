package grpcsvc

import (
	{{.ProjectKebab}}v1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/{{.ProjectKebab}}"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ {{.ProjectKebab}}v1.AAAServiceServer = (*AAAService)(nil)

type AAAService struct {
  {{.ProjectKebab}}v1.UnimplementedAAAServiceServer
}
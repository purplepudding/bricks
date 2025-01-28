package test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func ProtoEq[T proto.Message](t *testing.T, expected, actual T, opts ...cmp.Option) {
	t.Helper()

	opts = append(opts, protocmp.Transform())

	assert.True(t, cmp.Equal(expected, actual, opts...), cmp.Diff(expected, actual, opts...))
}

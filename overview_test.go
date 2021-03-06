package rabbitclient

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
)

func TestConn_GetOverview(t *testing.T) {
	c := NewConn("guest", "guest")
	ctx := context.TODO()
	errC := make(chan error)
	outC := make(chan Overview)
	go c.GetOverview(ctx, "0.0.0.0:15672", outC, errC)
	select {
	case err := <-errC:
		assert.NoError(t, err)
	case o := <-outC:
		assert.NotNil(t, o)
	}
}

package test

import (

	"github.com/mlambda-net/net/pkg/net"
	"github.com/mlambda-net/template/pkg/application/message"
	"github.com/mlambda-net/template/pkg/infrastructure/endpoint/server"
	"github.com/stretchr/testify/assert"
	os "os"

	"testing"
)

func Test_Create_User(t *testing.T) {
	s := server.NewServer()
	s.Start()
	t.Run("ping", ping)

}


func ping(t *testing.T) {
	ping := &message.Ping{
		Id:   1,
	}
	_ = os.Setenv("Debug", "true")
	c := net.NewClient("localhost", "8000")
	r, err := c.Actor("template").Request(ping).Unwrap()
	assert.Nil(t, err)
	assert.NotNil(t, r)
}

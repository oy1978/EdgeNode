package nodes

import (
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs"
	"testing"
)

func TestListener_Listen(t *testing.T) {
	listener := NewListener()

	group := serverconfigs.NewServerAddressGroup("https://:1234")

	listener.Reload(group)
	err := listener.Listen()
	if err != nil {
		t.Fatal(err)
	}
}

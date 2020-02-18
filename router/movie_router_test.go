package movierouter

import (
	"testing"
	. "github.com/franela/goblin"
)

func TestXxx(t *testing.T) {
	g := Goblin(t)
	g.Describe("Teste", func() {
		g.It("Should add two numbers ", func() {
			g.Assert(1+1).Equal(2)
		})
	})
}


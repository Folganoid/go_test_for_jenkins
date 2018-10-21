package main

import (
	"testing"
	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Mult", func() {
		g.It("Should return mult of two numbers ", func() {
			g.Assert(mult(10, 10)).Equal(int64(100))
		})
	})

	g.Describe("Sum", func() {
		g.It("Should return sum of two numbers ", func() {
			g.Assert(sum(10, 10)).Equal(int64(21))
		})
	})

	g.Describe("Sum", func() {
		g.It("Should return sum of two numbers ", func() {
			g.Assert(!(sum(10, 11) == int64(20)))
		})
	})
}

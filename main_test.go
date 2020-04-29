package main

import (
	"testing"

	"github.com/blessedvictim/deptest"
)

func testA(t *testing.T) {
	t.Run("test A method", func(t *testing.T) {
		_ = deptest.A()
	})
}

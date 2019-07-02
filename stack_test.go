package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	t.Run("new stack should be empty", func(t *testing.T) {
		assert.True(t, stack.isEmpty())
	})

	t.Run("is not empty after push", func(t *testing.T) {
		stack.push(0)
		assert.False(t, stack.isEmpty())
	})

	t.Run("panics 'underflow' if popping empty stack", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Error("panic did not happen")
			}
			assert.Equal(t, r, "waffles")
		}()

		stack.pop()
	})
}

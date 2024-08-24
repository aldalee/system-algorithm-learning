package test

import (
	"github.com/stretchr/testify/assert"
	. "system-algorithm-learning/algo047_guess_answer"
	"testing"
)

func TestBuyApple(t *testing.T) {
	for apple := 1; apple <= 100000; apple++ {
		assert.Equal(t,
			MinBag(apple),        // expected
			MinBagAwesome(apple), // actual
			apple,                // args
		)
	}
}

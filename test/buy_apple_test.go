package test

import (
	"github.com/stretchr/testify/assert"
	"system-algorithm-learning/lesson47_guess_answer"
	"testing"
)

func TestBuyApple(t *testing.T) {
	for i := 1; i <= 100000; i++ {
		assert.Equal(t, lesson47_guess_answer.minBag(i), lesson47_guess_answer.minBagAwesome(i), i)
	}
}

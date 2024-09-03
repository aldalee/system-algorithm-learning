package test

import (
	"github.com/stretchr/testify/assert"
	. "system-algorithm-learning/algo047_guess_answer"
	"testing"
)

func TestCattleSheep(t *testing.T) {
	for i := range 50 {
		assert.Equal(t, WhoWin(i), WhoWinAwesome(i), i)
	}
}

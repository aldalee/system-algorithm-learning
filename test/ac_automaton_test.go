package test

import (
	. "system-algorithm-learning/algo033_ac_automaton"
	"testing"
)

func TestAcAutomaton(t *testing.T) {
	NewAcAutomaton("dhe", "he", "abcdhekse", "fuck", "小三", "69")
	contains := ContainWords("abcdhekskdjfafhasldkflskdjhwqaeruv6996《小三》")
	for _, ch := range contains {
		t.Log(ch)
	}
}

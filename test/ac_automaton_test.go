package test

import (
	"system-algorithm-learning/ac_automaton"
	"testing"
)

func TestAcAutomaton(t *testing.T) {
	ac_automaton.NewAcAutomaton("dhe", "he", "abcdhekse", "fuck", "小三", "69")
	contains := ac_automaton.ContainWords("abcdhekskdjfafhasldkflskdjhwqaeruv6996《小三》")
	for _, ch := range contains {
		t.Log(ch)
	}
}

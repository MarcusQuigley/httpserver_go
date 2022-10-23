package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Joe wins\n")
	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore, in}
	cli.PlayPoker()

	assertPlayerWins(t, playerStore, "Joe")
}

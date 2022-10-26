package poker

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoker() error {
	reader := bufio.NewScanner(cli.in)
	reader.Scan()
	player, err := extractWinner(reader.Text())
	if err != nil {
		return err
	}
	cli.playerStore.RecordWin(player)
	return nil
}

func extractWinner(input string) (string, error) {
	var stringToReplace = " wins"
	if len(input) == 0 {
		return "", errors.New("no player name exists")
	}
	player := strings.Replace(input, stringToReplace, "", 1)
	return player, nil
}

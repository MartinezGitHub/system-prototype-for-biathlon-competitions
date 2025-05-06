package lexer

import (
	"fmt"
	"strings"
)

func SecondToken(line string) (op string, err error) {

	firstSpace := strings.Index(line, " ")
	if firstSpace == -1 {
		return "", fmt.Errorf("no first space")
	}

	secondSpace := strings.Index(line[firstSpace+1:], " ")
	if secondSpace == -1 {
		return strings.TrimSpace(line[firstSpace+1:]), nil
	}

	op = line[firstSpace+1 : firstSpace+1+secondSpace]
	return op, nil
}

func AllTokens(line string) []string {
	return strings.Fields(line)
}

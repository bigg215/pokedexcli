package main

import (
	"fmt"
	"strings"
)

const (
	ALPHA string = "abcde"
)

func check(e error) string {
	return fmt.Sprintf("%s", e)
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

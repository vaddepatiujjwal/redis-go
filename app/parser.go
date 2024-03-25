package main

import (
	"strconv"
	"strings"
)

func parseCommand(input string) ([]string, error) {
	inputStrings := strings.Split(input, "\r\n")
	var resultStrings []string
	for idx, value := range inputStrings[1:] {
		if (idx+1)%2 == 0 {
			resultStrings = append(resultStrings, value)
		}
	}
	return resultStrings, nil
}

func encodeRedisString(input string) string {
	return "$" + strconv.Itoa(len(input)) + "\r\n" + input + "\r\n"
}

package hl7

import (
	"strings"
)

func SplitComponents(field string) []string {
	messageType := strings.Split(field, "^")
	return messageType
}

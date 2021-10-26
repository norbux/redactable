package redactable

import (
	"fmt"
	"strings"
)

const Delim = ":"

type RedactableString string

func (rs RedactableString) Redact() string {
	if i := strings.Index(string(rs), Delim); i > 0 {
		if x := strings.Index(string(rs)[i+1:], Delim); x > 0 {
			return fmt.Sprintf("%s*****%s", 
				string(rs)[0:i], 
				string(rs)[x + i + 2:])
		}
	}

	return string(rs)
}

func (rs RedactableString) Sprint() string {
	return strings.ReplaceAll(string(rs), Delim, "")
}
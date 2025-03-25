package cli

import (
	"strings"
	"fmt"
)

func commandLookUp(text string) (cliProc, error) {

	if len(commandMapping) == 0 {
		initCommandMapping()
	}

	cmd, ok := commandMapping[text]
	if !ok {
		return nil, fmt.Errorf("Unknown command")
	}
	return cmd.callback, nil
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	fields := strings.Fields(text)
	return fields
}

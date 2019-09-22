package release

import (
	"fmt"
	"os"
	"strings"
)

func parseRev(rev string) (head string, base string) {
	switch {
	case rev == "":
		return "master", ""
	case strings.Contains(rev, "..."):
		_range := strings.Split(rev, "...")
		return _range[0], _range[1]
	default:
		return rev, ""
	}
}

func exit1(message string) {
	fmt.Fprintf(os.Stderr, "%s\n", message)
	os.Exit(1)
}

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

package operator

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

// Changelog should be keep
type Changelog struct {
	Base string
	Head string

	// Major
	Removed []string
	Changed []string

	// Minor
	Added      []string
	Deprecated []string

	// Patch
	Fixed    []string
	Security []string

	// Consider as Minor
	MissingDefination []string
}

func (cl *Changelog) String() string {
	buf := new(bytes.Buffer)

	if cl.Base == "" {
		return "WARNNING: First release as version 1.0.0!!!"
	}
	buf.WriteString(fmt.Sprintf("## Changelog(time=\"%s\", ancestor=\"%s\")", time.Now().Format(("2006-01-02 15:04")), cl.Base))

	write := func(messages []string, name string) {
		if len(messages) > 0 {
			for left, right := 0, len(messages)-1; left < right; left, right = left+1, right-1 {
				messages[left], messages[right] = messages[right], messages[left]
			}
			buf.WriteString(fmt.Sprintf("\n\n* %s:", name))
			buf.WriteString("\n  * ")
			buf.WriteString(strings.Join(messages, "\n  * "))
		}
	}

	write(cl.Removed, "Removed")
	write(cl.Changed, "Changed")
	write(cl.Added, "Added")
	write(cl.Deprecated, "Deprecated")
	write(cl.Fixed, "Fixed")
	write(cl.Security, "Security")
	write(cl.MissingDefination, "Missing Defination")

	return buf.String()
}

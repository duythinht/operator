package operator

import (
	"bytes"
	"fmt"
)

type Release struct {
	Semver    *Semver
	Changelog *Changelog
}

func (rls *Release) String() string {

	buf := new(bytes.Buffer)

	buf.WriteString(fmt.Sprintf("Release(ref=\"%s\", name=\"%s\")\n\n", rls.Changelog.Head, rls.Semver.Name()))

	if rls.Changelog != nil {
		buf.WriteString(rls.Changelog.String())
	}
	return buf.String()
}

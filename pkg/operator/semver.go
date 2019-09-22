package operator

import (
	"fmt"
	"regexp"
	"strconv"
)

type Semver struct {
	Major int
	Minor int
	Patch int
}

// Tag2Semver convert Tag to semver
func Tag2Semver(tag string) *Semver {

	semver := &Semver{}

	re, _ := regexp.Compile("(\\d+).(\\d+).(\\d+)")
	subs := re.FindStringSubmatch(tag)

	semver.Major, _ = strconv.Atoi(subs[1])
	semver.Minor, _ = strconv.Atoi(subs[2])
	semver.Patch, _ = strconv.Atoi(subs[3])

	return semver
}

// LiftUp semver based on changelog
func (s *Semver) Upgrade(cl *Changelog) {
	switch {
	case len(cl.Removed)+len(cl.Changed) > 0:
		s.Major++
		s.Minor = 0
		s.Patch = 0
	case len(cl.Deprecated)+len(cl.Added)+len(cl.MissingDefination) > 0:
		s.Minor++
		s.Patch = 0
	case len(cl.Fixed)+len(cl.Security) > 0:
		s.Patch++
	}
}

func (s *Semver) Less(s1 *Semver) bool {
	return s.Major < s1.Major || (s.Major == s1.Major && ((s.Minor < s1.Minor) || (s.Minor == s1.Minor && s.Patch < s1.Patch)))
}

func (s *Semver) Name() string {
	return fmt.Sprintf("v%d.%d.%d", s.Major, s.Minor, s.Patch)
}

func (s *Semver) String() string {
	return fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
}

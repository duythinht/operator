package operator

type Repository interface {
	Changelog(base string, head string) (*Changelog, error)
	Tags() ([]string, error)
	Vers() ([]*Semver, error)
	Release(base string, head string) error
}

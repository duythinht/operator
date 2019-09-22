package operator

type Operator struct {
	Repo Repository
}

func (o *Operator) Changelog(base, head string) (*Changelog, error) {
	return o.Repo.Changelog(base, head)
}

func (o *Operator) Tags() ([]string, error) {
	return o.Repo.Tags()
}

func (o *Operator) Vers() ([]*Semver, error) {
	return o.Repo.Vers()
}

func (o Operator) Release(base, head string) error {
	return o.Repo.Release(base, head)
}

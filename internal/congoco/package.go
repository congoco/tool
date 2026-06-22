package congoco

type Package struct {
	Path              string
	Include           []string
	ChangelogFileName string
	ChangelogPath     string
	Commits           []*Commit
}

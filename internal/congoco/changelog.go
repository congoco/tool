package congoco

type ChangelogBlock struct {
	Version Version
	Commits []Commit
}
type Changelog struct {
	Blocks []ChangelogBlock
}

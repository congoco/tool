package congoco

type BelongStrategy string

const (
	ALL   BelongStrategy = "all"
	SCOPE BelongStrategy = "scope"
	PATH  BelongStrategy = "path"
)

type InvalidCommitsStrategy string

const (
	FAIL   InvalidCommitsStrategy = "fail"
	IGNORE InvalidCommitsStrategy = "ignore"
	OTHER  InvalidCommitsStrategy = "other"
)

type FormatterType string

const (
	INI  FormatterType = "ini"
	JSON FormatterType = "json"
	TXT  FormatterType = "txt"
)

type Flags struct {
	Changelog  ChangelogFlags
	Init       InitFlags
	Next       NextFlags
	Persistent PersistentFlags
	Validate   ValidateFlags
}

type PersistentFlags struct {
	Config    string
	Formatter string
}

type InitFlags struct {
	Force bool
}

type ValidateFlags struct {
	Message string
}

type NextFlags struct {
	Belong              string
	Commit              bool
	Invalid             string
	NoChangelog         bool
	NoVersionFileUpdate bool
	Push                bool
}

type ChangelogFlags struct {
	From    string
	Invalid string
	To      string
}

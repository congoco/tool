package flags

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
	BelongsStrategy   string
	Changelog         bool
	Commit            bool
	InvalidStrategy   string
	Push              bool
	VersionFileUpdate bool
}

type ChangelogFlags struct {
	From    string
	Invalid string
	To      string
}

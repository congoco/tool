package congoco

type Flags struct {
	Persistent PersistentFlags
	Init       InitFlags
	Validate   ValidateFlags
	Changelog  ChangelogFlags
}

type PersistentFlags struct {
	Formatter string
	Config    string
}

type InitFlags struct {
	Force bool
}

type ValidateFlags struct {
	Message string
}

type ChangelogFlags struct {
	From    string
	To      string
	Invalid string
}

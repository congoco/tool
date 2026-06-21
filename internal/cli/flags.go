package cli

type Flags struct {
	Persistent Persistent
	Root       Root
	Next       Next
}

type Persistent struct {
	Formatter string
	Config    string
}

type Root struct {
	Version bool
}

type Next struct {
	Version bool
}

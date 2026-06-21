package cli

type Flags struct {
	Persistent Persistent
	Root       Root
	Init       Init
	Next       Next
}

type Persistent struct {
	Formatter string
	Config    string
}

type Root struct{}

type Init struct {
	Force bool
}

type Next struct {
	Version bool
}

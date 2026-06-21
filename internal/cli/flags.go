package cli

type Flags struct {
	Persistent Persistent
	Root       Root
	Next       Next
}

type Persistent struct{}

type Root struct {
	Version bool
}

type Next struct {
	Version bool
}

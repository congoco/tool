package congoco

type Flags struct {
	Persistent Persistent
	Root       Root
	Init       Init
	Next       Next
	Validate   Validate
}

type Persistent struct {
	Formatter string
	Config    string
}

type Root struct{}

type Init struct {
	Force bool
}

type Validate struct {
	Message string
}

type Next struct {
	Version bool
}

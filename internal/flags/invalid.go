package flags

import "fmt"

type InvalidCommitsStrategy string

const (
	FAIL   InvalidCommitsStrategy = "fail"
	IGNORE InvalidCommitsStrategy = "ignore"
	OTHER  InvalidCommitsStrategy = "other"
)

func (b *InvalidCommitsStrategy) String() string {
	return string(*b)
}

func (b *InvalidCommitsStrategy) Type() string {
	return "belong"
}

func (b *InvalidCommitsStrategy) Set(s string) error {
	switch InvalidCommitsStrategy(s) {
	case FAIL, IGNORE, OTHER:
		*b = InvalidCommitsStrategy(s)
		return nil
	default:
		return fmt.Errorf("invalid strategy %q (allowed: fail, ignore, other)", s)
	}
}

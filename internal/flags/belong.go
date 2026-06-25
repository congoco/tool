package flags

import "fmt"

type BelongStrategy string

const (
	ALL   BelongStrategy = "all"
	PATH  BelongStrategy = "path"
	SCOPE BelongStrategy = "scope"
)

func (b *BelongStrategy) String() string {
	return string(*b)
}

func (b *BelongStrategy) Type() string {
	return "belong"
}

func (b *BelongStrategy) Set(s string) error {
	switch BelongStrategy(s) {
	case ALL, PATH, SCOPE:
		*b = BelongStrategy(s)
		return nil
	default:
		return fmt.Errorf("invalid strategy %q (allowed: all, path, scope)", s)
	}
}

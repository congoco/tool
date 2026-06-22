package congoco

import (
	"fmt"
)

type CommitType string

const (
	Fix      CommitType = "fix"
	Feat     CommitType = "feat"
	Build    CommitType = "build"
	Chore    CommitType = "chore"
	CI       CommitType = "ci"
	Docs     CommitType = "docs"
	Style    CommitType = "style"
	Refactor CommitType = "refactor"
	Perf     CommitType = "perf"
	Test     CommitType = "test"
)

var CommitTypeNames = map[CommitType]string{
	Fix:      "Bug Fix",
	Feat:     "Feature",
	Build:    "Build",
	Chore:    "Maintenance",
	CI:       "CI/CD",
	Docs:     "Documentation",
	Style:    "Style",
	Refactor: "Refactoring",
	Perf:     "Performance",
	Test:     "Tests",
}

type CommitMessage struct {
	BreakingChange bool
	Scope          string
	Subject        string
	Type           CommitType
}

type Commit struct {
	CommitMessage
	Files []string
}

func (t CommitType) Valid() bool {
	_, ok := CommitTypeNames[t]
	return ok
}

func ParseCommitType(s string) (CommitType, error) {
	t := CommitType(s)

	if !t.Valid() {
		return "", fmt.Errorf("Unknown commit type: %s", s)
	}

	return t, nil
}

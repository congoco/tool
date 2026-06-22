package congoco

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
	Type           CommitType
	Scope          string
	BreakingChange bool
	Subject        string
}

type Commit struct {
	CommitMessage
	Files []string
}

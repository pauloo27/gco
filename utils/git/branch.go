package git

func GetCurrentBranchName() (string, error) {
	return RunGitCommand("rev-parse", "--abbrev-ref", "HEAD")
}

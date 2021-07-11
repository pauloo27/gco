package git

import (
	"strings"
)

type ChangedFileStatus string

const (
	MODIFIED = ChangedFileStatus("M")
)

type ChangedFile struct {
	Name    string
	Status  ChangedFileStatus
	Tracked bool
}

func parseChangedFile(statusLine string) *ChangedFile {
	parts := strings.Split(strings.TrimSpace(statusLine), " ")

	// TL;DR: one "extra column" is outputed when the file is tracked
	// the output (with leading/trailing spaces trimed) looks like this:
	// (conside _ as white space)
	// - _M name -> when file is modified, but not added
	// - M_ name -> when file is modified and added
	// - MM name -> when file is modified and added, but the file was modified...
	// after the add. So, a extra column is output when the file is tracked

	tracked := len(parts) == 3
	var name string
	if tracked {
		name = parts[2]
	} else {
		name = parts[1]
	}
	status := parts[0]

	return &ChangedFile{
		Name:    name,
		Status:  ChangedFileStatus(status),
		Tracked: tracked,
	}
}

func GetChangedFiles() ([]*ChangedFile, error) {
	// will list all files
	allFiles, err := RunGitCommand("status", "--porcelain")
	if err != nil {
		return nil, err
	}
	files := []*ChangedFile{}
	if allFiles == "" {
		return files, nil
	}
	for _, fileStatus := range strings.Split(allFiles, "\n") {
		files = append(files, parseChangedFile(fileStatus))
	}
	return files, nil
}

package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertFiles(t *testing.T, rawFiles []string, expectedIndexes []int, result map[string]bool) {
	shouldBeAdded := func(i int) bool {
		for _, value := range expectedIndexes {
			if value-1 == i {
				return true
			}
		}
		return false
	}

	for i, file := range rawFiles {
		if result[file] != shouldBeAdded(i) {
			t.Fatalf("File %s was %v\nResult: %v.", file, result[file], result)
			return
		}
	}
}

func TestParseInstructions(t *testing.T) {
	rawFiles := []string{
		// 1					2									3									4
		"hello.txt", "secrets.sqlite", "passwords.txt", "wifi-cracker.js",
		// 5								6
		"despacito-2.mp3", "marvel-movie-about-super-heros.mp4",
	}

	getFilesMap := func(initialValue bool) (map[string]bool, []int) {
		filesMap := map[string]bool{}

		allIndexes := []int{}

		for i, file := range rawFiles {
			filesMap[file] = initialValue
			allIndexes = append(allIndexes, i+1)
		}
		return filesMap, allIndexes
	}

	t.Run("Simple number", func(t *testing.T) {
		filesMap, _ := getFilesMap(false)
		newMap, err := parseInstruction("1", filesMap, rawFiles)
		assert.Nil(t, err)
		assertFiles(t, rawFiles, []int{1}, newMap)
	})
	t.Run("Star", func(t *testing.T) {
		filesMap, allIndexes := getFilesMap(false)
		newMap, err := parseInstruction("*", filesMap, rawFiles)
		assert.Nil(t, err)
		assertFiles(t, rawFiles, allIndexes, newMap)
	})
	t.Run("Range", func(t *testing.T) {
		filesMap, _ := getFilesMap(false)
		newMap, err := parseInstruction("2-5", filesMap, rawFiles)
		assert.Nil(t, err)
		assertFiles(t, rawFiles, []int{2, 3, 4, 5}, newMap)
	})
	t.Run("Negate", func(t *testing.T) {
		filesMap, _ := getFilesMap(true)
		newMap, err := parseInstruction("^5", filesMap, rawFiles)
		assert.Nil(t, err)
		assertFiles(t, rawFiles, []int{1, 2, 3, 4, 6}, newMap)
	})
	t.Run("Negate range", func(t *testing.T) {
		filesMap, _ := getFilesMap(true)
		newMap, err := parseInstruction("^2-5", filesMap, rawFiles)
		assert.Nil(t, err)
		assertFiles(t, rawFiles, []int{1, 6}, newMap)
	})
	// TODO: instruction list
	// TODO: index out of bounds (negative and 0 as well)
	// TODO: invalid syntax: empty string
	// TODO: invalid syntax: multiple negates
	// TODO: invalid syntax: NaN
	// TODO: invalid syntax: invalid range (start after the end)
	// TODO: invalid syntax: invalid range (missing parts)
	// TODO: invalid syntax: invalid range (multiple -)
}

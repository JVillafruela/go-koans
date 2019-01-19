package go_koans

import (
	"io/ioutil"
	"strings"
)

func aboutFiles() {
	filename := "about_files.go"
	contents, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(contents), "\n")

	assert(lines[0] == "package go_koans") // handling files is too trivial
	assert(lines[5] == ")")                // handling files is too trivial
}

package i18n

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// Parse is a function for parsing translates according to provided directory with structured files
// into global storage
func Parse(dirpath string) {
	// List translation files
	dirfiles, err := ioutil.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}
	// Parse each page file
	for _, file := range dirfiles {
		// Read page file
		tfile, err := ioutil.ReadFile(dirpath + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		// Load page data
		tfiledata := File{}
		err = yaml.Unmarshal(tfile, &tfiledata)
		if err != nil {
			panic(err)
		}
		// Determine page name
		tfilename := strings.Split(file.Name(), ".")[0]
		// Save to global map
		files[tfilename] = tfiledata
	}
}

// Clean is a function for global storage cleanup
func Clean() {
	files = map[string]File{}
}

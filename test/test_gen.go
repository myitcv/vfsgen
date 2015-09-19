// +build generate

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
	"golang.org/x/tools/godoc/vfs/httpfs"
	"golang.org/x/tools/godoc/vfs/mapfs"
)

func main() {
	var fs http.FileSystem = httpfs.New(mapfs.New(map[string]string{
		"sample-file.txt":                "This file compresses well. Blaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaah!",
		"not-worth-compressing-file.txt": "Its normal contents are here.",
		"folderA/file1.txt":              "Stuff.",
		"folderA/file2.txt":              "Stuff.",
		"folderB/folderC/file3.txt":      "Stuff.",
		// TODO: Empty folder somehow?
		//"folder-empty/":                  "",
	}))

	config := vfsgen.Config{
		Input:   fs,
		Package: "test_test",
		Output:  "test_vfsdata_test.go",
	}

	err := vfsgen.Generate(config)
	if err != nil {
		log.Fatalln(err)
	}
}
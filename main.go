package main

import (
	"os"
	"path/filepath"
)

func main() {
	const root = "f:/2020_Geometry/OpenSceneGraph/include/"
	const suffix = ".h"

	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// fmt.Println(files)
	for _, f := range files {
		if filepath.Ext(f) == "" {
			os.Rename(f, f+suffix)
		}
	}
}

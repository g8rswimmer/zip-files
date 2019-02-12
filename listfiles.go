package main

import "archive/zip"

type file struct {
	Name      string `json:"name,omitempty"`
	Directory bool   `json:"directory,omitempty"`
}

func listFiles(name string) ([]file, error) {

	reader, err := zip.OpenReader(name)

	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var files []file

	for _, f := range reader.File {
		rf := file{
			Name:      f.Name,
			Directory: f.FileInfo().IsDir(),
		}

		files = append(files, rf)
	}

	return files, nil
}

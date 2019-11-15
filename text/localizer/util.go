package localizer

import (
	"os"
	"path/filepath"
)

func FindMessageFiles(folderPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if ext := filepath.Ext(path); ext == ".toml" || ext == ".yaml" {
			files = append(files, path)
		}

		return err
	})

	return files, err
}

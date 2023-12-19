package utils

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func readLines(fileName string) ([]string, error) {
	content, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}
	contentString := string(content)
	lines := strings.Split(contentString, "\n")
	var res []string
	for _, l := range lines {
		a := strings.TrimSpace(l)
		if a != "" {
			res = append(res, a)
		}
	}
	return res, nil
}

func appendFiles(odooConfig, filesPath string) error {
	files, err := os.ReadDir(filesPath)
	if err != nil {
		return nil
	}
	odooLines, err := readLines(odooConfig)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() || f.Name() == "odoorc" {
			continue
		}
		lines2, err := readLines(path.Join(filesPath, f.Name()))
		if err != nil {
			return err
		}
		if len(lines2) > 0 {
			odooLines = append(odooLines, "\n")
			odooLines = append(odooLines, lines2...)
		}
	}
	content := []byte(strings.Join(odooLines, "\n"))
	log.Info("Saving Odoo config file content")
	if err := os.WriteFile(odooConfig, content, 0o600); err != nil {
		return err
	}
	return nil
}

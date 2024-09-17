package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func GenerateSqlc(cfg GenerationConfig) error {
	if cfg.Target == "" {
		return errors.New("u have to set up a folder")
	}

	// Step 0: Delete Folder
	if cfg.Schema != cfg.Target {
		err := os.RemoveAll(cfg.Target)
		if err != nil {
			return err
		}
	}

	// Step 1: Copy Schema folder to Target folder
	if cfg.Schema != cfg.Target {
		err := copyDir(cfg.Schema, cfg.Target)
		if err != nil {
			return fmt.Errorf("failed to copy schema: %v", err)
		}
	}
	// Step 2: Create x_views.sql in target folder
	err := createSQLFile(cfg.Views, filepath.Join(cfg.Target, "x_views.sql"))
	if err != nil {
		return fmt.Errorf("failed to create x_views.sql: %v", err)
	}

	// Step 3: Create z_procedures.sql in target folder
	err = createSQLFile(cfg.Procedures, filepath.Join(cfg.Target, "z_procedures.sql"))
	if err != nil {
		return fmt.Errorf("failed to create z_procedures.sql: %v", err)
	}

	return nil
}

// Helper function to copy a directory
func copyDir(src string, dest string) error {
	return filepath.Walk(src,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			relativePath, err := filepath.Rel(src, path)
			if err != nil {
				return err
			}
			destPath := filepath.Join(dest, relativePath)

			if info.IsDir() {
				return os.MkdirAll(destPath, info.Mode())
			}

			return copyFile(path, destPath)
		})
}

// Helper function to copy a file
func copyFile(srcFilePath, destFilePath string) error {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}

// Helper function to create a SQL file by concatenating contents of all files in a folder
func createSQLFile(sourceDir, destFilePath string) error {
	fileInfos, err := os.ReadDir(sourceDir)
	if err != nil {
		return err
	}

	destFile, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			filePath := filepath.Join(sourceDir, fileInfo.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}

			if _, err := destFile.Write(data); err != nil {
				return err
			}

			_, _ = destFile.Write([]byte("\n-- End of " + fileInfo.Name() + " --\n\n"))
		}
	}

	return destFile.Sync()
}

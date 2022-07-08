package file

import (
	"os"
	"path/filepath"
)

func GetAppDir() (exPath string) {
	appPath, err := os.Executable()
	if err != nil {
		return ""
	}
	exPath = filepath.Dir(appPath)
	return exPath
}

func GetWorkDir() (exPath string) {
	exPath, err := os.Getwd()
	if err != nil {
		return
	}
	return exPath
}

func GetHomeDir() (exPath string) {
	exPath, err := os.UserHomeDir()
	if err != nil {
		return
	}
	return exPath
}

func GetTmpDir() (exPath string) {
	return os.TempDir()
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

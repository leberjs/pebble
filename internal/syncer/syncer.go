package syncer

import (
	"log"
	"os"
	"path/filepath"
)

var (
	homeDir, _   = os.UserHomeDir()
	configDir    = ".pebble"
	fileSyncPath = filepath.Join(homeDir, configDir, "sync")
)

func EnsureSyncDir() {
	os.MkdirAll(fileSyncPath, 0750)
}

func SyncFiles() error {
	// TODO: call function to sync from s3
	return nil
}

func GetSyncFiles() []string {
	f, err := os.ReadDir(fileSyncPath)
	if err != nil {
		log.Fatal(err)
	}

	files := make([]string, len(f))
	for i, file := range f {
		files[i] = file.Name()
	}

	return files
}

func GetFileContent(fn string) string {
	d, err := os.ReadFile(filepath.Join(fileSyncPath, fn))
	if err != nil {
		log.Fatal(err)
	}

	return string(d)
}

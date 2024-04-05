package syncer

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	homeDir, _   = os.UserHomeDir()
	configDir    = ".pebble"
	fileSyncPath = filepath.Join(homeDir, configDir, "sync")
)

func SyncFiles() error {
	// TODO: call function to sync from s3
	return nil
}

func GetSyncFiles() []string {
	f, err := os.ReadDir(fileSyncPath)
	if err != nil {
		log.Fatal(err)
	}

  var files []string
	for _, file := range f {
    if strings.HasSuffix(strings.ToLower(file.Name()), ".json") {
		  files = append(files, file.Name())
    }
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

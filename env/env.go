package env

import (
	"bufio"
	"mvrp/util"
	"os"
	"path/filepath"
	"strings"
)

func Init() {
	err := loadEnv()
	if err != nil {
		panic(err)
	}
}

func loadEnv() error {
	rootDir, err := util.Util.FS.FindProjectRoot("go.mod")
	if err != nil {
		return err
	}

	envPath := filepath.Join(rootDir, ".env")
	file, err := os.Open(envPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			// Skip comment lines and lines without '='
			continue
		}
		pair := strings.SplitN(line, "=", 2)
		err = os.Setenv(pair[0], pair[1])
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

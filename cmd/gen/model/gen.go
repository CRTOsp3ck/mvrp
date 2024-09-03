package model

import (
	"fmt"
	"log"
	"mvrp/errors"
	"mvrp/util"
	"os"
	"os/exec"
)

func Generate() error {
	return generateModels()
}

func generateModels() error {
	// List of directories and their corresponding commands
	commands := []struct {
		dir  string   // Directory to cd into
		args []string // Arguments for the sqlboiler command
	}{
		{"data/model/base", []string{"psql"}},
		{"../entity", []string{"psql"}},
		{"../inventory", []string{"psql"}},
		{"../invoice", []string{"psql"}},
		{"../item", []string{"psql"}},
		{"../purchase", []string{"psql"}},
		{"../sale", []string{"psql"}},
	}

	// Store the initial working directory
	originalDir, err := util.Util.FS.FindProjectRoot("go.mod")
	if err != nil {
		return errors.WrapError(errors.ErrTypeFileNotFound,
			fmt.Sprintf("failed to find project root: %v", err))
	}

	// Change to the specified directory
	if err := os.Chdir(originalDir); err != nil {
		return errors.WrapError(errors.ErrTypeInternal,
			fmt.Sprintf("failed to change directory to %s: %v", originalDir, err))
	}

	// Run each command in its corresponding directory
	for _, cmd := range commands {
		if err := runSQLBoilerCommand(cmd.dir, cmd.args...); err != nil {
			return errors.WrapError(errors.ErrTypeInternal,
				fmt.Sprintf("failed to run sqlboiler command: %v", err))
		}
	}

	// Return to the original directory after all commands are executed
	if err := os.Chdir(originalDir); err != nil {
		return errors.WrapError(errors.ErrTypeInternal,
			fmt.Sprintf("failed to return to the original directory: %v", err))
	}

	fmt.Printf("%d sqlboiler commands executed\n", len(commands))
	return nil
}

func runSQLBoilerCommand(directory string, args ...string) error {
	if err := os.Chdir(directory); err != nil {
		return errors.WrapError(errors.ErrTypeInternal,
			fmt.Sprintf("failed to change directory to %s: %v", directory, err))
	}

	cmd := exec.Command("sqlboiler", args...)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return errors.WrapError(errors.ErrTypeInternal,
			fmt.Sprintf("failed to run sqlboiler command in directory %s with args %v: %v", directory, args, err))
	}

	return nil
}

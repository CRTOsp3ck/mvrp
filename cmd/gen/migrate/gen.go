package migrate

import (
	"fmt"
	"log"
	"mvrp/errors"
	"mvrp/util"
	"os"
	"os/exec"
)

func Run() error {
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

	fmt.Println("Running sql-migrate down")
	// Run the sql-migrate down command to delete all migrations
	cmd := exec.Command("sql-migrate", "down", "-limit=0")
	cmd.Stdout = log.Writer() // To capture the command's output in the console
	cmd.Stderr = log.Writer() // To capture any errors in the console
	if err := cmd.Run(); err != nil {
		return errors.WrapError(errors.ErrTypeInternal,
			fmt.Sprintf("failed to run sql-migrate down: %v", err))
	}

	fmt.Println("Running sql-migrate up")
	// Run sql-migrate up to apply all migrations
	cmd = exec.Command("sql-migrate", "up")
	cmd.Stdout = log.Writer() // To capture the command's output in the console
	cmd.Stderr = log.Writer() // To capture any errors in the console
	if err := cmd.Run(); err != nil {
		return errors.WrapError(errors.ErrTypeInternal,
			fmt.Sprintf("failed to run sql-migrate up: %v", err))
	}

	// Return to the original directory after all commands are executed
	if err := os.Chdir(originalDir); err != nil {
		return errors.WrapError(errors.ErrTypeInternal,
			fmt.Sprintf("failed to return to the original directory: %v", err))
	}

	fmt.Println("Migrations completed")
	return nil
}

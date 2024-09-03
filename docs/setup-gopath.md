To add the Go binary path permanently to your PATH environment variable on Ubuntu, you need to modify your shell configuration file. Here are the steps:

1. **Open your shell configuration file**:
   Depending on the shell you are using, open the appropriate configuration file in a text editor. For example, if you are using `bash`, you would edit `~/.bashrc`. If you are using `zsh`, you would edit `~/.zshrc`.

   ```sh
   nano ~/.bashrc  # For bash
   # or
   nano ~/.zshrc   # For zsh
   ```

2. **Add the Go binary path**:
   Add the following line to the end of the file to include the Go binary path in your PATH environment variable:

   ```sh
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

3. **Save and close the file**:
   Save the changes and close the text editor. In `nano`, you can do this by pressing `Ctrl+O` to write the changes and `Ctrl+X` to exit.

4. **Reload the shell configuration**:
   Apply the changes by reloading the shell configuration file. You can do this by running:

   ```sh
   source ~/.bashrc  # For bash
   # or
   source ~/.zshrc   # For zsh
   ```

After completing these steps, the Go binary path will be permanently added to your PATH environment variable, and you should be able to run `sqlboiler` from any terminal session.
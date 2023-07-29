# gopyvenv
Tool to automatically activate your python virtual environment inside project folder.

Works with `venv` and `.venv` directories.

# Install
## Auto
You can run `install.sh` or `make install` for autpinstallation. This will install `gopyvenv` to `$HOME/.pyenv` and add include to your `.zshrc` file with neccessary hooks.

## Manual
You can also build the file with `make build` or `go build -o bin/ cmd/gopyvenv/gopyvenv.go` and copy from `bin/gopyvenv` somewhere inside your `$PATH`.

Hooks and neede environment variables can be found in `config/zsh.inc.dist` 

# How it works?
Binary will return specific command on new instances of zsh or change directory
based on current venv status.

## No venv is active
If you have `venv` or `.venv` subdir it will be activated.
## Venv is active
Going to child directory will not deactivate, only parent or totally differnt one will 
cause deactivation. 

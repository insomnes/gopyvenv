# gopyvenv
Tool to automatically activate your python virtual environment inside project folder.

Works with `venv` and `.venv` directories.

# Install
## Auto
You can run `install.sh` or `make install` for auto-installation. This will install `gopyvenv` to `$HOME/.pyenv` and add include to your `.bashrc` or `.zshrc` file with necessary hooks. 

## Manual
You can also build the file with `make build` or `go build -o bin/ cmd/gopyvenv/gopyvenv.go` and copy from `bin/gopyvenv` somewhere inside your `$PATH`.

Hooks and needed environment variables can be found in `config/zsh.inc.dist` or `config/bash.inc.dist`

# How it works?
Binary will return specific command on new instances of zsh or change directory
based on current venv status.

## No venv is active
If you have any of: `venv`,`.venv`,`virtenv`, `.virtenv` subdirs with avaliable `bin/activate` srcipt => it will be activated. Dir list is configurable by `GOPYVENV_DIR_NAMES` env var.
It will also try to search combinations of `projectdir-{iter of venv dirs}` or `projectdir_{iter of venv dirs}` like `gopyvenv-venv` or `gopyvenv_virtenv`. 

## Venv is active
Going to child directory will not deactivate, only parent or totally different one will 
cause deactivation. Going to dir with new virtual environment avaliable will cause 
deactivation of previous and activation of new virtual environment.


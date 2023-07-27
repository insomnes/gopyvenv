# gopyvenv
Tool to automatically activate your python virtual environment inside project folder.

Works with `venv` and `.venv` directories.

# Install
1. Build the binary via:
```shell
go build
```

2. Add binary to your $PATH or copy (as root):
```shell
cp gopyvenv /usr/local/bin/
```

3. Add this to your `~/.zshrc`:
```shell
# GOPYVENV AUTOINSTALL GOES AFTER THIS LINE
# This will help with new terminal instances
eval "$(/usr/local/bin/gopyvenv)"
# This will help with cd
python_venv() {
      eval "$(/usr/local/bin/gopyvenv)"
}
autoload -U add-zsh-hook
add-zsh-hook chpwd python_venv
# GOPYVENV AUTOINSTALL ENDS HERE
```

# How it works?
Binary will return specific command on new instances of zsh or change directory
based on current venv status.

## No venv is active
If you have `venv` or `.venv` subdir it will be activated.
## Venv is active
Going to child directory will not deactivate, only parent or totally differnt one will 
cause deactivation. 

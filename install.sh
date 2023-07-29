#!/usr/bin/env bash

GOPYVENV_ROOT="$HOME/.gopyvenv"
GOPYVENV_BIN="$GOPYVENV_ROOT/bin"
GOPYVENV_INC="$GOPYVENV_ROOT/zsh.inc"


COMMENT_STRING="# GOPYVENV AUTOINSTALL DO NOT CHANGE THIS LINE OR COMMENT"
INCLUDE_STRING="if [ -f '$GOPYVENV_INC' ]; then . '$GOPYVENV_INC';fi"
FULL_INCLUDE="$INCLUDE_STRING  $COMMENT_STRING"


RELOAD_CMD="source $HOME/.zshrc"

echo "Creating $GOPYVENV_ROOT if needed"
mkdir -p "$GOPYVENV_BIN"

echo "Building ..."
go build -o "$GOPYVENV_BIN/" cmd/gopyvenv/gopyvenv.go

echo "Making $HOME/.zshrc backup at $HOME/.zshrc.pre_gopyvenv.bcp"
cp "$HOME/.zshrc" "$HOME/.zshrc.pre_gopyvenv.bcp"

echo "Setting include"

if ! [[ -f "$GOPYVENV_INC" ]]; then
    cp ./config/zsh.inc.dist "$GOPYVENV_INC"
fi

INCLUDE_FOUND=$(grep "$COMMENT_STRING" "$HOME/.zshrc")
if [[ $INCLUDE_FOUND == "" ]] ; then
    echo "Include not found, adding"
    echo "$FULL_INCLUDE" >> "$HOME/.zshrc"
fi

echo "Finished gopyvenv installation, reload your profile with:"
echo "$RELOAD_CMD"


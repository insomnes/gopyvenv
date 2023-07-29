#!/usr/bin/env bash

GOPYVENV_ROOT="$HOME/.gopyvenv"
GOPYVENV_BIN="$GOPYVENV_ROOT/bin"

USER_SHELL="$(echo "$SHELL" | rev | cut -d '/' -f1 | rev)"


if [[ $USER_SHELL != "zsh" ]] && [[ $USER_SHELL != "bash" ]]; then
    echo "Unsupported shell for autoinstall: $SHELL"
    exit 1
fi


GOPYVENV_INC="$GOPYVENV_ROOT/$USER_SHELL.inc"
RC_FILE="$HOME/.${USER_SHELL}rc"

COMMENT_STRING="# GOPYVENV AUTOINSTALL DO NOT CHANGE THIS LINE OR COMMENT"
INCLUDE_STRING="if [ -f '$GOPYVENV_INC' ]; then . '$GOPYVENV_INC';fi"
FULL_INCLUDE="$INCLUDE_STRING  $COMMENT_STRING"


RELOAD_CMD="source $RC_FILE"

echo "Creating $GOPYVENV_ROOT if needed"
mkdir -p "$GOPYVENV_BIN"

echo "Building ..."
go build -o "$GOPYVENV_BIN/" cmd/gopyvenv/gopyvenv.go

echo "Making $RC_FILE backup at $RC_FILE.pre_gopyvenv.bcp"
cp "$RC_FILE" "$RC_FILE.pre_gopyvenv.bcp"

echo "Setting include"

if ! [[ -f "$GOPYVENV_INC" ]]; then
    cp "./config/${USER_SHELL}.inc.dist" "$GOPYVENV_INC"
fi

INCLUDE_FOUND=$(grep "$COMMENT_STRING" "${RC_FILE}")
if [[ $INCLUDE_FOUND == "" ]] ; then
    echo "Include not found, adding"
    echo "$FULL_INCLUDE" >> "$RC_FILE"
fi

echo "Finished gopyvenv installation, reload your profile with:"
echo "$RELOAD_CMD"


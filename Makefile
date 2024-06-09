# Parent Makefile

# List of child directories
CHILD_DIRS := gateway users

.PHONY: run-all $(CHILD_DIRS)

run-all: $(CHILD_DIRS)

$(CHILD_DIRS):
	@echo "Running make run in $@"
	@osascript -e 'tell application "iTerm2"' \
	            -e 'create window with default profile' \
	            -e 'tell current session of current window' \
	            -e 'write text "cd $(PWD)/$@ && make run"' \
	            -e 'end tell' \
	            -e 'end tell'

# Optionally, you can add other targets here
.PHONY: all

all: run-all
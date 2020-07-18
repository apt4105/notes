INDEX_FILE ?= index-dev.html
VNU ?= vnu-serve


.PHONY: help
help:
	@echo "ui makefile help:"
	@echo
	@echo "dev-server   start up a dev server and vnu validator, run with -j"
	@echo "help         show this message"
	@echo
	@echo "variables:"
	@echo "VNU          command for running an html validator. set to empty to skip"
	@echo "INDEX_FILE   the index.html file"

.PHONY: dev-server
dev-server: dev-server-elm dev-server-vnu

.PHONY: dev-server-elm
dev-server-elm:
	elm-live src/Main.elm --open --start-page=$(INDEX_FILE) -- --output=build/elm.js

.PHONY: dev-server-vnu
dev-server-vnu:
	if [ -n "$(VNU)" ]; then vnu-serve 8888; fi


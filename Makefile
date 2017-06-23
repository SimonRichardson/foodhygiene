all: dist/hygiene

install:
	go get github.com/Masterminds/glide
	go get github.com/mjibson/esc
	glide install
	$(MAKE) clean all

dist/hygiene:
	go build -o dist/hygiene github.com/SimonRichardson/foodhygiene/cmd/hygiene

.PHONY: build-ui
build-ui: ui/index.html pkg/ui/static.go

pkg/ui/static.go:
	esc -o="pkg/ui/static.go" -ignore="elm-stuff|Makefile|src|elm-package.json" -pkg="ui" ui

ui/index.html:
	$(MAKE) -C ./ui

clean: FORCE
	rm -rf dist/hygiene

clean-ui: clean
	rm -rf pkg/ui/static.go
	$(MAKE) -C ./ui clean

FORCE:

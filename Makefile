all: dist/hygiene

install:
	go get github.com/Masterminds/glide
	go get github.com/mjibson/esc
	glide install
	$(MAKE) -C ./ui
	$(MAKE) clean all

dist/hygiene:
	go build -o dist/hygiene github.com/SimonRichardson/foodhygiene/cmd/hygiene

.PHONY: build-ui
build-ui: ui/dist/index.html pkg/ui/static.go

pkg/ui/static.go:
	esc -o="pkg/ui/static.go" -ignore=".babelrc|Makefile|src|package.json|webpack.config.js|yarn.lock" -pkg="ui" ui

ui/dist/index.html:
	$(MAKE) -C ./ui

clean: FORCE
	rm -rf dist/hygiene

clean-ui: clean
	rm -rf pkg/ui/static.go
	$(MAKE) -C ./ui clean

FORCE:

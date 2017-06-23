all: node_modules dist/index.html

node_modules:
	yarn install

dist/index.html:
	./node_modules/.bin/webpack --progress -p

clean: FORCE
	rm -rf dist

FORCE:

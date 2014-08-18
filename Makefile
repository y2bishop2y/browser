PATH:=$(VEYRON_ROOT)/environment/cout/node/bin:$(PATH)
PATH:=node_modules/.bin:$(PATH)
CURRENT_DIR = $(shell pwd)
GOPATH:=$(CURRENT_DIR)/go:$(GOPATH)

# All JS and CSS files except build.js and third party
BROWSERIFY_FILES = $(shell find src -name "*.js" -o -name "*.css")
VDL_FILES = $(shell find go -name "*.vdl")
GO_FILES = $(shell find go -name "*.go")
BROWSERIFY_OPTIONS = --transform ./css-transform --debug
PROVA_OPTIONS = --browser --launch chrome --plugin proxyquireify/plugin
PROVA_HEADLESS_OPTIONS = --headless --progress --quit

# Builds everything
all: public/bundle.js public/bundle.html public/platform.js go/bin

# Creating the bundle JS file
public/bundle.js: $(BROWSERIFY_FILES) node_modules
	#jshint src # lint all src JavaScript files
	browserify src/app.js $(BROWSERIFY_OPTIONS) $< | exorcist $@.map > $@ #Browserify and generate map file

# Creating the bundle HTML file
public/bundle.html: web-component-dependencies.html node_modules bower_components
	vulcanize --output public/bundle.html web-component-dependencies.html --inline

# Copies the web components platform file
public/platform.js: bower_components
	cp bower_components/platform/platform.js public/platform.js

go/bin: $(GO_FILES)
	$(VEYRON_ROOT)/veyron/scripts/build/go install sample/...

go/src/sample/sampled/generated: $(VDL_FILES)
	(cd go/src/sample && $(VEYRON_ROOT)/veyron/go/bin/vdl generate --go_out_dir="generated" ... )

# Install what we need from NPM
node_modules: package.json
	npm prune
	npm install
	touch node_modules

# Install non-JS dependencies from bower
bower_components: bower.json node_modules
	bower prune
	bower install
	touch bower_components

# PHONY targets:

# Uses prova to run tests in a headless chrome and then quit after all test finish
test:
	jshint test # lint all test JavaScript files
	prova test/**/*.js $(PROVA_OPTIONS) $(PROVA_HEADLESS_OPTIONS)

# Continuously watch for changes to .js, .html or .css files.
# Rebundles the appropriate bundles when local files change
watch:
	watch -n 1 make

# Continuously reruns the tests as they change
watch-test:
	@echo "Tests being watched at: http://0.0.0.0:7559"
	prova test/**/*.js $(PROVA_OPTIONS)

# Serves the needed daemons and starts a server at http://localhost:$(HTTP_PORT)
# CTRL-C to stop
start:
	./services.sh

# Clean all build artifacts
clean:
	rm -f public/bundle.js
	rm -f public/bundle.html
	rm -f public/platform.js
	rm -rf node_modules
	rm -rf bower_components

.PHONY: start clean watch test watch-test
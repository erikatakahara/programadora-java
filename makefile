#!/usr/bin/env make -f

# Storage path while running the SDK localy
STORAGE_PATH:=$(shell pwd)/appengine-generated

# Path to dev_appserver.py
DEV_APPSERVER:=dev_appserver.py
APPCFG=appcfg.py

# Top level directory of the project
TOP:=$(CURDIR)

# Defaul values
IP:=0.0.0.0
PORT:=8080
ADMIN_IP=0.0.0.0
ADMIN_PORT=8081

# Run the dev_appserver locally
run:
	$(DEV_APPSERVER) --skip_sdk_update_check \
		--host ${IP} --port ${PORT} \
		--admin_host ${ADMIN_IP} --admin_port $(ADMIN_PORT) \
		--storage_path $(STORAGE_PATH) \
		--log_level critical \
		--require_indexes true \
		--enable_sendmail true \
		--use_mtime_file_watcher true \
		--automatic_restart true \
		app.yaml

# Run unit tests on all packages
test: check
check:
	GOPATH=$(GOPATH):$(PWD) go test -v ./...

# Generates documentation
doc: godoc
godoc:
	GOPATH=$(GOPATH):$(PWD) godoc -http=:3000

format: fmt
fmt:
	GOPATH=$(GOPATH):$(PWD) go fmt ./...

# Clean local app data on STORAGE_PATH
clean:
	rm -rvf appengine-generated

# Deploy to Google AppEngine
deploy:
	$(APPCFG) update app.yaml \
		--oauth2 --noauth_local_webserver
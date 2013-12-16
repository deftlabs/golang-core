#
# Copyright 2013, Deft Labs
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at:
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

SHELL := /bin/bash

compile:
	@cd src/deftlabs.com/util; go build
	@cd src/deftlabs.com/log; go build
	@cd src/deftlabs.com/kernel; go build
	@cd src/deftlabs.com/net/http; go build

clean:
	@rm -Rf bin
	@rm -Rf pkg

test: compile
	@cd src/deftlabs.com/net/http; go test
	@cd src/deftlabs.com/util; go test
	@cd src/deftlabs.com/log; go test
	@cd src/deftlabs.com/kernel; go test

initlibs:
	@go get github.com/mreiferson/go-httpclient
	@go get labix.org/v2/mgo
	@go get github.com/daviddengcn/go-ljson-conf
	@go get github.com/gorilla/mux

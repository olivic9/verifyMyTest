PATH_THIS:=$(realpath $(dir $(lastword ${MAKEFILE_LIST})))
DIR:=$(PATH_THIS)

include example.env
VARS:=$(shell sed -ne 's/ *\#.*$$//; /./ s/=.*$$// p' example.env )
$(foreach v,$(VARS),$(eval $(shell echo export $(v)="$($(v))")))

include .env
VARS:=$(shell sed -ne 's/ *\#.*$$//; /./ s/=.*$$// p' .env )
$(foreach v,$(VARS),$(eval $(shell echo export $(v)="$($(v))")))

help:
	@echo "    server"
	@echo "        Run server"
	@echo "    migrate"
	@echo "        Run migrations"
	@echo "    test"
	@echo "        Run tests"


.PHONY: server
server:
	./verify-my-test server

.PHONY: migrate
migrate:
	./verify-my-test migrate

.PHONY: test
test:
	@cd $(DIR) \
	&& go test ./test

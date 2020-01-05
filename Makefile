SWAG_VERSION:=${shell swag --version |sed 's/[^0-9\.]//g'}
SWAG_VERSION_MAJAR=$(shell echo $(SWAG_VERSION) | cut -f1 -d.)
SWAG_VERSION_MINOR=$(shell echo $(SWAG_VERSION) | cut -f2 -d.)
SWAG_GT_1_6_0:=$(shell [ $(SWAG_VERSION_MAJAR) -gt 1 -o \( $(SWAG_VERSION_MAJAR) -eq 1 -a $(SWAG_VERSION_MINOR) -ge 6 \) ] && echo true)

%:
	@:

api-doc:
	@if [ "${SWAG_GT_1_6_0}" != "true" ];then \
		echo "swag version v${SWAG_VERSION} is lower than 1.6.0, plz update swag"; exit 2; \
	fi
	swag init --parseDependency -g http.go -dir pkg/api -o pkg/api/docs


bin/api:
	go build -o bin/api cmd/api/*.go

run-api:
	go run cmd/api/*.go

db-migration-create:
	migrate create -ext sql -dir cmd/migration/migration_sql_scripts $(filter-out $@,$(MAKECMDGOALS))

.PHONY: api-doc run-api db-migration-create

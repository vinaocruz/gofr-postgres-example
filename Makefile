# note: call scripts from /scripts
include configs/.env
# export $(shell sed 's/=.*//' .env)

MIGRATION_PATH=internal/database/migrations

create_migration:
	@if [ -z "$(name)" ]; then \
        echo "Error: name is not set"; \
        exit 1; \
    fi
	migrate create -ext=sql -dir=$(MIGRATION_PATH) $(name)

migrate_up:
	migrate -path=$(MIGRATION_PATH) -database ${DATABASE_URL} -verbose up

migrate_down:
	migrate -path=$(MIGRATION_PATH) -database ${DATABASE_URL} -verbose down

.PHONY: create_migration migrate_up migrate_down

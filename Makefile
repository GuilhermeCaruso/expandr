define docker_compose_up
	docker-compose -f ./deployments/$(1)/docker-compose.yaml up -d
endef

.PHONY: run/local
run/local:
	@./$(go env GOPATH)/bin/air


.PHONY: setup/local
setup/local:
	$(call docker_compose_up,local)

.PHONY: migration/create
migration/create:
	@go run cmd/migrations/main.go db create_go $(name)

.PHONY: migration/execute
migration/execute:
	@go run cmd/migrations/main.go db migrate

.PHONY: migration/init
migration/init:
	@go run cmd/migrations/main.go db init
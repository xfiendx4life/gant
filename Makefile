base_path := $(abspath $(firstword $(MAKEFILE_LIST)))
init_dir := $(CURDIR)/init-db

run_db:
	docker stop postgres || true && docker rm postgres || true
	docker run --name postgres \
	-e POSTGRESQL_USERNAME=xfiendx \
	-e POSTGRESQL_DATABASE=diagram \
	-e POSTGRESQL_PASSWORD=123 \
	-p 5432:5432 \
	-v $(init_dir):/docker-entrypoint-initdb.d \
	bitnami/postgresql:latest

run_test_db:
	docker stop postgres || true && docker rm postgres || true
	docker run --name postgres \
	-e POSTGRESQL_USERNAME=test \
	-e POSTGRESQL_DATABASE=test_diagram \
	-e POSTGRESQL_PASSWORD=123 \
	-p 5432:5432 \
	-v $(init_dir):/docker-entrypoint-initdb.d \
	bitnami/postgresql:latest
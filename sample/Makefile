compose-up:
	docker-compose -f deployment/dev/docker-compose.yaml -p gaia up -d

migrate:
	migrate -path migration -database postgres://gaia:gaia123@localhost:5432/gaia?sslmode=disable up

gen-mocks:
	mkdir -p internal/mocks && \
	touch internal/mocks/a.txt && rm internal/mocks/* && mockery --dir internal --output internal/mocks --all

swagger:
	swag init -g rest.go -d internal/app/rest
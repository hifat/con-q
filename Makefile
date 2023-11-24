run:
	go run ./cmd/rest/main.go

swag-init:
	swag init --generalInfo=./internal/app/routes/routeV1/v1.go --instanceName v1

migrate-diff:
	atlas migrate diff --env dev

migrate-apply:
	atlas migrate apply --env dev
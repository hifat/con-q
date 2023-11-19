run:
	go run ./cmd/rest/main.go

swag-init:
	swag init --generalInfo=./internal/app/routes/routeV1/v1.go --instanceName v1

.PHONY: run build test

run:
	docker-compose up --build

test-backend:
	cd backend && go test ./... -v
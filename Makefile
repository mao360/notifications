build:
	docker-compose up -d --build notifications

run:
	docker-compose up -d notifications

test:
	go test -v

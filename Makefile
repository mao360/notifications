run:
	docker-compose up -d --build notifications

test:
	go test -v

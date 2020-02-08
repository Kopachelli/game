lint:
	golangci-lint run --enable-all --fix

test:
	go test

fmt:
	gofmt -w .

#docker

docker_up du:
	docker-compose up --build -d

docker_up_mac dum:
	docker-compose -f docker-compose.yml -f docker-compose-mac.yml up --build -d

docker_down dd:
	docker-compose down

SERVICE_NAME:=greeting-service

tests:
	go test --coverprofile=coverage.out ./...

curl:
	curl -X POST -H 'Content-Type: application/json' -i http://localhost:8000/ --data '{"id": "385ad292-9830-4f14-bd4c-9118d85f0db8", "jsonrpc": "2.0", "method": "GreetingService.Greeting", "params": { "name": "Vasy" }}'

# DOCKER
build: 
	docker build -t $(SERVICE_NAME) .

run:
	docker run -it --rm --name "$(SERVICE_NAME)" -p 8000:8000 $(SERVICE_NAME)

up: build run

stop:
	docker stop $(SERVICE_NAME); docker rm $(SERVICE_NAME)
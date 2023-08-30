build: 
	cd go-app/ && go build -o ./ main.go

run: build
	cd go-app/ && ./main

run-docker:
	docker compose build --no-cache && docker compose up
	
	

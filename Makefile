build:
	docker build --no-cache -t js-user-admin -f ./Dockerfile .

run:
	docker run -p 5050:5050 -i js-user-admin

deploy: build run

app:
	docker-compose -f docker-compose.yml up --build

log-server:
	docker-compose -f docker-compose-logs.yml up --detach --build

local-app:
	docker-compose -f docker-compose-local.yml up --build

test-app:
	docker-compose -f ./cli/script/docker-compose-tests.yml up --build
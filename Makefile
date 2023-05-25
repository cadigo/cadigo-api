gen:
	- go get github.com/99designs/gqlgen
	- go run github.com/99designs/gqlgen generate
	- go run plugin/generate.go
bu:
	- go build -ldflags="-s -w" cmd/api/main.go
	- strip main 
make-env:
	export $(grep -v '^#' .env | xargs)

run-dev:
	-export $(grep -v '^#' .env | xargs);
	-air

set-mod:
	go env -w GOPRIVATE=github.com/,github.com/oasis-prime/
	git config --global url."https://:x-oauth-basic@github.com".insteadOf "https://github.com"

dcup-build:
	docker build \
		--build-arg ACCESS_TOKEN= \
		-t cadigo-api -f ./build/Dockerfile .

dcup-local:
	docker-compose up

dcup-prod:
	docker-compose -f ./docker-compose.prod.yaml up --build

dc-down:
	docker-compose down

dc-clear:
	docker-compose down
	docker rmi -f ss-platform-api

# hosts:
# 	sudo -- sh -c "echo 127.0.0.1  larler-dev.com >> /etc/hosts"
# 	sudo -- sh -c "echo 127.0.0.1  api.larler-dev.com >> /etc/hosts"
# 	sudo -- sh -c "echo 127.0.0.1  admin.larler-dev.com >> /etc/hosts"

# rm-hosts:
# 	sudo -- sh -c "sed -i '' '/127.0.0.1 larler-dev.com/d' /etc/hosts"
# 	sudo -- sh -c "sed -i '' '/127.0.0.1 api.larler-dev.com/d' /etc/hosts"
# 	sudo -- sh -c "sed -i '' '/127.0.0.1 admin.larler-dev.com/d' /etc/hosts"

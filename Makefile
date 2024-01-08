.PHONY: build
build:
		rm -f material-react-table-app-api
		go build -o material-react-table-app-api cmd/main.go

.PHONY: build_image
build_image:
		docker build -t material-react-table-app-api .

.PHONY: delete_image
delete_image:
		docker rmi material-react-table-app-api:latest

.PHONY: start
start:
		docker-compose up -d

.PHONY: stop
stop:
		docker-compose stop && docker-compose rm -f

.PHONY: start_auth
start_auth: build
		./material-react-table-app-api

.PHONY: api_test
api_test:
		bash scripts/register_test.sh

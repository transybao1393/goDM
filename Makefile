BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} app/*.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t godm .

run:
	docker-compose up --build -d

stop:
	@echo "Stopping Docker. Please wait" 
	docker-compose down

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

rm-pg:
	@echo "Removing Postgres" 
# check if container is running and remove container
ifeq (!$(docker ps -q --filter ancestor=postgres),)
	docker rm $(docker ps -q --filter ancestor=postgres)
else
	docker rmi postgres
endif

rm-pga:
	@echo "Removing Postgres Admin"
# check if container is running and remove container
ifeq (!$(docker ps -q --filter ancestor=dpage/pgadmin4),)
	docker rm $(docker ps -q --filter ancestor=dpage/pgadmin4)
else
	docker rmi dpage/pgadmin4
endif

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint rm-pg rm-pga

help:
	@echo ''
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo '  test    	Go test'
	@echo '  engine  	Go build with custom engine'
	@echo '  docker  	Docker build'
	@echo '  run  		Run application through Docker Compose'
	@echo '  rm-pg     	Remove Postgres container and related image'
	@echo '  rm-pga   	Remove Postgres Admin container and related image'
	@echo ''
	@echo 'Extra arguments:'
	@echo '..will be update..'
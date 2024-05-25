SRC_DIR=src
BIN_NAME=server

all: build

build:
	go build -o ${BIN_NAME} ${SRC_DIR}/main.go

run: 
	go run ${SRC_DIR}/main.go

clean:
	go clean
	rm ${BIN_NAME}

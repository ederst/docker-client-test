OUTPUT_DIR?=./dist
BIN_FILE?=${OUTPUT_DIR}/docker-client-test

default: build run

build:
	mkdir -p ${OUTPUT_DIR}
	go build -o ${BIN_FILE}

run:
	${BIN_FILE}


clean:
	rm -rf ${OUTPUT_DIR}

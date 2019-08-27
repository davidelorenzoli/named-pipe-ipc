CURRENT_DIR=${PWD}
BIN_DIR=${CURRENT_DIR}/bin
DISTR_MAC=darwin
DISTR_WIN=windows
ARCH=amd64
OS := $(shell uname)

all: clean build run

clean:
	@echo
	@echo Cleaning...
	rm -rf ${BIN_DIR}

build:
	@echo
	@echo Building...
ifeq ($(OS), Darwin)
	export GOOS=${DISTR_MAC}
	export GOARCH=${ARCH}
	go build -v -o ${BIN_DIR}/spawned_process ./cmd/spawned_process
	go build -v -o ${BIN_DIR}/process_executor ./cmd/process_executor
else ifeq ($(OS), Windows)
	set GOOS=${DISTR_WIN}
	set GOARCH=${ARCH}
	go build -v -o ${BIN_DIR}/spawned_process.exe ./cmd/spawned_process
	go build -v -o ${BIN_DIR}/process_executor.exe ./cmd/process_executor
else
	@echo Failed to build. OS $(OS) not supported
endif

run:
	@echo
	@echo Executing...
	${BIN_DIR}/process_executor -executablePath ${BIN_DIR}/spawned_process
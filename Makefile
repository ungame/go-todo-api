GOC = go
BUILD_DIR = build
TESTS_DIR = tests

build:
	- mkdir ${BUILD_DIR}
	${GOC} build -o ${BUILD_DIR}/main.exe .

run:
	${BUILD_DIR}/main.exe

clear:
	- rmdir /q /s ${BUILD_DIR}

ut:
	go clean --cache && go test --cover go-todo-api/api/...

all: clear build ut run

it:
	- mkdir ${BUILD_DIR}
	${GOC} build -o ${BUILD_DIR}/it.exe ${TESTS_DIR}/main.go
	${BUILD_DIR}/it.exe

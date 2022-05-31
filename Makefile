.PHONY: install build

EXE_ENDING 				:=
ifeq ($(OS),Windows_NT)
	EXE_ENDING = .exe
endif

install:
	dart pub get -C ./manager
	go mod download
	go get .

build:
	-mkdir build
	dart compile exe ./manager/bin/manager.dart -o build/manager$(EXE_ENDING)
	go build -o build/

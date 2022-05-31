.PHONY: install build all

all: install compile

EXE_ENDING :=
ifeq ($(OS),Windows_NT)
	EXE_ENDING = .exe
endif

build/hash.json: ./hash.json
	[ -d build/ ] || mkdir -p build/
	cp hash.json build

build/cert/localhost.crt: ./cert/localhost.crt
	[ -d build/cert/ ] || mkdir -p build/cert/
	cp ./cert/localhost.crt build/cert

build/cert/localhost.decrypted.key: ./cert/localhost.decrypted.key
	[ -d build/cert/ ] || mkdir -p build/cert/
	cp ./cert/localhost.decrypted.key build/cert

install:
	dart pub get -C ./manager
	go mod download
	go get .

compile: build/hash.json build/cert/localhost.crt build/cert/localhost.decrypted.key
	dart compile exe manager/bin/manager.dart -o build/manager$(EXE_ENDING)
	go build -o build/undoneserv$(EXE_ENDING)

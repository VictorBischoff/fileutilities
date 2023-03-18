build-temp:
	cd src && go build -o tmp ../main.go
build-main:
	cd src && go build -o main ../main.go
.PHONY: build-main
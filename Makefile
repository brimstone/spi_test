export GOARCH=arm
export GOARM=5

spi_test: main.go
	go build -i -v

.PHONY: scp
scp: spi_test
	scp spi_test pi@10.35.4.125:

.PHONY: run
run:
	ssh pi@10.35.4.125 ./spi_test

.PHONY: watch
watch:
	find *.go | entr make scp run

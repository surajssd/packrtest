.PHONY: all
all: build build-fat

.PHONY: build
build:
	go build -o bin/test github.com/surajssd/packrtest/cmd

.PHONY: build-fat
build-fat:
	cd cmd; packr2
	go build -o bin/test-fat github.com/surajssd/packrtest/cmd
	cd cmd; packr2 clean

.PHONY: clean
clean:
	rm -vf bin/*

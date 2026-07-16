BINARY_NAME = rhymeswith
DICT_URL = https://raw.githubusercontent.com/cmusphinx/cmudict/master/cmudict.dict

.PHONY: all build clean dict run help

all: dict build

dict:
	@if [ ! -f cmudict.txt ]; then \
		echo "Downloading CMUdict..."; \
		curl -L -o cmudict.txt $(DICT_URL); \
	fi

build: dict
	go build -o $(BINARY_NAME) main.go

clean:
	rm -f $(BINARY_NAME) cmudict.txt

run: build
	./$(BINARY_NAME) boy

help:
	@echo "Available targets:"
	@echo "  make          - Download dict + build"
	@echo "  make build    - Build the binary"
	@echo "  make run      - Build and run with 'boy'"
	@echo "  make clean    - Remove binary and dict"

BINARY_NAME = rhymeswith
INSTALL_PATH = /usr/local/bin
DICT_URL = https://raw.githubusercontent.com/cmusphinx/cmudict/master/cmudict.dict

.PHONY: all build clean dict run install help

all: dict build

dict:
	@if [ ! -f cmudict.txt ]; then \
		echo "Downloading CMUdict..."; \
		curl -L -o cmudict.txt $(DICT_URL); \
	fi

build: dict
	go build -o $(BINARY_NAME) main.go

install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	sudo install -m 755 $(BINARY_NAME) $(INSTALL_PATH)
	@echo "Installation complete! You can now run 'rhymeswith <word>' from anywhere."

clean:
	rm -f $(BINARY_NAME) cmudict.txt

run: build
	./$(BINARY_NAME) boy

help:
	@echo "Available targets:"
	@echo "  make          - Download dict + build"
	@echo "  make build    - Build the binary"
	@echo "  make install  - Build and install system-wide (needs sudo)"
	@echo "  make run      - Build and test with 'boy'"
	@echo "  make clean    - Remove binary and dict"
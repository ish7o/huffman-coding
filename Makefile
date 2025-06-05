UnhuffDir = ./cmd/unhuff/
HuffDir = ./cmd/huff/
BinDir = ./bin

HuffBin = $(BinDir)/huff
UnhuffBin = $(BinDir)/unhuff

all: build

$(BinDir):
	mkdir -p $(BinDir)

$(HuffBin): $(BinDir)
	go build -o $(HuffBin) $(HuffDir)

$(UnhuffBin): $(BinDir)
	go build -o $(UnhuffBin) $(UnhuffDir)

build: $(HuffBin) $(UnhuffBin)

clean:
	rm -rf $(BinDir)

.PHONY: all build clean

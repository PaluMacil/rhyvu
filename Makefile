.PHONY: default clean

build: ry.exe

ry.exe:
	go build ./cmd/ry

clean:
	rm ry.exe
	rm DEBUG

example: 
	./ry.exe new 

test:
	go test ./...

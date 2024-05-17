

preBuild:
	go mod tidy



build: preBuild clean
	go build -o ./build/ ./cmd/buff
	go build -o ./build/ ./cmd/spam



buffRun:
	./build/buff

spamRun:
	./build/spam

clean:
	rm -rf ./build/
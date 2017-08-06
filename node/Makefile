TAG?=latest

all:
	go build -o ./bin/download-node ./main.go

run: all
	./bin/download-node
stop:
	ps aux | grep "./bin/download-node" | grep -v "grep" | awk '{print $$2}' | xargs kill -9 

image: all
	docker build -t riclava/download-node:$(TAG) .
push: image
	docker push riclava/download-node:$(TAG)

clean:
	rm -rf ./bin/download-node
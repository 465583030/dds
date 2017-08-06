TAG?=latest

all:
	go build -o ./bin/download-master ./src/main.go

run: all
	./bin/download-master
stop:
	ps aux | grep "./bin/download-master" | grep -v "grep" | awk '{print $$2}' | xargs kill -9 

image: all
	docker build -t riclava/download-master:$(TAG) .
push: image
	docker push riclava/download-master:$(TAG)

clean:
	rm -rf ./bin/download-master
TAG?=latest

all:
	go build -o ./bin/dds ./main.go

run: all
	./bin/start
stop:
	./bin/stop

image: all
	docker build -t riclava/dds:$(TAG) .
push: image
	docker push riclava/dds:$(TAG)

clean:
	rm -rf ./bin/dds
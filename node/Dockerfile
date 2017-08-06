# You must compile project on Linux
FROM scratch

MAINTAINER riclava <riclavaz@gmail.com>

ADD ./bin/download-node /

ENTRYPOINT ["/download-node", "--master=http://172.23.8.94:8080"]
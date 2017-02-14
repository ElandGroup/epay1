FROM jaehue/golang-onbuild
MAINTAINER jang.jaehue@eland.co.kr

# install go packages

# add application
WORKDIR /go/src/epaygo/main
ADD . /go/src/epaygo
RUN go install

EXPOSE 5000

CMD ["/go/bin/main"]
FROM golang:1.9.1
COPY . /go/src/AdAlpha

WORKDIR /go/src/app
COPY ./main.go .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8000
CMD ["go", "run", "main.go"]

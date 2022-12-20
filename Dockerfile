FROM golang:alpine

WORKDIR /app

COPY . .

RUN go install && go build -o helloworld

CMD [ "/app/helloworld" ]
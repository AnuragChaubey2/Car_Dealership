FROM golang:1.16-alpine

WORKDIR /github.com/Mini_Projects/car

COPY go.mod ./
COPY go.sum ./
RUN go mod download

 COPY . ./



RUN go build -o /car_dealership

EXPOSE 8080

CMD [ "/car_dealership" ]
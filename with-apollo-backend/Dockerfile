FROM golang:latest
RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/go-sql-driver/mysql
COPY . .
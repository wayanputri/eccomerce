FROM golang:1.20-alpine

# create directory folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable file with name "belajar_go"
RUN go build -o belajar_go

# run executable file
CMD ["./belajar_go"]
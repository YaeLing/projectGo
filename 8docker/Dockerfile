FROM golang:1.21.4
# setting workspace
WORKDIR ./dockerWork

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .
CMD ["./main"]
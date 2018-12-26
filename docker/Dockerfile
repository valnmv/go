FROM golang:1.9
COPY main.go /
RUN go build -o /main /main.go
EXPOSE 8000
ENTRYPOINT ["/main"]

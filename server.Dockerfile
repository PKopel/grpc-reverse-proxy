FROM golang

RUN mkdir /app
RUN mkdir /app/gen
COPY gen /app/gen
COPY server/main.go /app
COPY go.mod /app

WORKDIR /app
RUN go mod tidy
EXPOSE 50051
CMD ["go", "run", "./main.go", ":50051"]
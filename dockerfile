## Build
FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod tidy
COPY . ./

RUN go build -o /fc_optimal_assignment

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /fc_optimal_assignment /fc_optimal_assignment
COPY .env .

ENTRYPOINT ["/fc_optimal_assignment"]

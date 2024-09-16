FROM golang:1.22 as build
WORKDIR /build
COPY . .

RUN mkdir out && \
    mkdir out/database && \
    mv database/migrations/ out/database/migrations/ && \
    mv .config/ out/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -o out/app

FROM alpine
EXPOSE 80

WORKDIR /app

COPY --from=build /build/out ./
RUN chmod +x ./app

ENTRYPOINT ./app
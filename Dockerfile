FROM golang:1.25.3-trixie AS builder
COPY main.go go.* /src/
ENV CGO_ENABLED=0
RUN cd /src && go build -v -o /highscore

FROM scratch
COPY --from=builder /highscore /app/highscore
EXPOSE 8080
ENTRYPOINT ["/app/highscore"]

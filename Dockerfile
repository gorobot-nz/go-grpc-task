FROM golang:latest
ENV PROJECT_REPO=github.com/gorobot-nz/go-grpc-task
ENV APP_PATH=/go/src/${PROJECT_REPO}/
WORKDIR ${APP_PATH}
COPY . ${APP_PATH}
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

FROM alpine:latest
ENV PROJECT_REPO=github.com/gorobot-nz/go-grpc-task
ENV APP_PATH=/go/src/${PROJECT_REPO}/
RUN adduser -S nonrootuser
WORKDIR ${APP_PATH}
COPY --from=0 ${APP_PATH}/server ${APP_PATH}/server
COPY --from=0 ${APP_PATH}/config/config.yml ${APP_PATH}/config/config.yml
USER nonrootuser
EXPOSE 8080
CMD ["./server"]

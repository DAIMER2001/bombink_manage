FROM public.ecr.aws/docker/library/golang:1.19-alpine3.17 as builder

ARG GO_MODULES_TOKEN=token

RUN apk update

RUN apk add git
RUN apk add build-base
RUN apk add librdkafka librdkafka-dev pkgconf musl-dev
RUN git config --global url.https://$GO_MODULES_TOKEN@github.com/.insteadOf https://github.com/

ENV GOPRIVATE=github.com/lealco/leal/*
WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=arm64 go build -tags musl -tags dynamic cmd/main.go

# Run the Go Binary in Alpine.
FROM public.ecr.aws/docker/library/alpine:3.14

RUN apk update
RUN apk add build-base
RUN apk add librdkafka librdkafka-dev pkgconf musl-dev

RUN addgroup -S appgroup && adduser -S appuser -G appgroup --uid 10001 

WORKDIR /app
COPY --from=builder  app/main main
RUN chmod +x ./main
USER appuser
HEALTHCHECK CMD curl --fail http://localhost:8080/commerce/manage-bombink/healthz || exit 1

EXPOSE 8080 

CMD ["./main"]
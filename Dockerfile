FROM public.ecr.aws/docker/library/golang:1.24.2

WORKDIR /app

COPY go.mod go.sum ./
COPY vendor/ ./vendor/
COPY *.go ./

RUN go build -mod=vendor -o main .
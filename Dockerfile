FROM public.ecr.aws/docker/library/golang:1.24.2
COPY go.mod .
COPY go.sum .
COPY main.go .
COPY vendor/ ./vendor/
RUN go build -mod=vendor ./...
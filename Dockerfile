#FROM golang:1.16 as builder
#
#WORKDIR /app
#
#ENV GO111MODULE=on
#
#ENV CGO_ENABLED=0
#
#COPY . .
#
#RUN go build -o engine app/main.go
#
#FROM alpine
#
#WORKDIR /app
#
#COPY --from=builder /app/engine /app/
#COPY .env /app/
#
#EXPOSE 3000
#
#CMD ./engine rest

FROM golang:1.16 as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o engine app/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/engine /app
COPY . /app
EXPOSE 3000
CMD ./engine rest
FROM golang:1.11.0 as builder
RUN mkdir /app
COPY ./ /app/
WORKDIR /app/
RUN go get -d -v 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .


FROM scratch
COPY --from=builder /app/ /
CMD ["/main"]

FROM golang:1.11.0 as builder
RUN mkdir /app
COPY ./ /dohapp/
WORKDIR /dohapp/
RUN go get -d -v 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /dohapp/main .


FROM scratch
COPY --from=builder /dohapp/ /
CMD ["/main"]

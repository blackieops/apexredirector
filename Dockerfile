FROM golang:1.17-alpine
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /app
ADD . .
RUN go build -o apexredirector .

FROM scratch
COPY --from=0 /app/apexredirector /apexredirector
CMD ["/apexredirector"]

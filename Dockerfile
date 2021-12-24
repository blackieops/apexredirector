FROM alpine:3.15
ENV CGO_ENABLED=0 GOOS=linux
RUN apk --update --no-cache add go
WORKDIR /app
ADD . .
RUN go build -o apexredirector .

FROM scratch
COPY --from=0 /app/apexredirector /apexredirector
CMD ["/apexredirector"]

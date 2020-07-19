FROM alpine:3.12
ENV CGO_ENABLED=1 GOOS=linux
RUN apk --update --no-cache add go
WORKDIR /go/src/github.com/blackieops/apexredirector
ADD . .
RUN go build -a -ldflags '-linkmode external -extldflags "-static"' .

FROM scratch
COPY --from=0 /go/src/github.com/blackieops/apexredirector/apexredirector /apexredirector
CMD ["/apexredirector"]

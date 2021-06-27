FROM golang:alpine as build

COPY . mta-sts
WORKDIR mta-sts

ENV LDFLAGS -static
RUN apk --no-cache add bash git gcc musl-dev
RUN mkdir /pkg/
RUN go build -trimpath -buildmode pie --ldflags '-linkmode external -extldflags "-static"' -tags "osusergo netgo static_build" mta-sts.go
CMD ["bash"]

FROM scratch
COPY --from=build /go/mta-sts/mta-sts /go/bin/mta-sts

EXPOSE 8080
ENTRYPOINT ["/go/bin/mta-sts"]
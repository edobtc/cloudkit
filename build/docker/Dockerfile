FROM golang:1.19-alpine3.17 AS build

RUN apk update

RUN apk --no-cache add gcc g++ make ca-certificates
RUN apk update -qq && apk add git

RUN mkdir -p /go/src/github.com/edobtc/cloudkit

ADD . /go/src/github.com/edobtc/cloudkit
WORKDIR /go/src/github.com/edobtc/cloudkit

# build the various artifacts
RUN go build -o ./bin/btck ./cmd/btck
RUN go build -o ./bin/agent ./cmd/agent
RUN go build -o ./bin/controlplane ./cmd/controlplane
RUN go build -o ./bin/cli ./cmd/cli
RUN go build -o ./bin/relay ./cmd/relay
RUN go build -o ./bin/sink ./cmd/sink
RUN go build -o ./bin/functions ./cmd/functions

# Build production image
FROM alpine:3.17

RUN apk update
RUN apk add ca-certificates
RUN update-ca-certificates

RUN mkdir -p /var/app

COPY --from=build /go/src/github.com/edobtc/cloudkit/bin/agent /var/bin/agent
COPY --from=build /go/src/github.com/edobtc/cloudkit/bin/cli /var/bin/cli
COPY --from=build /go/src/github.com/edobtc/cloudkit/bin/btck /var/bin/btck
COPY --from=build /go/src/github.com/edobtc/cloudkit/bin/controlplane /var/bin/controlplane
COPY --from=build /go/src/github.com/edobtc/cloudkit/bin/relay /var/bin/relay
COPY --from=build /go/src/github.com/edobtc/cloudkit/bin/sink /var/bin/sink
COPY --from=build /go/src/github.com/edobtc/cloudkit/bin/functions /var/app/functions

# default run the cli
CMD ["/var/bin/cli"]

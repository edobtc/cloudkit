# Bitcoin CloudKit

![EDO](./docs/images/edobtc.png)

The open source toolkit for building enterprise grade, cloud first, bitcoin infrastructure.

Building Bitcoin based cloud infrastructure platforms requires an advanced knowledge and experience of both Bitcoin and Cloud infrastructure. This project seeks to bake best practices from both, into an opinionated and delightful set of tools to address undifferentiated problems in a set components we have observed that "everyone" needs, and provide them in a robust enough way to be used as is.

## Problems

Problems we seek to address

- node provisioning
- integration with various cloud and infrastructure projects (eg: AWS, Digital Ocean, Cloudflare...)
- monitoring, reliability
- platform components
- data interchange (streaming, queues, async)
- analytics and storage

## Components

Each of these components is in a variety of readiness states, from "an idea" to "pretty full and working". Please reach out if any of these are relevant to you and you want help adding behavior, or making them work for your use case

[control plane](./cmd/controlplane/)

An overall control plane for orchestration of a Bitcoin platform, ie: node and cloud component provisioning, upgrading, rollout, deploy

[multiplex](./cmd/multiplex/)

A data multiplexer, ie: data from one source (ie: an Eclair node) and be multiplexed to other data providers (ie: an AWS kinesis stream)

[agent](./cmd/agent/)

an instance agent + sidecar for provisioned nodes which supports reliabilty, monitoring and provisioning in addition to supporting integration with other cloudkit components.

[relay](./cmd/relay/)

A event aggregator, where events are streamed/watched and conditions can be established as a ruleset upon which action can be taken

[cli](./cmd/relay/)

A cli toolkit of helpers for interacting with nodes, or cloudkit itself

[gateway](./cmd/gateway/)

A permission control proxy that can sit in front of any type of node, and add a layer regarding authentication, authorization and reliability

## Support

We endeavor to support all parts of the "Bitcoin ecosystem" that developers may be building upon, currently Bitcoin core, and the lightning implementation. If there is something that you think should be supported, or that is new that we don't know about, let us know.

**Warning**

This is alpha/WIP software in active development. Everything subject to change.


## Development


## Setup

### macOS X

install golang (at time of writing `1.22.1`):

`brew install golang`

### Protocol Buffers / GRPC

We try and generate all types using protobuf and in most cases support these with a GRPC service (though not always)

The goal in defining types with protobuf is to ensure some amount of modularity in building clients, adopting SOME parts of the cloudkit and not others, or cross language/project interoperability.

Protobuf files are to be self contained in `/proto` and ideally versioned.

#### Buf

We are currently using [buf](https://buf/build) to generate go types from the proto files and using remote plugins to ensure you don't need to set anything up on your machine.

To generate:

`buf mod update`
`buf generate`


#### Testing

verify things are set up and working by running `go test ./...` from the root and if there is a fail, something has gone wrong

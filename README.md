# Bitcoin CloudKit

![EDO](./docs/images/edobtc.png)

The open source toolkit for building enterprise grade, cloud first, bitcoin infrastructure

**Warning**

This is alpha/WIP software in active development. Everything subject to change.

This is

## Setup

### macOS X

install golang (at time of writing `1.19`):

`brew install golang`

### Protocol Buffers / GRPC

We try and generate all types using protobuf and in most cases support these with a GRPC service ( though not always)

The goal in defining types with protobuf is to ensure some amount of modularity in building clients, adopting SOME parts of the cloudkit and not others, or cross language/project interoperability.

Protobuf files are to be self contained in `/proto` and ideally versioned.

#### Buf

We are currently using [buf](https://buf/build) to generate go types from the proto files and using remote plugins to ensure you don't need to set anything up on your machine.

To generate:

`buf mod update`
`buf generate`

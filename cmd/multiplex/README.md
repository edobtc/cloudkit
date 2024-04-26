# cloudkit-multiplex

# About

cloudkit-multiplex is a means by which node operation events can be multiplexed to other cloud native transports, for example:

- eclair websocket subscription -> AWS SNS topic
- bitcoin zmq subscription -> AWS kinesis stream
- lightning/lnd invoice created hook (using rpc middleware interceptor) -> AWS SQS queue

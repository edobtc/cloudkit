# cloudkit-relay

# About

cloudkit-relay is a event/transaction relay batcher. It's goal is to sit between an event/transaction source and buffer collections of events into larger payloads, forwarding them when configured properties are met, such as:

- total sats amount
- event/transaction count
- time elapsed
- data size

The goal is to act as a buffer to reduce the number of broadcast events

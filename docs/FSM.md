# Message Consensus FSM

## block sealer determined

Based on a non network bound algorithm (modulo, hashring), and the last block a node hash, the node calculates the next block numbers sealer.

Basis

- nodes that are determined healthy
- nodes that are considered fully synced
- nodes that are considered sealers

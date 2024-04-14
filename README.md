# clockwise

clockwise is simple. clockwise is ephemeral. a blockchain for decentralized and finalized message passing.

implement a function for service discovery that continuously gets polled.
implement a signing and validation function for communication between nodes.
blocks are signed deterministically by a node, sealer is based on hash of block num -> node.
nodes have three types, archive, signer, and validator.
network config must match for nodes to participate.

## node types

### observers

observer nodes are important for receiving a stream of blocks and performing side effects as a result of which transactions are finalized.

can be configured to have persistence or pruning.

### sealers

sealers take part in the block creation process. they provide new blocks and broadcast newly sealed ones. they also periodically pull in missed blocks from their peers.

can be configured to have persistence or pruning as long as they can still perform sealing duties.

### archivers

archivers always have persistence configured and store records of all blocks. if a sealer or observer needs a record(s) of an old block they under the hood will query their selected archiver.

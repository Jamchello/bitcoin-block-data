# bitcoin-block-data

Simple http server written in golang.

Reads in bitcoin block data from a binary data blob from in a "blk00000.dat" file.

Exposes a /block endpoint which returns the decoded block data, eg:

`http://localhost:8080/block?blockNumber=20`

Returns:

```
{
    "version": 1,
    "previousHeaderHash": "6f187fddd5e28aa1b4065daa5d9eae0c487094fb20cf97ca02b81c8400000000",
    "merkleRootHash": "5b7b25b51797f83192f9fd2c3871bfb27570a7d6b56d3a50760613d1a2fc1aee",
    "timestamp": 1231565995,
    "nbits": 486604799,
    "nonce": 1901123894
}
```

# etherium_proxy

This service is a caching proxy over Cloudflare eth_getBlockByNumber.

Handler `/v1/block/{block_id}` will get block by numerical or hex IDs, also you can use `latest` as block ID 
to get last available block.
_Please note, that this handler will not cache last 25 blocks, since they can change_.

Handler `/v1/block/{block_id}/txs/{tx_id}` will get transaction by block ID(check 1st handler to find what IDs can you use)
and transaction index or hash. You can use hex or numerical numbers to specify index. If block does not contain 
such index or transaction hash, it will return an error message.

If you want to add something to this service, you have to add a new handler to `etherium-proxy.proto` and use 
`make generate` to generate new protobuf files. After that you can modify logic.

Also, it is a good idea to use `make deps` to get all dependencies.
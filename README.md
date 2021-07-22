# eth-mempool-whale-watcher

![Screenshot](https://i.ibb.co/Js6fSFC/2021-07-23-011128-903x402-scrot.png)


listen for large ETH movements in the mempool.

This task will monitor pending transactions and report all mvoements of ETH over a given threshold.
I mostly built this as a way to explore interacting with the mempool through geth but it may be useful as a tool for certain MEV strategies,
or for analytics.

## Usage
### docker
create a .env file and set your NODE_URL to a websocket enabed geth node url. (example: wss://eth-mainnet.alchemyapi.io/)

Then run
```
docker-compose up
```

### local
```
NODE_URL=<websocket geth node url> go run main.go
```



## Contributing

While this is at the present a very simple tool, please feel free to suggest any changes or improvements, or open a PR.

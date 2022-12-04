# grid-logger

POC of a logger with Web3 authentication.

# Logic

```mermaid
sequenceDiagram
    participant client
    participant server
    participant logger
    participant app
    Note over client,app: Authentication logic
    client->>+server: rpc nonce(addr)
    server->>-client: nonce
    client->>client: sig = sign nonce
    client->>+server: rpc signin(addr,nonce,sig)
    server->>-client: token

    Note over client,app: Logging logic
    app->>logger: pipe fifo
    logger->>server: stream rpc write(logname,user,data)
    opt if client logged
        client->>server: [auth] rpc readAndWatch(logname)
        server->>server: tail and follow logs
        server->>client: [auth] stream logs
    end

    Note over client,app: Closing logger/app side logic
    app->>logger: close app
    logger->>server: close write stream
    opt if client logged
        server->>server: tail and follow logs
        server->>client: [auth] close readAndWatch stream
    end

    Note over client,app: Closing client side logic
    opt if client logged
        client->>server: [auth] close readAndWatch stream
        note right of server: Will not close logger stream
    end


```

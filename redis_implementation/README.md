# Redis
## Pre Req
Install Redis Server [Link](https://redis.io/docs/getting-started/installation/install-redis-on-mac-os/)
``` bash
brew install redis
```

## Start/Stop Server
```
brew services start redis
```
```
brew services stop redis
```


```
go get github.com/go-redis/redis/v9
```

## FAQ

Limitation on types available [Link](https://redis.io/docs/manual/data-types/)

Backups [Link](https://redis.io/docs/manual/persistence/)
- Snapshotting
# spookystore
sample fullstack web application hosted on GKE

**disclaimer**: this is not an official Google product, and is intended solely for educational purposes. 

### protoc 

```
protoc -I . ./spookystore.proto --go_out=plugins=grpc:.
```


### Run from source

```
./bin/spookystore --addr=:8001 --google-project-id=spookystore-18
```

```
./web -addr=:8000 --spooky-store-addr=:8001 \
    --google-oauth2-config=/Users/mokeefe/spooky-oauth.json \
    --google-project-id=spookystore-18
``` 


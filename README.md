# spookystore
sample fullstack web application hosted on GKE

![build status](https://travis-ci.org/m-okeefe/spookystore.svg?branch=master)

**disclaimer**: this is not an official Google product, and is intended solely for educational purposes. 

### Run from source

```
./bin/users --addr=:8001 --google-project-id=spookystore-18
```

```
./bin/web -addr=:8000 --user-directory-addr=:8001 \
    --google-oauth2-config=/Users/mokeefe/spooky-oauth.json \
    --google-project-id=spookystore-18
``` 

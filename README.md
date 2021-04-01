[![Build Status](https://cloud.drone.io/api/badges/haibin/hello-service/status.svg)](https://cloud.drone.io/haibin/hello-service)

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/96353eef3eff4d1e918468638ccb4a2f)](https://www.codacy.com/gh/haibin/hello-service/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=haibin/hello-service&amp;utm_campaign=Badge_Grade)

# hello-service

```shell
$ go mod init github.com/haibin/hello-service
```

Install `staticcheck`.

```shell
$ go get honnef.co/go/tools/cmd/staticcheck 
```

## Flow

Request  -> Logger middleware -> Errors middleware -> handler
Response <- Logger middleware <- Errors middleware <- handler 


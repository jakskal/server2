# manual instalation

## Requirement
please make sure you already have:

* Golang installed

*  golang wire

    ```go get github.com/google/wire/cmd/wire```

    and ensuring that $GOPATH/bin is added to your $PATH.

## Setup environtment variable 
you can see the example from .env.example


## Run Application 
```go run cmd/api/main.go cmd/api/wire_gen.go```


# Run using docker

## Requirement
*  docker
see [docker documentation](https://docs.docker.com/) 

## Run application
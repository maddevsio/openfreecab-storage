# Openfreecabs.org storage

A thread-safe in-memory Geo storage for https://openfreecabs.org project based on R-tree index. With simple http api

## Prerequisites

1. [Go](https://golang.org/)
2. [Make](https://www.gnu.org/software/make/)

## Installation

```
mkdir -p $GOPATH/src/github.com/maddevsio/
cd $GOPATH/src/github.com/maddevsio
git clone https://github.com/maddevsio/openfreecab-storage
cd openfreecab-storage
make depends
make
```

Or golang way

```
mkdir -p $GOPATH/src/github.com/maddevsio/
cd $GOPATH/src/github.com/maddevsio
git clone https://github.com/maddevsio/openfreecab-storage
cd openfreecab-storage
go get -v
go build -v
go install
```

## Configure

```
GLOBAL OPTIONS:
   --http_bind_addr value  Define custom http port to bind to (default: ":8090") [$HTTP_BIND_ADDR]
   --base_url value        Define custom base url for project (default: "http://localhost:8090") [$BASE_URL]
   --loglevel value        set log level (default: "debug") [$LOG_LEVEL]
   --test_mode             set test mode [$TEST_MODE]
   --help, -h              show help
   --version, -v           print the version
```

## Run

```
./openfreecab-storage
```

## HTTP API

There are few= http api methods for storage

1. /nearest/:lat/:lon - return k nearest drivers for your location
2. /add/ - Add data from crawler
3. /clean/:companyName/ - purge items for key from storage


## Contributing

Feel free to create issues, sending pull requests.

1. Fork repo
2. Make a changes 
3. Commit
4. Create pull request

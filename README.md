# Getting started with generics - Go / Golang

Tested the first draft of [generics support in Go / Golang](https://go.dev/doc/tutorial/generics).

## Install Golang

For fedora 35
``` bash
sudo dnf -y install golang
```
Or check https://golang.org/doc/install

As of today the support is not in main, so at least `go1.18beta1` is needed or use `gotip` for the latest development version.
``` bash
export GOPATH=~/go
export GOROOT=~/sdk/gotip

go install golang.org/dl/gotip@latest
gotip download  # run again to update

export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# The repo was initialized with `go mod init example/generics` first
```
## Test

```bash
go test -v  # TODO try Fuzzy testing too 
```

## Build and Run

```bash
go run ./main.go

# or:: go build && ./generics
```

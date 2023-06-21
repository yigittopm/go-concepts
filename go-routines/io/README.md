# Go Routines - IO

## Run

To run the example launch:

```bash
go build .
```

```bash
GOMAXPROCS=1 GOGC=off GODEBUG=schedtrace=1000,scheddetail=1 ./io
```
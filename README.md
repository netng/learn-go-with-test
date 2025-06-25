## Run the go file
```
go run .
```

## Run the test
```
go test .
```

## Run the test with bencmark

In windows:
```
go test -bench="." -benchmem
```

In unix (Linux):
```
go test -bench=. -benchmem
```

## Run the test with coverage
```
go test -cover
```
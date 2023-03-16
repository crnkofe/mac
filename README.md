# Simple MAC converter

Convert between numeric and aa:bb:cc:dd:ee:ff form of MAC address.

To build:
```
go build -o mac.go
```

To convert from numeric:
```
./mac 61896753676446                       
38:4b:76:04:00:9e
```

To convert from string form:
```
./mac 38:4b:76:04:00:9e
61896753676446
```

Install RocksDB by compiling. Go inside rocksdb project and run following command

```sh
make static_lib
```

Compile Go Code using following command

```sh
CGO_CFLAGS="-I/home/math2do/Projects/cpp/rocksdb/include" CGO_LDFLAGS="-L/home/math2do/Projects/cpp/rocksdb -lrocksdb -lstdc++ -lm -lz -lsnappy -llz4 -lzstd" go build
```

Use following compatible version: `rocks: 8.9.1` `grocks: 1.8.12`

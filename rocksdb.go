package main

// import (
// 	"log"

// 	"github.com/linxGnu/grocksdb"
// )

// // Build project as follows

// // CGO_CFLAGS="-I/home/math2do/Projects/cpp/rocksdb/include" CGO_LDFLAGS="-L/home/math2do/Projects/cpp/rocksdb -lrocksdb -lstdc++ -lm -lz -lsnappy -llz4 -lzstd" go build

// var (
// 	wo           = grocksdb.NewDefaultWriteOptions()
// 	ro           = grocksdb.NewDefaultReadOptions()
// 	key          = []byte("Key")
// 	value        = []byte("Value")
// 	updatedValue = []byte("UpdatedValue")
// )

// func RocsDBTest() {
// 	db, err := InitDB()
// 	if err != nil {
// 		log.Printf("Open DB Error:%v", err)
// 		return
// 	}
// 	defer db.Close()

// 	// put key value pair
// 	err = db.Put(wo, key, value)
// 	if err != nil {
// 		log.Printf("DB Insert Error:%v", err)
// 		return
// 	}

// 	// retrieve the value for given key
// 	got, err := db.GetBytes(ro, key)
// 	if err != nil {
// 		log.Printf("DB Retrieve Error:%v", err)
// 		return
// 	}

// 	log.Printf("Expected Value:%s, Got:%s", value, string(got))
// 	// update value for existing key
// 	err = db.Put(wo, key, updatedValue)
// 	if err != nil {
// 		log.Printf("DB Insert Error:%v", err)
// 		return
// 	}

// 	// retrieve the value for given key
// 	got, err = db.GetBytes(ro, key)
// 	if err != nil {
// 		log.Printf("DB Retrieve Error:%v", err)
// 		return
// 	}

// 	log.Printf("Expected Value:%s, Got:%s", updatedValue, string(got))
// }

// func InitDB() (*grocksdb.DB, error) {
// 	bbto := grocksdb.NewDefaultBlockBasedTableOptions()
// 	bbto.SetBlockCache(grocksdb.NewLRUCache(3 << 30))
// 	opts := grocksdb.NewDefaultOptions()
// 	opts.SetBlockBasedTableFactory(bbto)
// 	opts.SetCreateIfMissing(true)

// 	db, err := grocksdb.OpenDb(opts, "/home/math2do/Projects/cpp/rocksdb-data/store")
// 	if err != nil {
// 		log.Printf("Open DB Error:%v", err)
// 		return nil, err
// 	}
// 	return db, nil
// }

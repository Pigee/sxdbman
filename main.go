package main

import (
	"fmt"
	badger "github.com/dgraph-io/badger/v4"
	"log"
	"sxdbman/datam"
        "encoding/json"
)

func main() {

	fmt.Println("Hello,DbMan")
	// https://blog.csdn.net/u013317445/article/details/114841853 struct slice
	mc := datam.Conn{"rockymysql", "127.0.0.1", "3306", "sxadmin", "sx@123", "sxdbman", "mysql"}
	//mcl := make([]datam.Conn, 1000)
	mc1 := datam.Conn{"ljpmysql", "127.0.0.1", "3306", "sxadmin", "sx@123", "sxdbman", "mysql"}
	var mcl datam.ConnList
	//append(mcl,myconn)
	mcl[0] = mc
	mcl[1] = mc1
	fmt.Println(mc)
	fmt.Println(mcl)
	fmt.Println(len(mcl))
	fmt.Println(cap(mcl))
    ///////////////////////////JSON/////////////////////////////
     // https://blog.csdn.net/togolife/article/details/121776683
         v, err := json.Marshal(mc)
      if err != nil {
          fmt.Println("marshal failed!", err)
          return
      }
      fmt.Println("marsha result: ", string(v))



	/////////////////badger///////////////////////////////
          // https://juejin.cn/post/6844903814571491335
	db, err := badger.Open(badger.DefaultOptions("./resources/badger"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(mc.Name), []byte(string(v)))
		return err
	})

	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	defer db.Close()
        ////////////////////////////////////////////////////////////
}

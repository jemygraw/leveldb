package cli

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
)

func Export(cmd string, params ...string) {
	if len(params) == 1 {
		dbPath := params[0]
		if _, statErr := os.Stat(dbPath); statErr != nil {
			fmt.Println("LevelDB `", dbPath, "' doesn't exist")
			return
		}
		ldb, lerr := leveldb.OpenFile(dbPath, nil)
		if lerr != nil {
			fmt.Println("Open db error due to", lerr)
			return
		}
		defer ldb.Close()

		iter := ldb.NewIterator(nil, nil)
		for iter.Next() {
			key := string(iter.Key())
			value := string(iter.Value())
			fmt.Println(fmt.Sprintf("%s\t%s", key, value))
		}
		iter.Release()
		ierr := iter.Error()
		if ierr != nil {
			fmt.Println("Export data error due to", ierr)
		}
		return
	} else {
		Help(cmd)
	}
}

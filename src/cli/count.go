package cli

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
)

func Count(cmd string, params ...string) {
	if len(params) == 1 {
		dbPath := params[0]
		if _, statErr := os.Stat(dbPath); statErr != nil {
			fmt.Println("LevelDB ", dbPath, " doesn't exist")
			return
		}
		ldb, lerr := leveldb.OpenFile(dbPath, nil)
		if lerr != nil {
			fmt.Println("Open db error due to", lerr)
			return
		}
		defer ldb.Close()
		totalCnt := 0
		iter := ldb.NewIterator(nil, nil)
		for iter.Next() {
			totalCnt += 1
		}
		iter.Release()
		ierr := iter.Error()
		if ierr != nil {
			fmt.Println("Export data error due to", ierr)
		}
		fmt.Println(totalCnt)
		return
	} else {
		Help(cmd)
	}
}

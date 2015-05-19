package cli

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
)

//leveldb create -path=/Users/jemy/Db/test.ldb
func Create(cmd string, params ...string) {
	if len(params) == 1 {
		path := params[0]
		if _, statErr := os.Stat(path); statErr == nil {
			fmt.Println("LevelDB ", path, " exists, plz delete it first!")
			return
		}
		ldb, err := leveldb.OpenFile(path, nil)
		if err != nil {
			fmt.Println("Create leveldb failed due to", err)
		}
		defer ldb.Close()
	} else {
		Help(cmd)
	}
}

package cli

import (
	"bufio"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"strings"
)

func Import(cmd string, params ...string) {
	if len(params) == 2 {
		dbPath := params[0]
		dataPath := params[1]
		if _, statErr := os.Stat(dbPath); statErr != nil {
			fmt.Println("LevelDB ", dbPath, " doesn't exist")
			return
		}
		if _, statErr := os.Stat(dataPath); statErr != nil {
			fmt.Println("Data file ", dataPath, " doesn't exist")
			return
		}
		ldb, lerr := leveldb.OpenFile(dbPath, nil)
		if lerr != nil {
			fmt.Println("Open db error due to", lerr)
			return
		}
		defer ldb.Close()
		dataFp, openErr := os.Open(dataPath)
		if openErr != nil {
			fmt.Println("Open data file failed due to", openErr)
			return
		}
		defer dataFp.Close()
		bReader := bufio.NewScanner(dataFp)
		bReader.Split(bufio.ScanLines)
		for bReader.Scan() {
			line := strings.TrimSpace(bReader.Text())
			items := strings.Split(line, "\t")
			if len(items) == 2 {
				key := items[0]
				value := items[1]
				ldb.Put([]byte(key), []byte(value), nil)
			}
		}
	} else {
		Help(cmd)
	}
}

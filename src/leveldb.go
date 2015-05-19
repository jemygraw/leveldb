package main

import (
	"cli"
	"flag"
)

//leveldb -create=/home/jemy/test.ldb
//leveldb -import=/home/jemy/test.ldb -data=data.txt
//leveldb -export=/home/jemy/test.ldb
func main() {
	var lcreate string

	var limport string
	var limportData string

	var lexport string

	var lcount string
	flag.StringVar(&lcreate, "create", "", "leveldb path")
	flag.StringVar(&limport, "import", "", "leveldb path")
	flag.StringVar(&limportData, "data", "", "data to import")
	flag.StringVar(&lexport, "export", "", "leveldb path")
	flag.StringVar(&lcount, "count", "", "leveldb path")

	flag.Usage = func() {
		cli.Help()
	}

	flag.Parse()

	if lcreate != "" {
		cli.Create("create", []string{lcreate}...)
		return
	} else if limport != "" {
		cli.Import("import", []string{limport, limportData}...)
		return
	} else if lexport != "" {
		cli.Export("export", []string{lexport}...)
		return
	} else if lcount != "" {
		cli.Count("count", []string{lcount}...)
	}
}

package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/pingcap/parser"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"vitess.io/vitess/go/vt/sqlparser"
)

func main() {
	f, err := os.Open("example/migrations/20200420143123_create_users_table.up.sql")
	if err != nil {
		log.Fatalln(err)
	}
	tz := sqlparser.NewTokenizer(f)
	for {
		ddlStmt, err := sqlparser.ParseNextStrictDDL(tz)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatalln(err)
		}
		spew.Dump(ddlStmt)
	}
}

func tidbParser(s string) {
	p := parser.New()
	s2, w, e := p.Parse(s, "", "")
	spew.Dump(s2, w, e)
}

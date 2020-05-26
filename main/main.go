package main

import (
	"fmt"
	pg_query "github.com/lfittl/pg_query_go"
)

func main() {
	stmt, err := pg_query.Parse("SELECT a from b group by grouping sets (a,b) union select a from d union select a from l")
	if err != nil {
		panic(err)
	}

	data, err := pg_query.Deparse(stmt)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	return

}

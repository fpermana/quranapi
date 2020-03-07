package main

import (
	"fmt"
	// "net/http"
	"log"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/fpermana/quranapi/paging"
	"github.com/fpermana/quranapi/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/quran")
	// db, err := sql.Open("mysql", "root:@/quran")
	if err != nil {
		log.Fatal(err)
	}

	/*fmt.Println(db)*/

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/greet/"):]
		fmt.Fprintf(w, "Hello %s\n", name)
	})

	http.ListenAndServe("localhost:9870", nil)*/
	// Setup repositories
	var (
		ayas         paging.AyaRepository
	)

	ayas, _ = mysql.NewAyaRepository(db)
	n,m := ayas.GetSuraAyaStart(2)
	o := ayas.GetNumberBySuraAya("quran_text", 2,1)
	p := ayas.GetAya("quran_text", "en_sahih", 1)

	fmt.Println("HELLO WORLD")
	fmt.Println(n,m)
	fmt.Println(o)
	fmt.Println(p)
}

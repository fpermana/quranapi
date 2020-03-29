package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	//"context"
	//"flag"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-kit/kit/log"
	"github.com/fpermana/quranapi/paging"
	"github.com/fpermana/quranapi/searching"
	"github.com/fpermana/quranapi/quran"
	"github.com/fpermana/quranapi/server"
	"github.com/fpermana/quranapi/mysql"
)

const (
	defaultRoutingServiceURL = "http://localhost:7878"
	defaultMysqlURL          = "localhost"
	defaultMysqlPort         = "3306"
	defaultMysqlUser         = "root"
	defaultMysqlPassword     = "password"
	defaultDBName            = "quran"
)

func main() {
	var (
		connection = fmt.Sprintf("%s:%s@tcp(%s)/%s",defaultMysqlUser,defaultMysqlPassword,defaultMysqlURL,defaultDBName)
	)

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		//log.Fatal(err)
	}
	defer db.Close()

	// Setup repositories
	var (
		ayas		paging.AyaRepository
		pages		paging.PageRepository
		translations	quran.TranslationRepository
		suras		quran.SuraRepository
		searchings	searching.SearchingRepository
	)

	ayas, _ = mysql.NewAyaRepository(db)
	pages, _ = mysql.NewPageRepository(db)
	translations, _ = mysql.NewTranslationRepository(db)
	suras, _ = mysql.NewSuraRepository(db)
	searchings, _ = mysql.NewSearchingRepository(db)

	var ps paging.Service
	ps = paging.NewService(ayas,pages)
	ps = paging.NewLoggingService(logger,ps)

	var qs quran.Service
	qs = quran.NewService(translations, suras)
	qs = quran.NewLoggingService(logger, qs)

	var ss searching.Service
	ss = searching.NewService(searchings)
	ss = searching.NewLoggingService(logger, ss)

	httpAddr := ":7878"

	srv := server.New(ps, qs, ss, log.With(logger, "component", "http"))
        errs := make(chan error, 2)
        go func() {
		logger.Log("transport", "http", "address", httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(httpAddr, srv)
	}()
        go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
        logger.Log("terminated", <-errs)
}

package mysql

import (
	"github.com/fpermana/quranapi/paging"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ayaRepository struct {
	db      *sql.DB
}

func (r *ayaRepository) GetSuraAyaStart(pageNumber int) (int,int) {

	var sura, aya int
	// Execute the query
	err := r.db.QueryRow("SELECT sura, aya FROM pages WHERE id = ?", pageNumber).Scan(&sura, &aya)
	if err != nil {
	    // panic(err.Error()) // proper error handling instead of panic in your app
	    return -1,-1
	}

	return sura, aya
}

func (r *ayaRepository) GetNumberBySuraAya(quran_text string, suraNumber int, ayaNumber int) paging.AyaNumber {
	var number int
	// Execute the query
	err := r.db.QueryRow("SELECT id FROM "+quran_text+" WHERE sura = ? AND aya = ?", suraNumber, ayaNumber).Scan(&number)
	if err != nil {
	    // panic(err.Error()) // proper error handling instead of panic in your app
	    return paging.AyaNumber(-1)
	}

	return paging.AyaNumber(number)
}

func (r *ayaRepository) GetAya(quran_text string, translation string, number paging.AyaNumber) *paging.AyaModel {
	var aya *paging.AyaModel

	rows, err := r.db.Query("SELECT "+quran_text+".id, "+quran_text+".text, "+translation+".text, "+quran_text+".sura, "+quran_text+".aya FROM "+quran_text+" JOIN "+translation+" ON "+translation+".id = "+quran_text+".id WHERE "+quran_text+".id = ?", number)
	if err != nil {
	    // panic(err.Error()) // proper error handling instead of panic in your app
	    return aya
	}
	aya = new(paging.AyaModel)
	for rows.Next() {
		if err := rows.Scan(&aya.Number, &aya.Text, &aya.Translation, &aya.Sura, &aya.Aya); err != nil {
			return nil
		}
	}

	return aya
}

/*func (r *ayaRepository) GetAyaListFrom(quran_text string, translation string, number paging.AyaNumber) []*paging.AyaModel {
}

func (r *ayaRepository) GetAyaListBetween(quran_text string, translation string, firstId int, secondId int, inclusive bool) []*paging.AyaModel {
}*/

// NewAyaRepository returns a new instance of a MySQL aya repository.
func NewAyaRepository(db *sql.DB) (paging.AyaRepository, error) {
	r := &ayaRepository{
		db:      db,
	}

	return r, nil
}
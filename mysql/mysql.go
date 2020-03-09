package mysql

import (
	"fmt"
	"github.com/fpermana/quranapi/paging"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ayaRepository struct {
	db      *sql.DB
}

func (r *ayaRepository) GetSuraAyaStart(pageNumber int) (int,int) {

	var sura, aya int
	var query string = fmt.Sprintf("SELECT sura, aya FROM pages WHERE id = ?")
	// Execute the query
	rows, err := r.db.Query(query, pageNumber)
	if err != nil {
	    // panic(err.Error()) // proper error handling instead of panic in your app
	    return -1,-1
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&sura, &aya); err != nil {
			return -1,-1
		}
	}

	return sura, aya
}

func (r *ayaRepository) GetNumberBySuraAya(quran_text string, suraNumber int, ayaNumber int) paging.AyaNumber {
	var number int
	var query string = fmt.Sprintf("SELECT id FROM %[1]s WHERE sura = ? AND aya = ?", quran_text)
	// Execute the query
	rows, err := r.db.Query(query, suraNumber, ayaNumber)
	if err != nil {
	    // panic(err.Error()) // proper error handling instead of panic in your app
	    return paging.AyaNumber(-1)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&number); err != nil {
			return paging.AyaNumber(-1)
		}
	}

	return paging.AyaNumber(number)
}

func (r *ayaRepository) GetAya(quran_text string, translation string, number paging.AyaNumber) *paging.AyaModel {
	var model *paging.AyaModel

	var query string = fmt.Sprintf("SELECT %[1]s.id, %[1]s.text, %[2]s.text, %[1]s.sura, %[1]s.aya FROM %[1]s JOIN %[2]s ON %[2]s.id = %[1]s.id WHERE %[1]s.id = ?", quran_text, translation)
	rows, err := r.db.Query(query, number)
	if err != nil {
	   // panic(err.Error()) // proper error handling instead of panic in your app
	    return nil
	}
	defer rows.Close()

	model = new(paging.AyaModel)
	for rows.Next() {
		var number, sura, aya int
		var text, translation string
		if err := rows.Scan(&number, &text, &translation, &sura, &aya); err != nil {
			return nil
		}
		model.Number = paging.AyaNumber(number)
		model.Text = text
		model.Translation = translation
		model.Sura = sura
		model.Aya = aya
	}

	return model
}

func (r *ayaRepository) GetAyaListFrom(quran_text string, translation string, ayaNumber paging.AyaNumber) []*paging.AyaModel {
	var ayaList []*paging.AyaModel

	var query string = fmt.Sprintf("SELECT %[1]s.id, %[1]s.text, %[2]s.text, %[1]s.sura, %[1]s.aya, suras.name FROM %[1]s JOIN  %[2]s ON %[1]s.id = %[2]s.id JOIN suras ON suras.id = %[1]s.sura WHERE %[1]s.id >= ?", quran_text, translation)
	// fmt.Println(query, ayaNumber)
	// Execute the query
	rows, err := r.db.Query(query,ayaNumber)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		var model *paging.AyaModel = new(paging.AyaModel)
		var number, sura, aya int
		var text, translation, suraName string
		if err := rows.Scan(&number, &text, &translation, &sura, &aya, &suraName); err != nil {
			return nil
		}
		model.Number = paging.AyaNumber(number)
		model.Text = text
		model.Translation = translation
		model.Sura = sura
		model.Aya = aya
		model.SuraName = suraName
		ayaList = append(ayaList, model)
        }
	return ayaList
}

func (r *ayaRepository) GetAyaListBetween(quran_text string, translation string, firstId int, secondId int, inclusive bool) []*paging.AyaModel {
	var ayaList []*paging.AyaModel

	var sign string = "<"
	if inclusive {
		sign = "<="
	}
	var query string = fmt.Sprintf("SELECT %[1]s.id, %[1]s.text, %[2]s.text, %[1]s.sura, %[1]s.aya, suras.name FROM %[1]s JOIN  %[2]s ON %[1]s.id = %[2]s.id JOIN suras ON suras.id = %[1]s.sura WHERE %[1]s.id >= ? AND %[1]s.id %[3]s ?", quran_text, translation, sign)
	// fmt.Println(query, firstId, secondId)
	// Execute the query
	rows, err := r.db.Query(query,firstId,secondId)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		var model *paging.AyaModel = new(paging.AyaModel)
		var number, sura, aya int
		var text, translation, suraName string
		if err := rows.Scan(&number, &text, &translation, &sura, &aya, &suraName); err != nil {
			return nil
		}
		model.Number = paging.AyaNumber(number)
		model.Text = text
		model.Translation = translation
		model.Sura = sura
		model.Aya = aya
		model.SuraName = suraName
		ayaList = append(ayaList, model)
        }
	return ayaList
}

// NewAyaRepository returns a new instance of a MySQL aya repository.
func NewAyaRepository(db *sql.DB) (paging.AyaRepository, error) {
	r := &ayaRepository{
		db:      db,
	}

	return r, nil
}

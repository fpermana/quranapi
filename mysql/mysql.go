package mysql

import (
	"fmt"
	"strconv"
	"github.com/fpermana/quranapi/paging"
	"github.com/fpermana/quranapi/searching"
	"github.com/fpermana/quranapi/quran"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ayaRepository struct {
	db      *sql.DB
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
		var number, sura, aya int
		var text, translation, suraName string
		if err := rows.Scan(&number, &text, &translation, &sura, &aya, &suraName); err != nil {
			return nil
		}
		var model *paging.AyaModel = new(paging.AyaModel)
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

func (r *ayaRepository) GetAyaListBetween(quran_text string, translation string, firstId paging.AyaNumber, secondId paging.AyaNumber, inclusive bool) []*paging.AyaModel {
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
		var number, sura, aya int
		var text, translation, suraName string
		if err := rows.Scan(&number, &text, &translation, &sura, &aya, &suraName); err != nil {
			return nil
		}
		var model *paging.AyaModel = new(paging.AyaModel)
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

type pageRepository struct {
	db      *sql.DB
}

func (r *pageRepository) GetSuraAyaStart(pageNumber int) (int,int) {

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

func (r *pageRepository) GetTotalPage() int {
	var pages int = 0
	_ = r.db.QueryRow("SELECT count(*) FROM pages").Scan(&pages)

	return pages;
}

// NewPageRepository returns a new instance of a MySQL page repository.
func NewPageRepository(db *sql.DB) (paging.PageRepository, error) {
	r := &pageRepository{
		db:      db,
	}

	return r, nil
}

type translationRepository struct {
	db      *sql.DB
}

func (r *translationRepository) GetTranslationList() []*quran.TranslationModel {
	var translationList []*quran.TranslationModel

	var query string = fmt.Sprintf("SELECT id, flag, lang, name, translator, tid, installed, is_default, visible, iso6391 FROM translations")
	// fmt.Println(query, ayaNumber)
	// Execute the query
	rows, err := r.db.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		var number int
		var flag, lang, name, translator, tid, iso6391 string
		var installed, is_default, visible bool
		if err := rows.Scan(&number, &flag, &lang, &name, &translator, &tid, &installed, &is_default, &visible, &iso6391); err != nil {
			return nil
		}
		var model *quran.TranslationModel = new(quran.TranslationModel)
		model.Number = quran.TranslationNumber(number)
		model.Flag = flag
		model.Lang = lang
		model.Name = name
		model.Translator = translator
		model.Tid = tid
		model.Installed = installed
		model.IsDefault = is_default
		model.Visible = visible
		model.Iso6391 = iso6391
		translationList = append(translationList, model)
        }
	return translationList
}

// NewTranslationRepository returns a new instance of a MySQL translation repository.
func NewTranslationRepository(db *sql.DB) (quran.TranslationRepository, error) {
	r := &translationRepository{
		db:      db,
	}

	return r, nil
}

type suraRepository struct {
	db      *sql.DB
}

func (r *suraRepository) GetSuraList() []*quran.SuraModel {
	var suraList []*quran.SuraModel

	var query string = fmt.Sprintf("SELECT id, ayas, start, name, tname, ename, type, `order`, rukus FROM suras")
	// fmt.Println(query, ayaNumber)
	// Execute the query
	rows, err := r.db.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		var number, ayas, start, stype,  order, rukus int
		var name, tname, ename string
		if err := rows.Scan(&number, &ayas, &start, &name, &tname, &ename, &stype, &order, &rukus); err != nil {
			return nil
		}
		var model *quran.SuraModel = new(quran.SuraModel)
		model.Number = quran.SuraNumber(number)
		model.Ayas = ayas
		model.Start = start
		model.Name = name
		model.TName = tname
		model.EName = ename
		model.Type = stype
		model.Order = order
		model.Rukus = rukus
		suraList = append(suraList, model)
        }
	return suraList
}

func (r *suraRepository) GetSuraPage(suraNumber quran.SuraNumber) int {
	var page, sura, aya int
	var query string = fmt.Sprintf("SELECT id,sura,aya FROM pages WHERE sura >= ? ORDER BY sura ASC LIMIT 1")
	//fmt.Println(query, int(suraNumber))
	// Execute the query
	var oSura = int(suraNumber)
	rows, err := r.db.Query(query,oSura)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&page, &sura, &aya); err != nil {
            if sura > oSura || (sura == oSura && aya > 1) {
                page--
            }
		}
	}

	return page
}
// NewSuraRepository returns a new instance of a MySQL sura repository.
func NewSuraRepository(db *sql.DB) (quran.SuraRepository, error) {
	r := &suraRepository{
		db:      db,
	}

	return r, nil
}

type searchingRepository struct {
	db      *sql.DB
}

func (r *searchingRepository) Search(keywords string, quran_text string, translation string, lastId int, limit int) []*paging.AyaModel {
	var ayaList []*paging.AyaModel

	var query string = fmt.Sprintf("SELECT %[1]s.id, %[1]s.text, %[2]s.text, %[1]s.sura, %[1]s.aya, suras.name FROM %[1]s JOIN %[2]s ON %[2]s.id = %[1]s.id JOIN suras ON suras.id = %[1]s.sura WHERE %[2]s.text LIKE ? AND %[1]s.id > ? ORDER BY %[1]s.id LIMIT %[3]d", quran_text, translation, limit)
	//fmt.Println(query)
	stacks := []string{fmt.Sprintf("%%%[1]s%%",keywords), strconv.Itoa(lastId)}
	args := make([]interface{}, len(stacks))
	for i := range stacks {
		args[i] = stacks[i]
	}

	// Execute the query
	rows, err := r.db.Query(query, args...)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()
	for rows.Next() {
		var number, sura, aya int
		var text, translation, suraName string
		if err := rows.Scan(&number, &text, &translation, &sura, &aya, &suraName); err != nil {
			return nil
		}
		var model *paging.AyaModel = new(paging.AyaModel)
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

// NewSearchingRepository returns a new instance of a MySQL searching repository.
func NewSearchingRepository(db *sql.DB) (searching.SearchingRepository, error) {
	r := &searchingRepository{
		db:      db,
	}

	return r, nil
}


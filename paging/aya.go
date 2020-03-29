package paging

// Number uniquely identifies a particular aya.
type AyaNumber int

// Aya is the central class in the domain model.
type AyaModel struct {
	Number			AyaNumber	`json:"id"`
	Text			string		`json:"text"`
	Translation		string		`json:"translation"`
	Aya			int		`json:"aya"`
	Sura			int		`json:"sura"`
	SuraName		string		`json:"sura_name"`
	Marked			bool		`json:"marked"`
}

// AyaRepository provides access a aya store.
type AyaRepository interface {
    GetNumberBySuraAya(quran_text string, suraNumber int, ayaNumber int) AyaNumber; // get number of specified aya and sura
    GetAya(quran_text string, translation string, number AyaNumber) *AyaModel; // get ayamodel by number
    GetAyaListFrom(quran_text string, translation string, number AyaNumber) []*AyaModel; // get aya list from number
    GetAyaListBetween(quran_text string, translation string, firstId AyaNumber, secondId AyaNumber, inclusive bool) []*AyaModel; // get aya list between 2 numbers
}

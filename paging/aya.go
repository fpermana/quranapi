package paging

// Number uniquely identifies a particular aya.
type AyaNumber int

// Aya is the central class in the domain model.
type AyaModel struct {
	Number			AyaNumber
	Text			string
	Translation		string
	Aya				int
	Sura			int
	SuraName		string
	Marked			bool
}

// AyaRepository provides access a aya store.
type AyaRepository interface {
    GetSuraAyaStart(pageNumber int) (int,int); // get first aya and sura for specified page
    GetNumberBySuraAya(quran_text string, suraNumber int, ayaNumber int) AyaNumber; // get number of specified aya and sura
    GetAya(quran_text string, translation string, number AyaNumber) *AyaModel; // get ayamodel by number
    /*GetAyaListFrom(quran_text string, translation string, number AyaNumber) []*AyaModel; // get aya list from number
    GetAyaListBetween(quran_text string, translation string, firstId int, secondId int, inclusive bool) []*AyaModel; // get aya list between 2 numbers*/
}

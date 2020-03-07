package quran

// Number uniquely identifies a particular sura.
type SuraNumber int

// Aya is the central class in the domain model.
type SuraModel struct {
	Number			SuraNumber
	Ayas			int
	Start			int
	Name			string
	TName			string
	EName			string
	Type			string
	Order			int
	Rukus			int
}

// SuraRepository provides access a sura store.
type SuraRepository interface {
    GetSuraList() []*SuraModel; // get all sura list
    GetSuraPage(int SuraNumber) int; //get page where sura started
}

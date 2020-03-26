package quran

// Number uniquely identifies a particular sura.
type SuraNumber int

// Aya is the central class in the domain model.
type SuraModel struct {
	Number			SuraNumber	`json:"id"`
	Ayas			int		`json:"ayas"`
	Start			int		`json:"start"`
	Name			string		`json:"name"`
	TName			string		`json:"tname"`
	EName			string		`json:"ename"`
	Type			int		`json:"type"`
	Order			int		`json:"order"`
	Rukus			int		`json:"rukus"`
}

// SuraRepository provides access a sura store.
type SuraRepository interface {
    GetSuraList() []*SuraModel; // get all sura list
    GetSuraPage(suraNumber SuraNumber) (int,int,int); //get page where sura started
}

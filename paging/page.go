package paging

// Number uniquely identifies a particular page.
type PageNumber int

// Page is the central class in the domain model.
type PageModel struct {
	Number			PageNumber	`json:"id"`
	Sura			int		`json:"sura"`
	Aya			int		`json:"aya"`
}

// PageRepository provides access a aya store.
type PageRepository interface {
    GetSuraAyaStart(pageNumber int) (int,int); // get first aya and sura for specified page
    GetTotalPage() int; // get total page
}

package quran

// Number uniquely identifies a particular translation.
type TranslationNumber int

// Translation is the central class in the domain model.
type TranslationModel struct {
	Number			TranslationNumber	`json:"id"`
	Flag			string			`json:"flag"`
	Lang			string			`json:"lang"`
	Name			string			`json:"name"`
	Translator		string			`json:"translator"`
	Tid			string			`json:"tid"`
	Installed		bool			`json:"installed"`
	IsDefault		bool			`json:"is_default"`
	Visible			bool			`json:"visible"`
	Iso6391			string			`json:"iso6391"`
}

// TranslationRepository provides access a translation store.
type TranslationRepository interface {
    GetTranslationList() []*TranslationModel; // get all translation list
}



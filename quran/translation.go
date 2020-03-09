package quran

// Number uniquely identifies a particular translation.
type TranslationNumber int

// Translation is the central class in the domain model.
type TranslationModel struct {
	Number			TranslationNumber
	Flag			string
	Lang			string
	Name			string
	Translator		string
	Tid			string
	Installed		bool
	IsDefault		bool
	Visible			bool
	Iso6391			string
}

// TranslationRepository provides access a translation store.
type TranslationRepository interface {
    GetTranslationList() []*TranslationModel; // get all translation list
}



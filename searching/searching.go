package searching

import (
	"github.com/fpermana/quranapi/paging"
)

// SearchingRepository provides access a aya store.
type SearchingRepository interface {
    Search(keywords string, quran_text string, translation string, lastId int, limit int) []*paging.AyaModel; // get aya list that contain keywords from translation
}

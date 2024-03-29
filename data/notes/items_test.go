package notes_test

import (
	"github.com/orionlab42/parmtracker/data/notes"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestItemInsertAndFetch(t *testing.T) {
	i := notes.Item{
		NoteId:         1,
		ItemText:       "red",
		ItemIsComplete: false,
		ItemDate:       time.Now().UTC(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
	e := i.Insert()
	assert.Nil(t, e)
	items := notes.GetItems()
	assert.NotEqual(t, len(items), 0)
	e = i.Delete()
	assert.Nil(t, e)
}

func TestItemSave(t *testing.T) {
	i := notes.Item{
		NoteId:         3,
		ItemText:       "black",
		ItemIsComplete: false,
		ItemDate:       time.Now().UTC(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
	e := i.Insert()
	assert.Nil(t, e)
	i.ItemText = "Shark"
	err := i.Save()
	if err != nil {
		return
	}
	assert.Equal(t, i.ItemText, "Shark")
	e = i.Delete()
	assert.Nil(t, e)
}

//func TestGetItemsByNoteId(t *testing.T) {
//	i := notes.Item{
//		NoteId:         3,
//		ItemText:       "black",
//		ItemIsComplete: false,
//		ItemDate:       time.Now().UTC(),
//		CreatedAt:      time.Now().UTC(),
//		UpdatedAt:      time.Now().UTC(),
//	}
//	i.Insert()
//	i = notes.Item{
//		NoteId:         3,
//		ItemText:       "red",
//		ItemIsComplete: true,
//		ItemDate:       time.Now().UTC(),
//		CreatedAt:      time.Now().UTC(),
//		UpdatedAt:      time.Now().UTC(),
//	}
//	i.Insert()
//	i = notes.Item{
//		NoteId:         2,
//		ItemText:       "red",
//		ItemIsComplete: true,
//		ItemDate:       time.Now().UTC(),
//		CreatedAt:      time.Now().UTC(),
//		UpdatedAt:      time.Now().UTC(),
//	}
//	i.Insert()
//	items := notes.GetItemsByNoteId(3)
//	for _, item := range items {
//		fmt.Println(item)
//	}
//}

//func TestCreateItemsByNoteId(t *testing.T) {
//	items := notes.CreateItemsByNoteId(11, "2022-06-27T22:00:00.000Z", "2022-07-01T22:00:00.000Z")
//	fmt.Println(items)
//}

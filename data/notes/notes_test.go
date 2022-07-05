package notes_test

import (
	"github.com/orionlab42/parmtracker/data/notes"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNoteInsertAndFetch(t *testing.T) {
	n := notes.Note{
		UserId:    1,
		NoteType:  2,
		NoteTitle: "Shark",
		NoteText:  "red",
		NoteEmpty: false,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	e := n.Insert()
	assert.Nil(t, e)
	notes := notes.GetNotes()
	assert.NotEqual(t, len(notes), 0)
	e = n.Delete()
	assert.Nil(t, e)
}

func TestNoteSave(t *testing.T) {
	n := notes.Note{
		UserId:    1,
		NoteType:  2,
		NoteTitle: "Shark",
		NoteText:  "red",
		NoteEmpty: false,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	e := n.Insert()
	assert.Nil(t, e)
	n.NoteTitle = "Not Shark"
	n.NoteText = "two"
	err := n.Save()
	if err != nil {
		return
	}
	assert.Equal(t, n.NoteTitle, "Not Shark")
	e = n.Delete()
	assert.Nil(t, e)
}

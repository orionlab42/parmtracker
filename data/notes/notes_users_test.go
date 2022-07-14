package notes_test

import (
	"github.com/orionlab42/parmtracker/data/notes"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNoteUsersInsertAndFetch(t *testing.T) {
	n := notes.NoteUser{
		NoteId:    1,
		UserId:    2,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	e := n.Insert()
	assert.Nil(t, e)
	noteUsers := notes.GetNotesUsers()
	assert.NotEqual(t, len(noteUsers), 0)
	e = n.Delete()
	assert.Nil(t, e)
}

func TestNoteUsersSave(t *testing.T) {
	n := notes.NoteUser{
		NoteId:    1,
		UserId:    2,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	e := n.Insert()
	assert.Nil(t, e)
	n.UserId = 102
	err := n.Save()
	if err != nil {
		return
	}
	//notes := notes.GetNotesUsers()
	//fmt.Println(notes)
	assert.Equal(t, n.UserId, 102)
	e = n.Delete()
	assert.Nil(t, e)
}

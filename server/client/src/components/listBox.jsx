import React, {useEffect, useState} from "react";
import Note from "./common/note";
import CheckList from "./common/checkList";
import AgendaNote from "./common/agendaNote";
// import {searchKeywordNotes} from "../utils/search";
// import SearchBox from "./searchBox";


const ListBox = (props) => {
    const [notes, setNotes] = useState([]);
    // const [searchQuery, setSearchQuery] = useState("");
    const [id, setId] = useState(1);

    // Sorts the notes by date even while a date is edited which is annoying
    // const sortedNotes = notes.sort((a,b) => b.date - a.date);

    const giveId = () => {
        setId(id + 1);
        return id;
    };

    useEffect(() => {
        let savedNotes = JSON.parse(localStorage.getItem('react-notes-app-data'));
        if (savedNotes) {
            setNotes(savedNotes);
            if (savedNotes[0]) {
                setId(savedNotes[0].id + 1);
            }
        }
    }, []);

    useEffect(() => {
        localStorage.setItem('react-notes-app-data', JSON.stringify(notes));
    }, [notes]);

    const addNote = () => {
        let newNote = {
            id: giveId(),
            type: "simple-note",
            empty: true,
            title: "",
            text: "",
            date: Date.now()
        };
        setNotes([newNote, ...notes]);
    };

    const addCheckNote = () => {
        let newList = {
            id: giveId(),
            type: "checklist",
            empty: true,
            title: "",
            list: [],
            date: Date.now()
        }
        setNotes([newList, ...notes]);
    };

    const addAgendaNote = () => {
        let newAgenda = {
            id: giveId(),
            type: "agenda",
            empty: true,
            // list: [],
            title: "",
            date: Date.now()
        }
        setNotes([newAgenda, ...notes]);
    };

    const updateNote = (newNote) => {
        let newNotes = notes.map(note => {
            if (note.id === newNote.id) {
                note.empty = false;
                note.title = newNote.title;
                note.text = newNote.text;
                note.date = Date.now();
            }
            return note
        });
        setNotes(newNotes);
    };

    const deleteNote = (id) => {
        let newNotes = notes.filter(note => note.id !== id);
        setNotes(newNotes);
    };

    // const searchNote = (text) => {
    //     setSearchQuery(text);
    // };

    // let notesToDisplay = notes;
    // if (searchQuery) {
    //     notesToDisplay = searchKeywordNotes(notes, searchQuery);
    // }
    //
    // console.log("All notes", notes);
    return (
        <div className="notes-list-container">
            {/*<SearchBox value={searchQuery} onChange={searchNote}/>*/}
            <button className="button is-link is-light add-note-button" onClick={addNote}><span
                className="mdi mdi-note-outline"/> &nbsp; Add Simple Note</button>
            <button className="button is-link is-light add-note-button" onClick={addCheckNote}><span
                className="mdi mdi-playlist-check"/> &nbsp; Add Checklist</button>
            <button className="button is-link is-light add-note-button" onClick={addAgendaNote}><span
                className="mdi mdi-calendar-text"/> &nbsp; Add Planner</button>
            <div className="notes-list">
                {notes.map(note => {
                    if (note.type === "simple-note") {
                         return <Note key={note.id}
                                      note={ note }
                                      handleUpdateNote={updateNote}
                                      handleDeleteNote={deleteNote}/>
                    }
                    if (note.type === "checklist") {
                        return <CheckList key={note.id}
                                            items={note}
                                            handleUpdateCheckList={updateNote}
                                            handleDeleteCheckList={deleteNote}/>
                    }
                    if (note.type === "agenda") {
                        return <AgendaNote key={note.id}
                                           items={note}
                                           handleUpdateAgendaNote={updateNote}
                                           handleDeleteAgenda={deleteNote}
                        />
                    }
                    return null
                })}
            </div>

        </div>
    );
};

export default ListBox;
import React, {useEffect, useState} from "react";
import Note from "./common/note";
import CheckList from "./common/checkList";
import {searchKeywordNotes} from "../utils/search";
import SearchBox from "./searchBox";

const ListBox = (props) => {
    const [notes, setNotes] = useState([]);
    // const [searchQuery, setSearchQuery] = useState("");
    const [id, setId] = useState(1);

    const giveId = () => {
        setId(id + 1);
        return id;
    };

    useEffect(() => {
        const savedNotes = JSON.parse(localStorage.getItem('react-notes-app-data'));
        if (savedNotes) {
            setNotes(savedNotes);
        }
    }, []);

    useEffect(() => {
        localStorage.setItem('react-notes-app-data', JSON.stringify(notes));
    }, [notes]);

    const addNote = (text) => {
        const date = new Date();
        const newNote = {
            id: giveId(),
            type: "simple-note",
            title: "",
            text: "",
            date: date.toLocaleDateString()
        };
        setNotes([...notes, newNote]);
    };

    const updateNote = (newNote) => {
        const newNotes = notes.map(note => {
            if (note.id === newNote.id) {
                note.text = newNote.text;
                note.title = newNote.title;
            }
            return note
        });
        setNotes(newNotes);
    };

    const deleteNote = (deleteNote) => {
        const newNotes = notes.filter(note => note.id !== deleteNote.id);
        setNotes(newNotes);
    };

    const addCheckNote = (itemList) => {
        const date = new Date();
        let newList = {
            id: giveId(),
            type: "checklist",
            title: "",
            list: [],
            date: date.toLocaleDateString()
        }
        setNotes([...notes, newList]);
    };

    const updateCheckList = (itemList) => {
        const newChecklists = notes.map(checkList => {
            if (checkList.id === itemList.id) {
                checkList.list = itemList.list
            }
            return checkList
        });
        setNotes(newChecklists);
    };

    const deleteCheckList = (id) => {
        const newChecklists = notes.filter(checkList => checkList.id !== id);
        setNotes(newChecklists);
    };

    // const searchNote = (text) => {
    //     setSearchQuery(text);
    // };

    // let notesToDisplay = notes;
    // if (searchQuery) {
    //     notesToDisplay = searchKeywordNotes(notes, searchQuery);
    // }

    console.log("All notes", notes);
    return (
        <div className="notes-list-container">
            {/*<SearchBox value={searchQuery} onChange={searchNote}/>*/}
            <button className="button is-link is-light add-note-button" onClick={addNote}><span
                className="mdi mdi-note-outline"/> &nbsp; Add Simple Note</button>
            <button className="button is-link is-light add-note-button" onClick={addCheckNote}><span
                className="mdi mdi-playlist-check"/> &nbsp; Add Checklist</button>
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
                                            handleUpdateCheckList={updateCheckList}
                                            handleDeleteCheckList={deleteCheckList}/>
                    }
                })}
            </div>
        </div>
    );
};

export default ListBox;
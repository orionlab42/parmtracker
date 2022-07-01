import React, {useEffect, useState} from "react";
import NotesList from "./notesList";
import SearchBox from "./searchBox";
import {searchKeywordNotes} from "../utils/search";
import CheckListsList from "./checkListsList";

const ListBox = (props) => {
    const [notes, setNotes] = useState([]);
    const [checkLists, setCheckLists] = useState([]);
    const [searchQuery, setSearchQuery] = useState("");
    const [idCheckLists, setIdCheckLists] = useState(1);
    const [idNotes, setIdNotes] = useState(1);

    const giveIdChecklist = () => {
        setIdCheckLists(idCheckLists + 1);
        return idCheckLists;
    };

    const giveIdNotes = () => {
        setIdNotes(idNotes + 1);
        return idNotes;
    };

    useEffect(() => {
        const savedNotes = JSON.parse(localStorage.getItem('react-notes-app-data'));
        if (savedNotes) {
            setNotes(savedNotes);
        }
    }, []);

    useEffect(() => {
        const savedCheckLists = JSON.parse(localStorage.getItem('react-lists-app-data'));
        if (savedCheckLists) {
            setCheckLists(savedCheckLists);
        }
    }, []);

    useEffect(() => {
        localStorage.setItem('react-notes-app-data', JSON.stringify(notes));
    }, [notes]);

    useEffect(() => {
        localStorage.setItem('react-lists-app-data', JSON.stringify(checkLists));
    }, [checkLists]);

    const addNote = (text) => {
        console.log("Add new empty note in home");
        const date = new Date();
        const newNote = {
            id: giveIdNotes(),
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
            id: giveIdChecklist(),
            type: "checklist",
            title: "",
            list: [],
            date: date.toLocaleDateString()
        }
        setCheckLists([...checkLists, newList]);
    };

    const updateCheckList = (itemList) => {
        const newChecklists = checkLists.map(checkList => {
            if (checkList.id === itemList.id) {
                checkList.list = itemList.list
            }
            return checkList
        });
        setCheckLists(newChecklists);
    };

    const deleteCheckList = (id) => {
        const newChecklists = checkLists.filter(checkList => checkList.id !== id);
        setCheckLists(newChecklists);
    };

    const searchNote = (text) => {
        setSearchQuery(text);
    };

    let notesToDisplay = notes;
    if (searchQuery) {
        notesToDisplay = searchKeywordNotes(notes, searchQuery);
    }

    return (
        <div className="notes-list-container">
            <SearchBox value={searchQuery} onChange={searchNote}/>
            <button className="button is-link is-light add-note-button" onClick={addNote}><span
                className="mdi mdi-note-outline"/> &nbsp; Add Simple Note</button>
            <button className="button is-link is-light add-note-button" onClick={addCheckNote}><span
                className="mdi mdi-playlist-check"/> &nbsp; Add Checklist</button>
            <NotesList
                notes={notesToDisplay}
                handleAddNote={addNote}
                handleUpdateNote={updateNote}
                handleDeleteNote={deleteNote}
            />
            <CheckListsList
                checkLists={checkLists}
                handleUpdateCheckList={updateCheckList}
                handleDeleteCheckList={deleteCheckList}
            />
        </div>
    );
};

export default ListBox;
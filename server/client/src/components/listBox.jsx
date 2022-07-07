import React, {useEffect, useState} from "react";
import Note from "./common/note";
import CheckList from "./common/checkList";
import AgendaNote from "./common/agendaNote";
import {deleteNote, getNotes, saveNote} from "../services/noteService";
import {toast} from "react-toastify";
// import {searchKeywordNotes} from "../utils/search";
// import SearchBox from "./searchBox";
import { v4 as uuidv4 } from 'uuid';


const typeSimpleNote = 1;
const typeChecklist = 2;
const typeAgenda = 3;

const ListBox = ({user}) => {
    // const [notes, setNotes] = useState([]);
    const [notesFromServer, setNotesFromServer] = useState([]);
    // const [searchQuery, setSearchQuery] = useState("");
    const [updateNotes, setUpdateNotes] = useState(true);

    // Sorts the notelist by date even while a date is edited which is annoying
    // const sortedNotes = notelist.sort((a,b) => b.date - a.date);


    useEffect(() => {
        async function getAllNotes() {
            const {data: notes} = await getNotes();
            // console.log("Notes from server", notes)
            if (notes != null) {
                setNotesFromServer(notes);
                // setNotes(notes);
            }
        }
        getAllNotes();
    }, [updateNotes]);

    // useEffect(() => {
    //     let savedNotes = JSON.parse(localStorage.getItem('react-notelist-app-data'));
    //     if (savedNotes) {
    //         setNotes(savedNotes);
    //         if (savedNotes[0]) {
    //             setId(savedNotes[0].id + 1);
    //         }
    //     }
    // }, []);
    //
    // useEffect(() => {
    //     localStorage.setItem('react-notelist-app-data', JSON.stringify(notes));
    // }, [notes]);

    const addNote = async () => {
        let newNote = {
            // note_id: uuidv4(),
            user_id: user.user_id,
            note_type: typeSimpleNote,
            note_empty: true,
            note_title: "",
            note_text: ""
        };
        console.log("Type of", typeof newNote.note_id !== "number")
        // setNotes([newNote, ...notes]);
        // setNotesFromServer([newNote, ...notesFromServer]);
        setUpdateNotes(!updateNotes);
        await saveNote(newNote);
    };

    const addCheckNote = async () => {
        let newList = {
            // note_id: uuidv4(),
            user_id: user.user_id,
            note_type: typeChecklist,
            note_empty: true,
            note_title: "",
            note_text: "",
            note_list: []
        }
        // setNotes([newList, ...notes]);
        // setNotesFromServer([newList, ...notesFromServer]);
        setUpdateNotes(!updateNotes);
        await saveNote(newList);
    };

    const addAgendaNote = async () => {
        let newAgenda = {
            // note_id: uuidv4(),
            user_id: user.user_id,
            note_type: typeAgenda,
            note_empty: true,
            note_title: "",
            note_text: ""
        }
        // setNotes([newAgenda, ...notes]);
        // setNotesFromServer([newAgenda, ...notesFromServer]);
        setUpdateNotes(!updateNotes);
        await saveNote(newAgenda);
    };

    // const updateNote = async (newNote) => {
    //     // setUpdateNotes(!updateNotes);
    //     await saveNote(newNote);
    // };

    const handleDeleteNote = async (id) => {
        const originalNotes = notesFromServer;
        const notes = notesFromServer.filter(note => note.note_id !== id);
        // setNotesFromServer(notes);
        // setNotes(notes);
        try {
            await deleteNote(id)
            setUpdateNotes(!updateNotes);
        } catch (ex) {
            if (ex.response && ex.response.status === 404)
                toast('This note has already been deleted.');
            // setNotesFromServer(originalNotes);
            // setNotes(originalNotes);
        }
    };

    // const searchNote = (text) => {
    //     setSearchQuery(text);
    // };

    // let notesToDisplay = notelist;
    // if (searchQuery) {
    //     notesToDisplay = searchKeywordNotes(notelist, searchQuery);
    // }


    // console.log("All notes", notes);
    console.log("All notes from server", notesFromServer);


    return (
        <div className="notes-list-container">
            {/*<SearchBox value={searchQuery} onChange={searchNote}/>*/}
            <button className="button is-link is-light add-note-button" onClick={addNote}><span
                className="mdi mdi-note-outline"/> &nbsp; Add Simple Note
            </button>
            <button className="button is-link is-light add-note-button" onClick={addCheckNote}><span
                className="mdi mdi-playlist-check"/> &nbsp; Add Checklist
            </button>
            <button className="button is-link is-light add-note-button" onClick={addAgendaNote}><span
                className="mdi mdi-calendar-text"/> &nbsp; Add Planner
            </button>
            <div className="notes-list">
                {notesFromServer.map(note => {
                    if (note.note_type === typeSimpleNote) {
                        return <Note key={note.note_id}
                                     note={note}
                                     // handleUpdateNote={updateNote}
                                     onDeleteNote={handleDeleteNote}/>
                    }
                    if (note.note_type === typeChecklist) {
                        return <CheckList key={note.note_id}
                                          items={note}
                                          // handleUpdateCheckList={updateNote}
                                          onDeleteCheckList={handleDeleteNote}/>
                    }
                    if (note.note_type === typeAgenda) {
                        return <AgendaNote key={note.note_id}
                                           items={note}
                                           // handleUpdateAgendaNote={updateNote}
                                           onDeleteAgendaNote={handleDeleteNote}
                        />
                    }
                    return null
                })}
            </div>

        </div>
    );
};

export default ListBox;
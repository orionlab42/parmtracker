import React, {useEffect, useState} from "react";
import Note from "./common/note";
import CheckList from "./common/checkList";
import AgendaNote from "./common/agendaNote";
import {deleteItems, deleteNote, getNotes, saveNote, saveNoteUser} from "../services/noteService";
import {toast} from "react-toastify";
// import {searchKeywordNotes} from "../utils/search";
// import SearchBox from "./searchBox";
import { v4 as uuidv4 } from 'uuid';
import {getUsers} from "../services/userService";

const typeSimpleNote = 1;
const typeChecklist = 2;
const typeAgenda = 3;

const ListBox = ({user}) => {
    const [notes, setNotes] = useState([]);
    const [users, setUsers] = useState([]);
    // const [searchQuery, setSearchQuery] = useState("");

    useEffect(() => {
        async function getALlUsers() {
            const { data: users } = await getUsers();
            if (users != null) {
                setUsers(users);
            }
        }
        getALlUsers();
    }, [user]);

    useEffect(() => {
        async function getInitialNotes() {
            if (user.user_id) {
                const {data: notes} = await getNotes(user.user_id);
                if (notes != null) {
                    setNotes(notes);
                }
            }
        }
        getInitialNotes();
    }, [user]);


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

    const getAllNotes = async () => {
        const {data: newNotes} = await getNotes(user.user_id);
        if (notes != null) {
            setNotes(newNotes);
        }
    };

    const addNote = async (type) => {
        let newNote = {
            note_id: uuidv4(),
            note_type: type,
            note_empty: true,
            note_title: "",
            note_text: ""
        };
        await saveNote(newNote, user.user_id);
        getAllNotes().then();
    };

    const handleDeleteNote = async (noteId) => {
        let originalNotes = notes;
        let updatedNotes = notes.filter(n => n.note_id !== noteId);
        setNotes(updatedNotes);
        try {
            await deleteNote(noteId, user.user_id)
            await deleteItems(noteId)
        } catch (ex) {
            if (ex.response && ex.response.status === 404)
                toast('This note has already been deleted.');
            setNotes(originalNotes);
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
    // console.log("All notes from server", notes);

    return (
        <div className="notes-list-container">
            {/*<SearchBox value={searchQuery} onChange={searchNote}/>*/}
            <button className="button is-link is-light add-note-button" onClick={() => addNote(typeSimpleNote)}><span
                className="mdi mdi-note-outline"/> &nbsp; Add Simple Note
            </button>
            <button className="button is-link is-light add-note-button" onClick={() => addNote(typeChecklist)}><span
                className="mdi mdi-playlist-check"/> &nbsp; Add Checklist
            </button>
            <button className="button is-link is-light add-note-button" onClick={() => addNote(typeAgenda)}><span
                className="mdi mdi-calendar-text"/> &nbsp; Add Planner
            </button>
            <div className="notes-list">
                {notes.map(note => {
                    if (note.note_type === typeSimpleNote) {
                        return <Note key={note.note_id}
                                     note={note}
                                     users={users}
                                     onDeleteNote={handleDeleteNote}/>
                    }
                    if (note.note_type === typeChecklist) {
                        return <CheckList key={note.note_id}
                                          note={note}
                                          users={users}
                                          onDeleteNote={handleDeleteNote}/>
                    }
                    if (note.note_type === typeAgenda) {
                        return <AgendaNote key={note.note_id}
                                           note={note}
                                           user={user}
                                           users={users}
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
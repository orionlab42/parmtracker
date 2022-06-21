import React, {useState} from "react";
import Link from "react-router-dom/Link";
import NotesList from "./notesList";
import SearchBox from "./searchBox";
import {searchKeywordNotes} from "../utils/search";

const Home = (props) => {
    const [notes, setNotes] = useState([
        {
        id: 0,
        text: "This is my first note!",
        date: "15/04/2021"
        },
        {
        id: 1,
        text: "This is my second note!",
        date: "16/04/2021"
        },
        {
        id: 2,
        text: "This is my third note!",
        date: "17/04/2021"
        },
    ]);
    const [searchQuery, setSearchQuery] = useState("");

    let menu;
    if (props.user === "") {
        menu =  ( <h3 className="title">{'You are not logged in' }</h3>)
    } else {
        menu = (    <div>
                        <div className="add-more">
                            <Link to="/expenses/new" className="button is-link is-medium add-more-button">+</Link>
                        </div>
                        {/*<h3 className="title is-3 center-text">{'Hello ' + props.user.user_name}</h3>*/}
                     </div>)
    }

    const addNote = (text) => {
        let id = 0;
        if (notes.length > 0)  {
            id = notes[notes.length - 1].id + 1;
        }
        const date = new Date();
        const newNote = {
            id: id,
            text: text,
            date: date.toLocaleDateString()
        }
        const newNotes = [...notes, newNote];
        setNotes(newNotes);
    };

    const deleteNote = (deleteNote) => {
        const newNotes = notes.filter(note => note.id !== deleteNote.id);
        setNotes(newNotes);
    };

    const searchNote = (text) => {
        setSearchQuery(text);
    };

    let notesToDisplay = notes;
    if (searchQuery) {
        notesToDisplay = searchKeywordNotes(notes, searchQuery);
    }

    return (
        <div>
            {menu}
            <div className="notes-list-container">
                <SearchBox value={searchQuery} onChange={searchNote}/>
                <NotesList
                    notes={ notesToDisplay }
                    handleAddNote={addNote}
                    handleDeleteNote={deleteNote}/>
            </div>
        </div>
    );
};

export default Home;
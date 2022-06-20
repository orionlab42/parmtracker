import React from "react";
import Note from "./common/note";
import AddNote from "./common/addNote";

const NotesList = ({ notes }) => {
    return (
        <div className="notes-list-container">
            <h1 className="title is-3 center-text">NotesList</h1>
            <div className="notes-list">
                {notes.map(note => <Note note={ note }/>)}
                <AddNote/>
            </div>
        </div>
    );
};

export default NotesList;
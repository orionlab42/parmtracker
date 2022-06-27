import React from "react";
import Note from "./common/note";
import AddNote from "./common/addNote";
import AddCheckList from "./common/addCheckList";

const NotesList = ({ notes, handleAddNote, handleDeleteNote }) => {
    return (
        <div className="notes-list-container">
            <h1 className="title is-3 center-text">NotesList</h1>
            <div className="notes-list">
                {notes.map(note => <Note
                                        key={note.id}
                                        note={ note }
                                        handleDeleteNote={ handleDeleteNote }/>)}
                <AddNote handleAddNote={ handleAddNote }/>
                <AddCheckList handleAddNote={ handleAddNote }/>
            </div>
        </div>
    );
};

export default NotesList;
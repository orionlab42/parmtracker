import React from "react";
import Note from "./common/note";

const NotesList = ({ notes, handleUpdateNote, handleDeleteNote }) => {
    return (
        <div className="notes-list">
            {notes.map(note => <Note
                                    key={note.id}
                                    note={ note }
                                    handleUpdateNote={ handleUpdateNote }
                                    handleDeleteNote={ handleDeleteNote }/>)}
        </div>
    );
};

export default NotesList;
import React from "react";

const Note = ({ note, handleDeleteNote }) => {
    return (
        <div className="note">
            <span>{ note.text }</span>
            <div className="note-footer">
                <small>{ note.date }</small>
                <button onClick={() => handleDeleteNote(note)}><span className="mdi mdi-trash-can-outline"/></button>
            </div>
        </div>
    );
};

export default Note;
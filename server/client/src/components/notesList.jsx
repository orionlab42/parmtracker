import React from "react";
import Note from "./common/note";
import AddNote from "./common/addNote";
import AddCheckList from "./common/addCheckList";
import CheckList from "./common/checkList";

const NotesList = ({ notes, checkLists, handleAddNote, handleDeleteNote, handleAddChecklist }) => {
    return (
        <div className="notes-list-container">
            <h1 className="title is-3 center-text">NotesList</h1>
            <div className="notes-list">
                {notes.map(note => <Note
                                        key={note.id}
                                        note={ note }
                                        handleDeleteNote={ handleDeleteNote }/>)}

                {checkLists.map(list => <CheckList
                                            key={list.id}
                                            items={list}
                                            handleDeleteNote={handleDeleteNote}
                                                />)}
                <AddNote handleAddNote={ handleAddNote }/>
                <AddCheckList handleAddChecklist={ handleAddChecklist }/>
            </div>
        </div>
    );
};

export default NotesList;
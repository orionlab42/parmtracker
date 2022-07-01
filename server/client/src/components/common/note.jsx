import React, {useEffect, useState} from "react";

const Note = ({ note, handleUpdateNote, handleDeleteNote }) => {
    const [updateNote, setUpdateNote] = useState(note);
    const [editOn, setEditOn] = useState(false);
    const [titleOn, setTitleOn] = useState(false);

    const renderEdit = () => {
        setEditOn(!editOn);
    };

    useEffect(() => {
        setUpdateNote(note);
    }, [note]);

    const editChange = (e) => {
        let newNote = updateNote;
        newNote.text = e.target.value;
        setUpdateNote(newNote);
        handleUpdateNote(newNote);
    }

    const edit = (
        <textarea rows="10" cols="32" placeholder="Type to add a note..." value={updateNote.text} onChange={ editChange }/>
    );

    const renderTitleInput = () => {
        setTitleOn(!titleOn);
    };

    const titleChange = (e) => {
        let newNote = updateNote;
        newNote.title = e.target.value
        setUpdateNote(newNote);
        handleUpdateNote(newNote);
    }

    const title = (
        <input
            className="title-input edit"
            type="text"
            placeholder="Title here..."
            value={note.title}
            name="text"
            onChange={titleChange}
        />
    );

    return (
            <div className="note">
                {!titleOn && <h4 className="note-title">{note.title}</h4>}
                {titleOn && title}
                <div className="note-body">
                    <div className="note-content">
                        { editOn && edit}
                        { !editOn && <span >{ note.text }</span>}
                        <div className="note-footer">
                            <small>{ note.date }</small>
                        </div>
                    </div>
                    <div className="simple-note-buttons">
                        <div className="edit-note-buttons">
                            <button className="button is-link is-light  mdi mdi-plus"
                                    onClick={renderEdit}/>
                            <button className="button is-link is-light  mdi mdi-format-title"
                                    onClick={renderTitleInput}/>
                        </div>
                        <button className="button is-link is-light  mdi mdi-trash-can-outline"
                                onClick={() => handleDeleteNote(note)}/>
                    </div>
                </div>
            </div>
    );
};

export default Note;
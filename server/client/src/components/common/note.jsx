import React, {useEffect, useState} from "react";
import ReactMarkdown from "react-markdown";
import {saveNote} from "../../services/noteService";

const Note = ({ note, onDeleteNote }) => {
    const [textOn, setTextOn] = useState(false);
    const [titleOn, setTitleOn] = useState(false);
    const [editText, setEditText] = useState({});

    useEffect(() => {
        setEditText({
            note_empty: note.note_empty,
            note_title: note.note_title,
            note_text: note.note_text,
            updated_at: note.updated_at});
    }, []);

    const editChange = (e, type) => {
        if (type === "text") {
            setEditText({
                note_title: editText.note_title,
                note_text: e.target.value,
                updated_at: new Date()
            })
        }
        if (type === "title") {
            setEditText({
                note_title: e.target.value,
                note_text: editText.note_text,
                updated_at: new Date()
            })
        }
        console.log("edit", editText);
    }

    const renderTitleInput = async () => {
        setTitleOn(!titleOn);
        let noteUpdate = {note_id: note.note_id, note_empty: false, note_title: editText.note_title, note_text: editText.note_text};
        await saveNote(noteUpdate);
    };

    const title = (
        <input
            className="title-input edit"
            type="text"
            placeholder="Title here..."
            value={ editText.note_title }
            name="text"
            onChange={(e) => editChange(e, "title")}
            autoFocus
        />
    );

    const renderTextArea = async () => {
        setTextOn(!textOn);
        let noteUpdate = {note_id: note.note_id, note_empty: false, note_title: editText.note_title, note_text: editText.note_text};
        await saveNote(noteUpdate);
    };

    const edit = (
        <textarea rows="10" cols="32" placeholder="Type to add a note..." value={editText.note_text} onChange={(e) => editChange(e,"text") }/>
    );

    return (
            <div className="note">
                {!titleOn && <h4 className="note-title">{ editText.note_title }</h4>}
                {titleOn && title}
                <div className="note-body">
                    <div className="note-content">
                        { !textOn && <ReactMarkdown >{editText.note_text}</ReactMarkdown>}
                        { textOn && edit}
                        <div className="note-footer">
                            <small>{!note.note_empty ? "Last modified: " +  new Date(editText.updated_at).toLocaleDateString("en-GB", {
                                    hour: "2-digit",
                                    minute:  "2-digit",
                                }) : ""}</small>
                        </div>
                    </div>
                    <div className="simple-note-buttons">
                        <div className="edit-note-buttons">
                            <button className="button is-link is-light  mdi mdi-format-title"
                                    onClick={renderTitleInput}/>
                            <button className="button is-link is-light  mdi mdi-plus"
                                    onClick={renderTextArea}/>
                        </div>
                        <button className="button is-link is-light  mdi mdi-trash-can-outline"
                                onClick={() => onDeleteNote(note.note_id)}/>
                    </div>
                </div>
            </div>
    );
};

export default Note;
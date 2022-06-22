import React, {useState} from "react";
import CheckListForm from "./checkListForm";
import CheckList from "./checkList";

const AddNote = ({ handleAddNote }) => {
    const [noteText, setNoteText] = useState("");
    const characterLimit = 200;

    const handleChange = (event) => {
        if (characterLimit - event.target.value.length >= 0) {
            setNoteText(event.target.value);
        }
    };

    const handleSaveClick = () => {
        if (noteText.trim().length > 0) {
            handleAddNote(noteText);
            setNoteText("");
        }
    };

    return (
        <div className="note add-new-note">
            <CheckList/>
            <textarea rows="8" cols="10" placeholder="Type to add a note..." value={noteText} onChange={ handleChange }/>
            <div className="note-footer">
                <div>
                    <small>{ characterLimit - noteText.length }/200</small>
                    <div className="note-options">
                        <button className="add-title-button button is-small is-link is-light" data-title="Add title"><span className="mdi mdi-format-title"/></button>
                        {/*<button className="simple-note-button button is-small is-link is-light" data-title="Change to simple note"><span className="mdi mdi-note-outline"/></button>*/}
                        <button className="list-note-button button is-small is-link is-light" data-title="Change to list"><span className="mdi mdi-playlist-check"/></button>
                        <button className="planner-note-button button is-small is-link is-light" data-title="Change to planner"><span className="mdi mdi-calendar-month-outline"/></button>
                    </div>
                </div>

                <button className="save-button button is-link is-light" onClick={handleSaveClick}><span
                    className="mdi mdi-content-save"/> &nbsp; Save
                </button>
            </div>
        </div>
    );
};

export default AddNote;
import React from "react";

const AddNote = (props) => {
    return (
        <div className="note add-new-note">
            <textarea rows="8" cols="10" placeholder="Type to add a note..."></textarea>
            <div className="note-footer">
                <small>200 Remaining</small>
                <button className="save-button button is-light is-small"><span className="mdi mdi-content-save"/> &nbsp; Save</button>
            </div>
        </div>
    );
};

export default AddNote;
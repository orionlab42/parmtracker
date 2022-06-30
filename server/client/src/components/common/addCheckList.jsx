import React, {useState} from "react";
import CheckList from "./checkList";

const AddCheckList = ({ handleAddChecklist }) => {
    const [noteText, setNoteText] = useState("");
    const [items, setItems] = useState([]);

    const addItem = item => {
        if (item.text.trim().length > 0) {
            const newItems = [...items, item];
            setItems(newItems);
        }
        // console.log("AddItem", item);
    };

    const completeItem = (id) => {
        const updatedItems = items.map(item => {
            if (item.id === id) {
                item.isComplete = !item.isComplete;
            }
            return item
        });
        setItems([...updatedItems]);
    };

    const deleteItem = (id) => {
        const updatedItems = items.filter(item => item.id !== id);
        setItems([...updatedItems]);
    };

    const updateItem = (id, newValue, isComplete) => {
        console.log("Update", id);
        const updatedItems = items.map(item => {
            if (item.id === id) {
                item.text = newValue.text;
                item.isComplete = isComplete;
            }
            return item
        });
        setItems([...updatedItems]);
    };

    // const handleChange = (event) => {
    //     // setNoteText(event.target.value);
    //     console.log("Event", event);
    // };

    const handleSaveClick = () => {
        // if (noteText.trim().length > 0) {
        //     handleAddChecklist(noteText);
        //     setNoteText("");
        // }
        handleAddChecklist(items);
        setItems([]);
        console.log("To save", items);
    };

    return (
        <div className="note add-new-note">
            <CheckList
                items={items}
                handleAddItem={addItem}
                handleCompleteItem={completeItem}
                handleDeleteItem={deleteItem}
                handleUpdateItem={updateItem}
            />
            {/*<textarea rows="8" cols="10" placeholder="Type to add a note..." value={noteText} onChange={ handleChange }/>*/}
            <div className="note-footer">
                <div>
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

export default AddCheckList;
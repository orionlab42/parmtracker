import React, {useState, useEffect} from "react";
import CheckListForm from "./checkListForm";
import CheckListItems from "./checkListItems";
import {saveNote, saveItem, deleteItem, getNotes, getItems} from "../../services/noteService";

const CheckList = ({ note, onDeleteNote }) => {
    const [items, setItems] = useState([]);
    const [titleOn, setTitleOn] = useState(false);
    const [editText, setEditText] = useState({});
    const [timeToGetItems, setTimeToGetItems] = useState(true);

    useEffect(() => {
        async function getAllItems() {
            const {data: newItems} = await getItems(note.note_id);
            console.log("items from server", newItems)
            if (newItems != null) {
                setItems(newItems);
            }
        }
        getAllItems();
    }, [timeToGetItems]);

    const addItem = async item => {
        if (item.item_text.trim().length > 0) {
            let newCheckList = [...items, item];
            setItems(newCheckList);
            item.note_id = note.note_id;
            await saveItem(item).then();
            setTimeToGetItems(!timeToGetItems);
        }
    };

    const completeItem = async (id) => {
        let updatedItem;
        let newCheckList = items.map(item => {
            if (item.item_id === id) {
                item.item_is_complete = !item.item_is_complete;
                updatedItem = item;
            }
            return item
        });
        setItems(newCheckList);
        await saveItem(updatedItem);
    };

    const removeItem = async (id) => {
        let newCheckList = items.filter(item => item.item_id !== id);
        setItems(newCheckList);
        await deleteItem(id);
        setTimeToGetItems(!timeToGetItems);
    };

    const updateItem = async (id, newValue, isComplete) => {
        let updatedItem;
        let newCheckList = items.map(item => {
            if (item.item_id === id) {
                item.item_text = newValue.item_text;
                item.item_is_complete = isComplete;
                updatedItem = item;
            }
            return item
        });
        setItems(newCheckList);
        await saveItem(updatedItem);
    };

    const editTitle = (e) => {
        setEditText({
            note_title: e.target.value,
            updated_at: new Date()
        })
    }

    const renderTitleInput = async () => {
        setTitleOn(!titleOn);
        let noteUpdate = {note_id: note.note_id, note_empty: false, note_title: editText.note_title };
        await saveNote(noteUpdate);
    };

    const title = (
        <input
            className="title-input edit"
            type="text"
            placeholder="Title here..."
            value={ editText.note_title }
            name="text"
            onChange={editTitle}
            autoFocus
        />
    );

    return (
        <div className="note checklist">
            <div className="checklist-container">
                {!titleOn && <h4 className="note-title">{editText.note_title}</h4>}
                {titleOn && title}
                <div className="checklist-top">
                    <CheckListForm onSubmit={addItem}/>
                    <button className="button is-link is-light  mdi mdi-format-title"
                            onClick={renderTitleInput}/>
                </div>
                <div className="checklist-main">
                    <div className="checklist-body">
                        {!note.note_empty && <CheckListItems
                            items={items}
                            handleCompleteItem={completeItem}
                            handleDeleteItem={removeItem}
                            handleUpdateItem={updateItem}
                        />}
                        <div className="note-footer">
                            <small>{!note.note_empty ? "Last modified:" +  new Date(note.updated_at).toLocaleDateString("en-GB", {
                                hour: "2-digit",
                                minute:  "2-digit",
                            }) : ""}</small>
                        </div>
                    </div>
                    <button className="button is-link is-light  mdi mdi-trash-can-outline"
                            onClick={() => onDeleteNote(note.note_id)}/>
                </div>
            </div>
        </div>
    );
}

export default CheckList;
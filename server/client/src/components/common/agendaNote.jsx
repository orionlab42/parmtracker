import React, {useEffect, useState} from "react";
import DatePicker from "react-datepicker";
import {getItems, saveItem, saveItems, saveNote} from "../../services/noteService";
import "react-datepicker/dist/react-datepicker.css";


const AgendaNote = ({note, user, onDeleteAgendaNote}) => {
    const [items, setItems] = useState([]);
    const [dateRange, setDateRange] = useState([null, null]);
    const [startDate, endDate] = dateRange;
    const [titleOn, setTitleOn] = useState(false);
    const [editText, setEditText] = useState({});

    useEffect(() => {
        setEditText({note_empty: note.note_empty, note_title: note.note_title, updated_at: note.updated_at,});
    }, []);

    useEffect(() => {
        async function getAllItems() {
            const {data: newItems} = await getItems(note.note_id);
            if (newItems != null) {
                setItems(newItems);
            }
        }
        getAllItems();
        console.log("Went through1", note)
    }, []);

    useEffect(() => {
        async function saveNewItems() {
            const {data: newItems} = await saveItems(note.note_id, startDate, endDate).then();
            if (newItems != null) {
                console.log("Went through3", newItems)
                note.note_empty = false;
                await saveNote(note, user.user_id);
                setItems(newItems);
            }
        }
        saveNewItems();
        console.log("Went through2", note)
    }, [dateRange]);

    // useEffect(() => {
    //     async function sendNote() {
    //         note.note_empty = false;
    //         await saveNote(note);
    //     }
    //
    //     sendNote().then();
    // }, [items]);

    const editTitle = (e) => {
        setEditText({note_title: e.target.value, updated_at: new Date()});
    }

    const renderTitleInput = async () => {
        setTitleOn(!titleOn);
        let noteUpdate = {note_id: note.note_id, note_empty: false, note_title: editText.note_title};
        await saveNote(noteUpdate);
    };

    const title = (
        <input
            className="title-input edit"
            type="text"
            placeholder="Title here..."
            value={editText.note_title}
            name="text"
            onChange={editTitle}
            autoFocus
        />
    );

    const sendItemToServer = async (item) => {
        await saveItem(item);
    };

    const itemChange = (e, date) => {
        let newItems = items.map(item => {
            if (item.item_date === date) {
                item.item_text = e.target.value;
                sendItemToServer(item).then();
            }
            return item;
        });
        setItems(newItems);
    };

    const isCompleteItem = (id) => {
        let newAgenda = items.map(item => {
            if (item.item_id === id) {
                item.item_is_complete = !item.item_is_complete;
                sendItemToServer(item).then();
            }
            return item;
        });
        setItems(newAgenda);
    };

    const handleEnter = (event) => {
        if (event.key.toLowerCase() === "enter") {
            let form = event.target.form;
            let index = [...form].indexOf(event.target);
            let formLength = form.elements.length;
            if (index < formLength - 1) {
                form.elements[index + 1].focus();
                event.preventDefault();
            }
            if (index === formLength - 1) {
                form.elements[0].focus();
                event.preventDefault();
            }
        }
    };

    const itemList = (
        <form className="agenda-item-list">
            {items && items.map((item) => <div className={item.item_is_complete ? 'checked agenda-item' : 'agenda-item'}
                                      key={item.item_id}>
                <div className="agenda-item-date" onClick={(e) => isCompleteItem(item.item_id)}>
                    <span>{new Date(item.item_date).toLocaleDateString("en-US", {
                        month: "short",
                        day: "numeric"
                    })}</span> /&nbsp;
                    <span>{new Date(item.item_date).toLocaleDateString("en-US", {
                        weekday: "short"
                    })}</span>
                </div>
                <input
                    value={item.item_text}
                    onChange={(e) => itemChange(e, item.item_date)}
                    onKeyDown={handleEnter}
                />
            </div>)}
        </form>
    );

    return (
        <div className="note">
            {!titleOn && <h4 className="note-title">{editText.note_title}</h4>}
            {titleOn && title}
            <div className="note-body">
                <div className="note-content">
                    <div className="date-range-picker">
                        <DatePicker
                            placeholderText="Select date range..."
                            selectsRange={true}
                            startDate={startDate}
                            endDate={endDate}
                            onChange={setDateRange}
                            isClearable={true}
                        />
                    </div>
                    {!note.note_empty && itemList}
                    <div className="note-footer">
                        <small>{!note.note_empty ? "Last modified:" + new Date(note.updated_at).toLocaleDateString("en-GB", {
                            hour: "2-digit",
                            minute: "2-digit",
                        }) : ""}</small>
                    </div>
                </div>
                <div className="simple-note-buttons">
                    <div className="edit-note-buttons">
                        <button className="button is-link is-light  mdi mdi-format-title"
                                onClick={renderTitleInput}/>

                    </div>
                    <button className="button is-link is-light  mdi mdi-trash-can-outline"
                            onClick={() => onDeleteAgendaNote(note.note_id)}/>
                </div>
            </div>
        </div>
    );
};

export default AgendaNote;
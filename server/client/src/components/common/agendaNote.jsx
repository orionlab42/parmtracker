import React, {useEffect, useState} from "react";
import DatePicker from "react-datepicker";
import { v4 as uuidv4 } from 'uuid';
import {getItems, saveNote, saveItem, deleteItem} from "../../services/noteService";
import "react-datepicker/dist/react-datepicker.css";


const AgendaNote = ({ note, onDeleteAgendaNote }) => {
    const [items, setItems] = useState([]);
    const [dateRange, setDateRange] = useState([null, null]);
    const [startDate, endDate] = dateRange;
    const [titleOn, setTitleOn] = useState(false);
    const [editText, setEditText] = useState({});
    const [timeToGetItems, setTimeToGetItems] = useState(true);

    useEffect(() => {
        setEditText({
            note_empty: note.note_empty,
            note_title: note.note_title,
            updated_at: note.updated_at});
    }, []);

    useEffect(() => {
        async function getAllItems() {
            const {data: newItems} = await getItems(note.note_id);
            if (newItems != null) {
                setItems(newItems);
            }
        }
        getAllItems();
    }, [timeToGetItems]);

    useEffect(() => {
        createItems();
    }, [dateRange]);

    useEffect(() => {
        async function sendNote() {
            note.note_empty = false;
            await saveNote(note);
        }
        sendNote().then();
    }, [items]);

    const sendItemToServer = async (item) => {
        await saveItem(item);
    };

    const removeItemFromServer = async (id) => {
        await deleteItem(id);
    };

    const createItems = () => {
        let newItems = [];
        let lengthItemList = 0;
        if (items !== null) {
            lengthItemList = items.length;
        }
        // this is the whenever it loads
        if (startDate === null || endDate === null) {
            if (lengthItemList <= 1) {
                return [];
            }
           return  note.note_items;
        }

        // this is the first time we enter a date range
        if (lengthItemList <= 1) {
            for (let d = new Date(startDate); d <= new Date(endDate); d.setDate(d.getDate() + 1)) {
                newItems.push({
                    note_id: note.note_id,
                    item_id: uuidv4(),
                    item_date: new Date(d),
                    item_text: "",
                    item_is_complete: false});
            }
            newItems.map((item) => sendItemToServer(item).then());
            setTimeToGetItems(!timeToGetItems);
            // console.log("Triggered from here the update 1");
        }

        if (lengthItemList > 1) {
            let prevStartDate = new Date(items[0].item_date);
            let prevEndDate = new Date(items[lengthItemList-1].item_date);
            if (prevStartDate <= startDate && endDate <= prevEndDate) {
                let itemsToDelete = items.filter(item => (new Date(item.item_date) < startDate || endDate < new Date(item.item_date)));
                itemsToDelete.map(item => removeItemFromServer(item.item_id))
                setTimeToGetItems(!timeToGetItems);
                // console.log("Triggered from here the update 2");
            }
            // in case on one end of the interval we need to add new items
            if (startDate < prevStartDate || prevEndDate < endDate) {
                for (let d = new Date(startDate); d <= new Date(endDate); d.setDate(d.getDate() + 1)) {
                    let itemsToDelete = items.filter(item => (new Date(item.item_date) < startDate || endDate < new Date(item.item_date)));
                    itemsToDelete.map(item => removeItemFromServer(item.item_id));
                    setTimeToGetItems(!timeToGetItems);
                    // console.log("Triggered from here the update 3");

                    let existingItem = items.filter(item => (new Date(item.item_date).getTime() === d.getTime()));
                    if (existingItem.length > 0) {
                        let itemUpdate = {
                            note_id: note.note_id,
                            item_id: existingItem[0].item_id,
                            item_date: new Date(d),
                            item_text: existingItem[0].item_text,
                            item_is_complete: existingItem[0].item_is_complete};
                        newItems.push(itemUpdate);
                        sendItemToServer(itemUpdate).then();
                        setTimeToGetItems(!timeToGetItems);
                    } else {
                        let itemNew = {
                            note_id: note.note_id,
                            item_id: uuidv4(),
                            item_date: new Date(d),
                            item_text: "",
                            item_is_complete: false};
                        newItems.push(itemNew);
                        sendItemToServer(itemNew).then();
                        setTimeToGetItems(!timeToGetItems);
                        // console.log("Triggered from here the update 4",itemNew.item_id);
                    }
                }
                setTimeToGetItems(!timeToGetItems);
                // console.log("Triggered from here the update 5");
            }
        }
        return newItems
    };

    const editTitle = (e) => {
        setEditText({note_title: e.target.value, updated_at: new Date()});
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
            {items.map((item) =>  <div  className={item.item_is_complete ? 'checked agenda-item' : 'agenda-item'} key={item.item_id}>
                <div className="agenda-item-date" onClick={(e) => isCompleteItem(item.item_id)}>
                    <span>{new Date(item.item_date).toLocaleDateString("en-US", {
                        month:  "short",
                        day:"numeric"
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
                    { !note.note_empty && itemList}
                    <div className="note-footer">
                        <small>{!note.note_empty ? "Last modified:" +  new Date(note.updated_at).toLocaleDateString("en-GB", {
                            hour: "2-digit",
                            minute:  "2-digit",
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
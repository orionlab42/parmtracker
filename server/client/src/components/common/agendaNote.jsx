import React, {useEffect, useState} from "react";
import DatePicker from "react-datepicker";
import { v4 as uuidv4 } from 'uuid';

import "react-datepicker/dist/react-datepicker.css";
import {getItems, saveNote} from "../../services/noteService";

const AgendaNote = ({ note, onDeleteAgendaNote }) => {
    const [items, setItems] = useState({});

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
        let newItems = createItems();
        setItems(newItems);
    }, [dateRange]);

    const createItems = () => {
        let newItems = [];
        let lengthItemList = 0;
        if (note.note_items !== null) {
            lengthItemList = note.note_items.length;
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
                newItems.push({item_id: uuidv4(), item_date: new Date(d), item_text: "", item_is_complete: false});
            }
        }

        if (lengthItemList > 1) {
            let prevStartDate = note.note_items[0].date;
            let prevEndDate = note.note_items[lengthItemList-1].date;

            // in case the new interval is smaller than the previous, we loose items
            if (prevStartDate <= startDate && endDate <= prevEndDate) {
                newItems = note.note_items.filter(item => (startDate <= item.date && item.date <= endDate))
            }
            // in case on one end of the interval we need to add new items
            if (startDate < prevStartDate || prevEndDate < endDate) {
                let idItem = 0;
                for (let d = new Date(startDate); d <= new Date(endDate); d.setDate(d.getDate() + 1)) {
                    let existingItem = note.note_items.filter(item => (item.date.getTime() === d.getTime()));
                    if (existingItem.length > 0) {
                        newItems.push({id: idItem, date: new Date(d), text: existingItem[0].text, isComplete: existingItem[0].isComplete});
                        idItem++;
                    } else {
                        newItems.push({id: idItem, date: new Date(d), text: ""});
                        idItem++;
                    }
                }
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
        let newAgenda = items;
        newAgenda.list.map(item => {
            if (item.date === date) {
                item.text = e.target.value;
            }
            return item;
        });
        setItems(newAgenda);
        // handleUpdateAgendaNote(updateAgendaNote);
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

    const isCompleteItem = (id) => {
        let newAgenda = items;
        newAgenda.list.map(item => {
            if (item.id === id) {
                item.isComplete = !item.isComplete;
            }
            return item;
        });
        setItems(newAgenda);
        // handleUpdateAgendaNote(updateAgendaNote);
    };

    const itemList = (
        <form className="agenda-item-list">
            {!note.note_empty && items.map((item)=>  <div  className={item.isComplete ? 'checked agenda-item' : 'agenda-item'} key={item.id}>
                <div className="agenda-item-date" onClick={(e) => isCompleteItem(item.id)}>
                    <span>{new Date(item.date).toLocaleDateString("en-US", {
                        month:  "short",
                        day:"numeric"
                    })}</span> /&nbsp;
                    <span>{new Date(item.date).toLocaleDateString("en-US", {
                        weekday: "short"
                    })}</span>
                </div>
                <input
                    value={item.text}
                    onChange={(e) => itemChange(e, item.date)}
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
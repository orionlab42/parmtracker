import React, {useEffect, useState} from "react";
import DatePicker from "react-datepicker";

import "react-datepicker/dist/react-datepicker.css";

const AgendaNote = ({ items, onDeleteAgendaNote }) => {
    const [updateAgendaNote, setUpdateAgendaNote] = useState(items);
    const [titleOn, setTitleOn] = useState(false);
    const [dateRange, setDateRange] = useState([null, null]);
    const [startDate, endDate] = dateRange;

    useEffect(() => {
        setUpdateAgendaNote(items);
    }, [items]);

    useEffect(() => {
        let newAgenda = updateAgendaNote;
        newAgenda.list = createItems();
        setUpdateAgendaNote(newAgenda);
        // handleUpdateAgendaNote(newAgenda);
    }, [dateRange]);

    const createItems = () => {
        let newItems = [];
        let lengthItemList = 0;
        if (items.list) {
            lengthItemList = items.list.length;
        }

        // this is the whenever it loads
        if (startDate === null || endDate === null) {
            if (lengthItemList <= 1) {
                return [];
            }
           return  items.list;
        }

        // this is the first time we enter a date range
        if (lengthItemList <= 1) {
            let idItem = 0;
            for (let d = new Date(startDate); d <= new Date(endDate); d.setDate(d.getDate() + 1)) {
                newItems.push({id: idItem, date: new Date(d), text: "", isComplete: false});
                idItem++;
            }
        }

        if (lengthItemList > 1) {
            let prevStartDate = items.list[0].date;
            let prevEndDate = items.list[lengthItemList-1].date;

            // in case the new interval is smaller than the previous, we loose items
            if (prevStartDate <= startDate && endDate <= prevEndDate) {
                newItems = items.list.filter(item => (startDate <= item.date && item.date <= endDate))
            }
            // in case on one end of the interval we need to add new items
            if (startDate < prevStartDate || prevEndDate < endDate) {
                let idItem = 0;
                for (let d = new Date(startDate); d <= new Date(endDate); d.setDate(d.getDate() + 1)) {
                    let existingItem = items.list.filter(item => (item.date.getTime() === d.getTime()));
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

    const renderTitleInput = () => {
        setTitleOn(!titleOn);
    };

    const titleChange = (e) => {
        let newNote = updateAgendaNote;
        newNote.title = e.target.value
        setUpdateAgendaNote(newNote);
        // handleUpdateAgendaNote(newNote);
    }

    const title = (
        <input
            className="title-input edit"
            type="text"
            placeholder="Title here..."
            value={items.title}
            name="text"
            onChange={titleChange}
            autoFocus
        />
    );

    const itemChange = (e, date) => {
        let newAgenda = updateAgendaNote;
        newAgenda.list.map(item => {
            if (item.date === date) {
                item.text = e.target.value;
            }
            return item;
        });
        setUpdateAgendaNote(newAgenda);
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
        let newAgenda = updateAgendaNote;
        newAgenda.list.map(item => {
            if (item.id === id) {
                item.isComplete = !item.isComplete;
            }
            return item;
        });
        setUpdateAgendaNote(newAgenda);
        // handleUpdateAgendaNote(updateAgendaNote);
    };

    const itemList = (
        <form className="agenda-item-list">
            {!items.note_empty && items.list.map((item)=>  <div  className={item.isComplete ? 'checked agenda-item' : 'agenda-item'} key={item.id}>
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
            {!titleOn && <h4 className="note-title">{items.note_title}</h4>}
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
                    { !items.note_empty && itemList}
                    <div className="note-footer">
                        <small>{!items.note_empty ? "Last modified:" +  new Date(items.updated_at).toLocaleDateString("en-GB", {
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
                            onClick={() => onDeleteAgendaNote(items.note_id)}/>
                </div>
            </div>
        </div>
    );
};

export default AgendaNote;
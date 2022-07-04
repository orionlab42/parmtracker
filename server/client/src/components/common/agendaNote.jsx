import React, {useEffect, useState} from "react";
import ReactMarkdown from "react-markdown";
import DatePicker from "react-datepicker";

import "react-datepicker/dist/react-datepicker.css";

const AgendaNote = ({ items, handleDeleteAgenda, handleUpdateAgendaNote }) => {
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
        handleUpdateAgendaNote(newAgenda);
    }, [dateRange]);

    const createItems = () => {
        let emptyItems = [];

        if (startDate === null || endDate === null) {
            let lengthItemList = 0;
            if (items.list) {
                lengthItemList = items.list.length;
            }
            if (lengthItemList <= 1) {
                return [];
            }
           return  items.list;
        }
        let idItem = 0;
        for (let d = new Date(startDate); d <= new Date(endDate); d.setDate(d.getDate() + 1)) {
            emptyItems.push({id: idItem, date: new Date(d), text: ""});
            idItem = idItem + 1;
        }

        return emptyItems
    };

    const renderTitleInput = () => {
        setTitleOn(!titleOn);
    };

    const titleChange = (e) => {
        let newNote = updateAgendaNote;
        newNote.title = e.target.value
        setUpdateAgendaNote(newNote);
        handleUpdateAgendaNote(newNote);
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
        console.log("text",e.target.value);
        let newAgenda = updateAgendaNote;
        newAgenda.list.map(item => {
            if (item.date === date) {
                item.text = e.target.value;
            }
            return item;
        });
        setUpdateAgendaNote(newAgenda);
        handleUpdateAgendaNote(updateAgendaNote);
    };

    const itemList = (
        <div>
            {!items.empty && items.list.map(item =>  <div key={item.id}>
                <span>{new Date(item.date).toLocaleDateString()}</span>
                <input
                    placeholder="Text here..."
                    value={item.text}
                    onChange={(e) => itemChange(e, item.date)}
                />
            </div>)}
        </div>
    );

    return (
        <div className="note">
            {!titleOn && <h4 className="note-title">{items.title}</h4>}
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
                    { !items.empty && itemList}
                    <div className="note-footer">
                        <small>{!items.empty ? "Last modified:" +  new Date(items.date).toLocaleDateString("en-GB", {
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
                            onClick={() => handleDeleteAgenda(items.id)}/>
                </div>
            </div>
        </div>
    );
};

export default AgendaNote;
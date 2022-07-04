import React, {useEffect, useState} from "react";
import ReactMarkdown from "react-markdown";
import DatePicker from "react-datepicker";

import "react-datepicker/dist/react-datepicker.css";

const AgendaNote = ({ items, handleDeleteAgenda, handleUpdateAgendaNote }) => {
    const [updateAgendaNote, setUpdateAgendaNote] = useState(items);
    const [titleOn, setTitleOn] = useState(false);
    const [dateRange, setDateRange] = useState([null, null]);
    const [startDate, endDate] = dateRange;
    const [start1Date, setStart1Date] = useState(new Date());
    const [end1Date, setEnd1Date] = useState(new Date());

    console.log("Date range", dateRange);
    console.log("StartDate", startDate);
    console.log("EndDate", endDate);

    useEffect(() => {
        setUpdateAgendaNote(items);
    }, [items]);


    const renderCalendar = () => {
       console.log("Pick a date");
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

    return (
        <div className="note">
            {!titleOn && <h4 className="note-title">{items.title}</h4>}
            {titleOn && title}
            <div className="note-body">
                <div className="note-content">
                    {/*<p>{note.list}</p>*/}
                    <div className="note-footer">
                        <small>{!items.empty ? "Last modified:" +  new Date(items.date).toLocaleDateString("en-GB", {
                            hour: "2-digit",
                            minute:  "2-digit",
                        }) : ""}</small>
                    </div>
                </div>
                <div className="simple-note-buttons">
                    <div className="edit-note-buttons">
                        {/*<button className="button is-link is-light  mdi mdi-plus"*/}
                        {/*        onClick={renderEdit}/>*/}
                        {/*<button className="button is-link is-light mdi mdi-calendar-plus"*/}
                        {/*        onClick={renderCalendar}/>*/}
                        {/*<div className="control picker" id="date-picker" >*/}
                        {/*    <DatePicker*/}
                        {/*        id="date"*/}
                        {/*        selected={selectedStartingDate}*/}
                        {/*        onChange={setSelectedStartingDate}*/}
                        {/*        name="Date"*/}
                        {/*        dateFormat={"dd/MM/yyyy"}*/}
                        {/*    />*/}
                        {/*</div>*/}
                        <DatePicker
                            // placeholderText="Select date range..."
                            selectsRange={true}
                            startDate={startDate}
                            endDate={endDate}
                            onChange={setDateRange}
                            isClearable={true}
                        />
                        <DatePicker
                            selected={start1Date}
                            onChange={(date) => setStart1Date(date)}
                            selectsStart
                            startDate={start1Date}
                            endDate={end1Date}
                        />
                        <DatePicker
                            selected={end1Date}
                            onChange={(date) => setEnd1Date(date)}
                            selectsEnd
                            startDate={start1Date}
                            endDate={end1Date}
                            minDate={startDate}
                        />
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
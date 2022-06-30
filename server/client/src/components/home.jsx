import React, {useEffect, useState} from "react";
import Link from "react-router-dom/Link";
import NotesList from "./notesList";
import SearchBox from "./searchBox";
import {searchKeywordNotes} from "../utils/search";
import CheckListsList from "./checkListsList";

const Home = (props) => {
    const [notes, setNotes] = useState([]);
    const [checkLists, setCheckLists] = useState([]);
    const [searchQuery, setSearchQuery] = useState("");
    const [idCheckLists, setIdCheckLists] = useState(1);

    const giveId = () => {
        setIdCheckLists(idCheckLists + 1);
        return idCheckLists;
    };

    useEffect(() => {
        const savedNotes = JSON.parse(localStorage.getItem('react-notes-app-data'));
        if (savedNotes) {
            setNotes(savedNotes);
        }
    }, []);

    useEffect(() => {
        localStorage.setItem('react-notes-app-data', JSON.stringify(notes));
    }, [notes]);

    let menu;
    if (props.user === "") {
        menu =  ( <h3 className="title">{'You are not logged in' }</h3>)
    } else {
        menu = (    <div>
                        <div className="add-more">
                            <Link to="/expenses/new" className="button is-link is-medium add-more-button">+</Link>
                        </div>
                        {/*<h3 className="title is-3 center-text">{'Hello ' + props.user.user_name}</h3>*/}
                     </div>)
    }

    const addNote = (text) => {
        let id = 0;
        if (notes.length > 0)  {
            id = notes[notes.length - 1].id + 1;
        }
        const date = new Date();
        const newNote = {
            id: id,
            text: text,
            date: date.toLocaleDateString()
        };
        const newNotes = [...notes, newNote];
        setNotes(newNotes);
    };

    const deleteNote = (deleteNote) => {
        const newNotes = notes.filter(note => note.id !== deleteNote.id);
        setNotes(newNotes);
    };

    const addCheckNote = (itemList) => {
        console.log("Add new empty checklist in home");
        const date = new Date();
        let newList = {
            id: giveId(),
            type: "checklist",
            list: [],
            date: date.toLocaleDateString()
        }
        const newLists = [...checkLists, newList];
        setCheckLists(newLists);
    };

    const updateCheckNote = (itemList) => {
        const newChecklists = checkLists.map(checkList => {
            if (checkList.id === itemList.id) {
                checkList.list = itemList.list
            }
            return checkList
        });
        setCheckLists(newChecklists);
    };

    const searchNote = (text) => {
        setSearchQuery(text);
    };

    let notesToDisplay = notes;
    if (searchQuery) {
        notesToDisplay = searchKeywordNotes(notes, searchQuery);
    };

    return (
        <div>
            {menu}
            <div className="notes-list-container">
                <SearchBox value={searchQuery} onChange={searchNote}/>
                <NotesList
                    notes={notesToDisplay}
                    handleAddNote={addNote}
                    handleDeleteNote={deleteNote}
                />
                <CheckListsList
                    checkLists={checkLists}
                    handleUpdateCheckList={updateCheckNote}
                />
                <button className="save-button button is-link is-light" onClick={addCheckNote}><span
                    className="mdi mdi-content-save"/> &nbsp; Add Checklist</button>
            </div>
        </div>
    );
};

export default Home;
import React, {useState} from "react";
import Link from "react-router-dom/Link";
import NotesList from "./notesList";

const Home = (props) => {
    const [notes, setNotes] = useState([
        {
        id: 0,
        text: "This is my first note!",
        date: "15/04/2021"
        },
        {
        id: 1,
        text: "This is my second note!",
        date: "16/04/2021"
        },
        {
        id: 2,
        text: "This is my third note!",
        date: "17/04/2021"
        },
    ]);

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
    return (
        <div>
            {menu}
            <div className="notes-list-container">
                <NotesList notes={ notes }/>
            </div>
        </div>
    );
};

export default Home;
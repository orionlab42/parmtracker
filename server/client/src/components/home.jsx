import React from "react";
import Link from "react-router-dom/Link";
import ListBox from "./listBox";

const Home = (props) => {
    let menu;
    if (props.user === "") {
        menu =  ( <h3 className="title">{'You are not logged in' }</h3>)
    } else {
        menu = (    <div>
                        <div className="add-more">
                            <Link to="/expenses/new" className="button is-link is-medium add-more-button">+</Link>
                        </div>
                     </div>)
    }

    return (
        <div>
            {menu}
            <ListBox user={props.user}/>
        </div>
    );
};

export default Home;
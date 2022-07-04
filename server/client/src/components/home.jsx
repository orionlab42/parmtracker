import React, {useState} from "react";
import Link from "react-router-dom/Link";
import ListBox from "./listBox";
import DatePicker from "react-datepicker";

const Home = (props) => {
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
            <ListBox/>
        </div>
    );
};

export default Home;
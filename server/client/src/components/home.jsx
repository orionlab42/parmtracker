import React from "react";
import Link from "react-router-dom/Link";

const Home = (props) => {
    let menu;
    if (props.user === "") {
        menu =  ( <h3 className="title">{'You are not logged in' }</h3>)
    } else {
        menu = (<h3 className="title">{'Hello ' + props.user.user_name}</h3>)
    }
    return (
        <div>
            <h1 className="title">Home</h1>
            {menu}
            <div className="add-more">
                <Link to="/expenses/new" className="button is-link is-medium add-more-button">+</Link>
            </div>
        </div>
    );
};

export default Home;
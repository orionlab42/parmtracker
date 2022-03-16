import React from "react";
import NavLink from "react-router-dom/NavLink";

const Home = (props) => {
    let menu;
    if (props.user === '') {
        menu =  ( <h3 className="title">{'You are not logged in' }</h3>)
    } else {
        menu = (<h3 className="title">{'Hello ' + props.user.user_name }</h3>)
    }
    return (
        <div>
            <h1 className="title">Home</h1>
            {menu}
        </div>
    );
};

export default Home;
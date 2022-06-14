import React from "react";
import UserColor from "./userColor";
import UserDarkMode from "./common/userDarkMode";


const Settings = (props) => {
    return (
        <div>
            <h1 className="title is-3 center-text">Settings</h1>
            <UserDarkMode user={props.user} onChange={props.onChange}/>
            <UserColor user={props.user}/>
        </div>
    );
};

export default Settings;
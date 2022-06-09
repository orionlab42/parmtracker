import React, { useState } from "react";
import UserColorForm from "./userColorForm";
import Link from "react-router-dom/Link";
import {Route} from "react-router-dom";


const Settings = (props) => {
    const [addButtonToggle, setAddButtonToggle] = useState(false);
    const [currentColor, setCurrentColor] = useState("");

    let currentColorDisplay;
    let user = props.user;
    if (user !== "") {
        currentColorDisplay = (
            <div className="user-color" style={{backgroundColor:  user.user_color}}>
                <h4 className="title is-5 center-text">Currently saved color:</h4>
            </div>
        )
    }

    const handleToggle = () => {
        let addButtonToggleChanged;
        addButtonToggleChanged = !addButtonToggle;
        setAddButtonToggle(addButtonToggleChanged);
    }

    const passData = (data) => {
        setCurrentColor(data);
        console.log("Color2", data);
        // window.location = '/client/settings';
    };

    return (
        <div>
            <h1 className="title is-3 center-text">Settings</h1>
            <h4 className="title is-5 center-text">Dark mode:</h4>
            <div className="dark-mode-toggle">
                <input type="checkbox" id="switch"/>
                <div className="app">
                    <div className="body">
                        <div className="content">
                            <label htmlFor="switch">
                                <div className="toggle"></div>
                                <div className="names">
                                    <p className="light">Light</p>
                                    <p className="dark">Dark</p>
                                </div>
                            </label>
                        </div>
                    </div>
                </div>
            </div>
            {currentColorDisplay}
            <div className="add-new-color">
                <div className="add-new-color-button">
                    <Link to="/settings/new-color"
                          className="button is-link is-medium add-more-button"
                          onClick={handleToggle}
                    >{addButtonToggle ? "x" : "+"}</Link>
                </div>
                {addButtonToggle &&
                <Route
                    path="/settings/new-color"
                    render={(props) => (<UserColorForm {...props} user={user} passData={passData}/>)}
                />}
            </div>
        </div>
    );
};

export default Settings;
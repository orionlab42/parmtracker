import React, {useEffect, useState} from "react";
import UserColorForm from "./userColorForm";
import Link from "react-router-dom/Link";
import {Route} from "react-router-dom";
import {updateUserSettings} from "../services/userService";


const Settings = (props) => {
    const [addButtonToggle, setAddButtonToggle] = useState(false);
    const [darkModeToggle, setDarkModeToggle] = useState(false);
    const [currentColor, setCurrentColor] = useState('#fff');

    useEffect(() => {
        setDarkModeToggle(props.user.dark_mode);
    }, [props.user.dark_mode]);

    useEffect(() => {
        console.log("color 1: ", props.user.user_color);
        console.log("color 2: ", currentColor);
        console.log("Props 1: ", props);
        setCurrentColor(props.user.user_color);
        console.log("color 3: ", props.user.user_color);
        console.log("color 4: ", currentColor);
    }, []);

    useEffect( () => {
        async function setUserColor() {
            let user = props.user
            user.user_color = currentColor;
            await updateUserSettings(user);

            console.log("color 5: ", props.user.user_color);
            console.log("color 6: ", currentColor);
        }
        setUserColor();

    }, [currentColor]);

    const handleChangeComplete = (color) => {
        const newColor = color.hex;
        setCurrentColor(newColor);
        console.log("color 7: ", props.user.user_color);
        console.log("color 8: ", currentColor);
    };

    let currentColorDisplay;
    console.log("Props 2: ", props);
    if (props.user !== "") {
        console.log("color 9: ", props.user.user_color);
        console.log("color 10: ", currentColor);
        console.log("Props 3: ", props);
        currentColorDisplay = (
            <div className="user-color" style={{backgroundColor:  currentColor}}>
                <h4 className="title is-5 center-text">Currently saved color: {currentColor} </h4>
            </div>
        )
    }

    const handleToggle = () => {
        let addButtonToggleChanged;
        addButtonToggleChanged = !addButtonToggle;
        setAddButtonToggle(addButtonToggleChanged);
        // window.location = "/client/settings";
    }

    const handleDarkModeToggle = () => {
        let darkModeToggleChanged;
        darkModeToggleChanged = !darkModeToggle;
        setDarkModeToggle(darkModeToggleChanged);
    }

    return (
        <div>
            <h1 className="title is-3 center-text">Settings</h1>
            <h4 className="title is-5 center-text">Dark mode:</h4>
            <div className="dark-mode-toggle">
                <input type="checkbox" id="switch" onChange={props.onChange} onClick={handleDarkModeToggle} checked={darkModeToggle}/>
                <div className="toggle-body">
                    <div className="toggle-container">
                        <label htmlFor="switch">
                            <div className="toggle"/>
                            <div className="names">
                                <p className="light">Light</p>
                                <p className="dark">Dark</p>
                            </div>
                        </label>
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
                <Route path="/settings/new-color" render={ () => (<UserColorForm currentColor={currentColor} onChangeComplete={handleChangeComplete}/>)}/>}
            </div>
        </div>
    );
};

export default Settings;
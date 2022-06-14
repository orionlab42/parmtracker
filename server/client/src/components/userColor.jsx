import React, {useEffect, useState} from "react";
import ColorForm from "./common/colorForm";
import Link from "react-router-dom/Link";
import {Route} from "react-router-dom";
import {updateUserSettings} from "../services/userService";


const UserColor = (props) => {
    const [addButtonToggle, setAddButtonToggle] = useState(false);
    const [currentColor, setCurrentColor] = useState('');

    useEffect(() => {
        setCurrentColor(props.user.user_color);
    }, [props.user.user_color]);

    const handleChangeComplete = async (color) => {
        let user = props.user;
        user.user_color = color.hex;
        await updateUserSettings(user);
        setCurrentColor(user.user_color);
    };

    const handleToggle = () => {
        let addButtonToggleChanged;
        addButtonToggleChanged = !addButtonToggle;
        setAddButtonToggle(addButtonToggleChanged);
    };

    let currentColorDisplay;
    if (props.user !== "") {
        currentColorDisplay = (
            <div className="user-color">
                <h4 className="title is-5 center-text settings-title">Currently saved color: {currentColor} </h4>
            </div>
        )
    }

    return (
        <div className="settings" style={{backgroundColor:  currentColor}}>
            {currentColorDisplay}
            <div className="add-new-color settings-body">
                <div className="add-new-color-button">
                    <Link to="/settings/new-color"
                          className="button is-link is-medium add-more-button"
                          onClick={handleToggle}>{addButtonToggle ? "x" : "+"}</Link>
                </div>
                {addButtonToggle && <Route path="/settings/new-color" render={ () => (<ColorForm currentColor={currentColor} onChangeComplete={handleChangeComplete}/>)}/>}
            </div>
        </div>
    );
};

export default UserColor;
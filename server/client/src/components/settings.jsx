import React, {useEffect, useState} from "react";
import UserColorForm from "./userColorForm";
import Link from "react-router-dom/Link";
import {Route} from "react-router-dom";


const Settings = (props) => {
    const [addButtonToggle, setAddButtonToggle] = useState(false);

    let currentColor;
    let user = props.user;
    if (user !== "") {
        // console.log("User Color", props.user.user_color);
        currentColor = (
            <div className="user-color" style={{backgroundColor: user.user_color}}>
                <h4 className="title is-5 center-text">Currently saved color:</h4>
            </div>
        )
    }

    const handleToggle = () => {
        let addButtonToggleChanged;
        addButtonToggleChanged = !addButtonToggle;
        setAddButtonToggle(addButtonToggleChanged);
    }
    const handleUpdate = (e) => {
        console.log("hello change");
        console.log(e);

    }

    return (
        <div>
            <h1 className="title is-3 center-text">Settings</h1>
            {currentColor}
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
                    render={(props) => (<UserColorForm {...props} user={user}
                                                      onUpdate={handleUpdate}/>)}
                />}
            </div>

            {/*<SliderPicker*/}
            {/*    color={background}*/}
            {/*    onChangeComplete={handleChangeComplete}*/}
            {/*/>*/}
            {/*<CirclePicker*/}
            {/*    color={background}*/}
            {/*    onChangeComplete={handleChangeComplete}*/}
            {/*/>*/}
        </div>
    );
};

export default Settings;
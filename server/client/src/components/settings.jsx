import React, {useEffect, useState} from "react";
import SelectColor from "./common/selectColor";
import {getEntriesByMonth, getEntriesByWeek} from "../services/chartsService";
import {getCategoryColor} from "../services/categoryService";
import UserColorForm from "./userColorForm";
import { SliderPicker, CirclePicker } from 'react-color';
import Link from "react-router-dom/Link";
import {Route} from "react-router-dom";
import CategoryForm from "./categoryForm";

const Settings = (props) => {
    const [addButtonToggle, setAddButtonToggle] = useState(false);

    let currentColor;
    if (props.user !== "") {
        // console.log("User Color", props.user.user_color);
        currentColor = (
            <div className="user-color" style={{backgroundColor: props.user.user_color}}>
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
                    render={(props) => (<UserColorForm {...props}
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
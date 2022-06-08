import React, {useEffect, useState} from "react";
import SelectColor from "./common/selectColor";
import {getCategoryColor} from "../services/categoryService";
import { CirclePicker } from 'react-color';
import {getEntriesByCategoryAndUser} from "../services/chartsService";
import {saveUserColor} from "../services/userService";

const UserColorForm = (props) => {
    const [newColor, setNewColor] = useState('#fff');

    console.log("Color", newColor);
    useEffect( () => {
        async function setUserColor() {
            let user = props.user
            user.user_color = newColor
            await saveUserColor(user);
        }
        setUserColor();
    }, [newColor]);

    const handleChangeComplete = (color) => {
        setNewColor(color.hex);
    };

    return (
        <div>
            <CirclePicker
                color={newColor}
                onChangeComplete={handleChangeComplete}
            />
        </div>
    );
};

export default UserColorForm;
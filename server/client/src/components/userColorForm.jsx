import React, {useEffect, useState} from "react";
import { CirclePicker } from 'react-color';
import {updateUserSettings} from "../services/userService";

const UserColorForm = (props) => {
    const [newColor, setNewColor] = useState('#fff');

    useEffect( () => {
        async function setUserColor() {
            let user = props.user
            user.user_color = newColor;
            await updateUserSettings(user);
        }
        setUserColor();
    }, [newColor]);

    const handleChangeComplete = (color) => {
        setNewColor(color.hex);
        props.passData(newColor);
        console.log("Color1", newColor);
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
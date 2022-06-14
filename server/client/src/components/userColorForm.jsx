import React from "react";
import { CirclePicker } from 'react-color';

const UserColorForm = ({currentColor, onChangeComplete}) => {
    return (
        <div>
            <CirclePicker
                color={currentColor}
                onChangeComplete={onChangeComplete}
            />
        </div>
    );
};

export default UserColorForm;
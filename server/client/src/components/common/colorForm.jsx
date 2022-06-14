import React from "react";
import { CirclePicker } from 'react-color';

const ColorForm = ({currentColor, onChangeComplete}) => {
    return (
        <div>
            <CirclePicker
                color={currentColor}
                onChangeComplete={onChangeComplete}
            />
        </div>
    );
};

export default ColorForm;
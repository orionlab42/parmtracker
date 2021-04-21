import React from "react";

const SelectColor = ({ name, label, options, error, ...rest }) => {
    return (
        <div className="field">
            <label htmlFor={name} className="label">{label}</label>
            <div className="control">
                <div className="select">
                    <select name={name} id={name} {...rest}>
                        <option value=""/>
                        {options.map((option, index) =>
                            <option
                                key={index}
                                value={option}
                                style={{backgroundColor:option}}>
                                {option}
                            </option>)}
                    </select>
                    {error && <div>{error}</div>}
                </div>
            </div>
        </div>
    );
};



export default SelectColor;
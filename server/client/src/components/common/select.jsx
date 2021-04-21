import React from "react";

const Select = ({ name, label, options, error, ...rest }) => {
    return (
        <div className="field">
            <label htmlFor={name} className="label">{label}</label>
            <div className="control">
                <div className="select">
                    <select name={name} id={name} {...rest}>
                        <option value=""/>
                        {options.map(option => <option key={option.id} value={option.id}>{option.category_name}</option>)}
                    </select>
                    {error && <div>{error}</div>}
                </div>
            </div>
        </div>
    );
};



export default Select;
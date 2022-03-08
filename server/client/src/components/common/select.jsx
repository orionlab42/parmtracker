import React from "react";

const Select = ({ name, label, dbId, dbName, options, error, ...rest }) => {
    return (
        <React.Fragment>
            <label htmlFor={name} className="label">{label}</label>
            <div className="control" id={name}>
                <div className="select">
                    <select name={name} id={name} {...rest}>
                        <option value=""/>
                        {options.map(option => <option key={option[dbId]} value={option[dbId]}>{option[dbName]}</option>)}
                    </select>
                    {error && <div>{error}</div>}
                </div>
            </div>
        </React.Fragment>
    );
};

export default Select;
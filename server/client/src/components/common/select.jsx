import React from "react";

const Select = ({ name, label, dbId, dbName, dbStyle, options, error, ...rest }) => {
    return (
        <React.Fragment>
            <label htmlFor={name} className="label">{label}</label>
            <div className="control" id={name}>
                <div className="select">
                    <select name={name} id={name} {...rest}>
                           {options.map(option => <option
                                                        key={option[dbId]}
                                                        value={option[dbId]}
                                                        style={{backgroundColor:option[dbStyle]}}
                                                >{option[dbName]}</option>)}
                    </select>
                    {error && <div>{error}</div>}
                </div>
            </div>
        </React.Fragment>
    );
};

export default Select;

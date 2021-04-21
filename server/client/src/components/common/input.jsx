import React from "react";


// stateless functional component
const Input = ({ name, label, error, ...rest }) => {
    return (
        <div className="field">
            <label htmlFor={name} className="label">{label}</label>
            <div className="control">
                <input
                    {...rest}
                    name={name}
                    id={name}
                    className="input"/>
                {error && <div>{error}</div>}
            </div>
        </div>
    );
};

export default Input;
import React from "react";


// stateless functional component
const Checkbox = ({ name, label, error, ...rest }) => {
    return (
        <div className="field">
            <label htmlFor={name} className="checkbox label">
                <input
                    name={name}
                    id={name}
                    type="checkbox"
                    {...rest}/>
            &nbsp;&nbsp;{label}</label>
            {error && <div>{error}</div>}
        </div>
    );
};

export default Checkbox;


import React from "react";


// stateless functional component
const Checkbox = ({ name, label, error, ...rest }) => {
    return (
        <React.Fragment>
            <label htmlFor={name} className="checkbox label">
                <input
                    name={name}
                    id={name}
                    type="checkbox"
                    {...rest}/>
            &nbsp;&nbsp;{label}</label>
            {error && <div>{error}</div>}
        </React.Fragment>
    );
};

export default Checkbox;

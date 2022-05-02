import React from "react";


// stateless functional component
const Checkbox = ({ name, label, error, ...rest }) => {
    return (
        <React.Fragment>
            <label htmlFor={name} className="checkbox label" id={name}>
                <input
                    name={name}
                    id={name}
                    type="checkbox"
                    defaultChecked={true}
                    {...rest}/>
            &nbsp;&nbsp;{label}</label>
            {error && <div>{error}</div>}
        </React.Fragment>
    );
};

export default Checkbox;


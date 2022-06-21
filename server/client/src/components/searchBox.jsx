import React from "react";

const SearchBox = ({ value, onChange}) => {
    return (
        <div className="container search-box">
            <div className="control has-icons-left">
                <span className="icon is-small is-left"><i className="mdi mdi-magnify"/></span>
                <input
                    type="text"
                    name="query"
                    className="input "
                    placeholder="Search..."
                    value={value}
                    onChange={e => onChange(e.currentTarget.value)}
                />
            </div>

        </div>
    );
}


export default SearchBox;
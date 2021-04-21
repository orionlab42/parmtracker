import React from "react";

const SearchBox = ({ value, onChange}) => {
    return (
        <div className="container search-box">
            <div className="control has-icons-right">
                <input
                    type="text"
                    name="query"
                    className="input "
                    placeholder="Search..."
                    value={value}
                    onChange={e => onChange(e.currentTarget.value)}
                />
                <span className="icon is-small is-right"><i className="fa fa-search"/></span>
            </div>

        </div>
    );
}


export default SearchBox;
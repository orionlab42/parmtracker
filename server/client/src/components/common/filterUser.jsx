import React from "react";

function getUserName(items, selectedItem) {
    const emptyUser = {user_name: ""};
    const user = items.filter(u => u.user_id === selectedItem);
    if (user.length === 0) return emptyUser;
    return user[0].user_name;
}

const FilterUser = ({selectedItem, items, onItemSelect, textProperty, valueProperty}) => {
    return (
        <div className="dropdown is-hoverable filter-position filter-user">
            <div className="dropdown-trigger">
                <button className="button" aria-haspopup="true" aria-controls="dropdown-menu4">
                    <span>Filter by user: {getUserName(items, selectedItem)}</span>
                </button>
            </div>
            <div className="dropdown-menu" id="dropdown-menu4" role="menu">
                <div className="dropdown-content">
                    {items.map(filter =>
                        <button
                            style={{backgroundColor: filter.user_color}}
                            key={filter.user_id} className={filter.user_id === selectedItem ? "dropdown-item is-active" : "dropdown-item"}
                            onClick={() => onItemSelect(filter.user_id)}>
                            {filter.user_name}
                        </button>)}
                </div>
            </div>
        </div>
    );
}
export default FilterUser;
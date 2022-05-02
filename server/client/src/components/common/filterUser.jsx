import React from "react";


const FilterUser = ({selectedItem, items, onItemSelect, textProperty, valueProperty}) => {
    return (
        <div className="dropdown is-hoverable filter-position filter-user">
            <div className="dropdown-trigger">
                <button className="button" aria-haspopup="true" aria-controls="dropdown-menu4">
                    <span>Filter by User</span>
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
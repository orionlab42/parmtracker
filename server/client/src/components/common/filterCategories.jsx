import React from "react";


const FilterCategory = ({selectedItem, items, onItemSelect, textProperty, valueProperty}) => {
    return (
        <div className="dropdown is-hoverable filter-position filter-category">
            <div className="dropdown-trigger">
                <button className="button" aria-haspopup="true" aria-controls="dropdown-menu4">
                    <span>Filter by Category</span>
                </button>
            </div>
            <div className="dropdown-menu" id="dropdown-menu4" role="menu">
                <div className="dropdown-content">
                    {items.map(filter =>
                        <button
                            style={{backgroundColor: filter.category_color}}
                            key={filter.id} className={filter.id === selectedItem ? "dropdown-item is-active" : "dropdown-item"}
                            onClick={() => onItemSelect(filter.id)}>
                            <span className={["mdi", filter.category_icon].join(" ")}/>
                            {filter.category_name}
                        </button>)}
                </div>
            </div>
        </div>
    );
}

export default FilterCategory;
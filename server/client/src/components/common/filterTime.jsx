import React from "react";


const FilterTime = props => {
    const timeLength = ['Get all entries','Last week','This month','Last month']
    return (
        <div className="dropdown is-hoverable filter-position filter-time">
            <div className="dropdown-trigger">
                <button className="button" aria-haspopup="true" aria-controls="dropdown-menu4">
                    <span>Filter by Time</span>
                </button>
            </div>
            <div className="dropdown-menu" id="dropdown-menu4" role="menu">
                <div className="dropdown-content">
                    {timeLength.map(filter =>
                        <button key={filter} className={filter === props.currentTimeFilter ? "dropdown-item is-active" : "dropdown-item"}
                            onClick={() => props.onFilterChange(filter)}>{filter}
                        </button>)}
                </div>
            </div>
        </div>
    );
}

export default FilterTime;
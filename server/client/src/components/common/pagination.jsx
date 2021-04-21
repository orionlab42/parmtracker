import React from "react";
import PropTypes from 'prop-types';
import _ from 'lodash';

const Pagination = props => {
    const { itemsCount, pageSize, currentPage, onPageChange } = props;
    const pagesCount = Math.ceil(itemsCount / pageSize);
    if (pagesCount === 1) return null;
    const pages = _.range(1,pagesCount + 1);

    return(
    <nav className="pagination is-centered" role="navigation" aria-label="pagination">
        <ul className="pagination-list">
            {pages.map(page =>
                <li key={page}>
                    <button
                        className={ page === currentPage ? 'pagination-link is-current': 'pagination-link'}
                        onClick={() => onPageChange(page)}
                    >{page}</button>
                </li>)}
        </ul>
    </nav>
    );
};

Pagination.propTypes = {
    itemsCount: PropTypes.number.isRequired,
    pageSize: PropTypes.number.isRequired,
    currentPage: PropTypes.number.isRequired,
    onPageChange: PropTypes.func.isRequired
};

export default Pagination;
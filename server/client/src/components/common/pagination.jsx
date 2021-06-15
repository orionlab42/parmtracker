import React from "react";
import PropTypes from 'prop-types';
import _ from 'lodash';

const Pagination = props => {
    const {itemsCount, pageSize, currentPage, onPageChange} = props;
    const pagesCount = Math.ceil(itemsCount / pageSize);
    if (pagesCount === 1) return null;
    const pages = _.range(1, pagesCount + 1);
    function renderPrev() {
        return <button
            className="pagination-previous"
            onClick={() => currentPage === 1 ? onPageChange(pagesCount) : onPageChange(currentPage - 1)}>Previous
        </button>
    }

    function renderFirst() {
        if (currentPage <= 2) {
            return null;
        }
        return <React.Fragment>
            <li key={1}>
                <button
                    className='pagination-link'
                    onClick={() => onPageChange(1)}
                >1
                </button>
            </li>
        </React.Fragment>;
    }

    function renderFirstEllip() {
        if (currentPage <= 3) {
            return null;
        }
        return <li><span className="pagination-ellipsis">&hellip;</span></li>
    }

    function renderLast() {
        if (currentPage >= pagesCount - 1) {
            return null;
        }
        return <React.Fragment>
            <li key={pagesCount}>
                <button
                    className='pagination-link'
                    onClick={() => onPageChange(pagesCount)}
                >{pagesCount}</button>
            </li>
        </React.Fragment>;
    }

    function renderLastEllip() {
        if (currentPage >= pagesCount - 2) {
            return null;
        }
        return <li><span className="pagination-ellipsis">&hellip;</span></li>
    }

    function renderMid() {
        return <React.Fragment>
            {currentPage > 1 ? <li key={currentPage - 1}>
                <button className='pagination-link' onClick={() => onPageChange(currentPage - 1)}>{currentPage - 1}</button>
            </li> : null }
            <li key={currentPage}>
                <button className='pagination-link is-current'>{currentPage}</button>
            </li>
            {currentPage < pagesCount ? <li key={currentPage + 1}>
                <button className='pagination-link' onClick={() => onPageChange(currentPage + 1)}>{currentPage + 1}</button>
            </li> : null }
        </React.Fragment>;
    }

    function renderFull() {
        return(
            <ul className="pagination-list">
                { pages.map(page =>
                    <li key={page}>
                        <button
                            className={page === currentPage ? 'pagination-link is-current' : 'pagination-link'}
                            onClick={() => onPageChange(page)}
                        >{page}</button>
                    </li>)
                }
            </ul>
        );
    }

    return (
        <nav className="pagination is-centered" role="navigation" aria-label="pagination">
            {renderPrev()}
            <button
                className="pagination-next"
                onClick={() => currentPage === pagesCount ? onPageChange(1) : onPageChange(currentPage + 1)}>Next
            </button>
            {pagesCount < 8 ? renderFull() : <ul className="pagination-list">
                {renderFirst()}
                {renderFirstEllip()}
                {renderMid()}
                {renderLastEllip()}
                {renderLast()} </ul>}
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
import React, {Component} from "react";
import {getEntries, deleteEntry} from '../services/entryService';
import {getCategories} from '../services/categoryService';
import ExpensesTable from './expensesTable';
import Pagination from './common/pagination';
import {paginate} from '../utils/paginate';
import FilterTime from "./common/filterTime";
import FilterCategory from "./common/filterCategories";
import {filterByCategory, filterByTime} from "../utils/filters";
import {searchKeyword} from "../utils/search";
import _ from 'lodash';
import Link from "react-router-dom/Link";
import SearchBox from "./searchBox";
import {toast} from "react-toastify";


class Entries extends Component {
    state = {
        entries: [],
        categories: [],
        currentPage: 1,
        currentTimeFilter: "Get all entries",
        selectedCategory: 0,
        searchQuery: '',
        pageSize: 4,
        sortColumn: {path: 'entry_date', order: 'desc'}
    };

    async componentDidMount() {
        const { data } = await getCategories();
        const categories = [{id:0 , category_name: "Get all entries"}, ...data];

        const { data: entries } = await getEntries();
        this.setState({entries, categories});
    }

    totalCalculation = entries => {
        let total = 0;
        // eslint-disable-next-line array-callback-return
        entries.map(entry => total += entry.amount)
        return total
    }

    handleDelete = async entryToDelete => {
        const originalEntries = this.state.entries;
        const entries = this.state.entries.filter(entry => entry.id !== entryToDelete.id);
        this.setState({entries})
        try {
            await deleteEntry(entryToDelete.id)
        }
        catch (ex) {
            console.log('HANDLE DELETE CATCH BLOCK.');
            if (ex.response && ex.response.status === 404)
                toast('This entry has already been deleted.');
            this.setState({entries: originalEntries});
        }
    };

    handlePageChange = page => {
        this.setState({currentPage: page});
    };

    handleTimeFilterChange = time => {
        this.setState({currentTimeFilter: time, searchQuery: "", currentPage: 1});
    };

    handleCategoryFilterChange = category => {
        this.setState({selectedCategory: category, searchQuery: "", currentPage: 1});
    };

    handleSort = sortColumn => {
        this.setState({sortColumn})
    };

    handleSearch = query => {
        this.setState({searchQuery: query, currentPage: 1})
    };

    getPagedData = () => {
        const {
            pageSize,
            currentPage,
            sortColumn,
            searchQuery,
            entries: allEntries
        } = this.state;

        // filtering
        const entriesFilteredByTime = filterByTime(allEntries, this.state.currentTimeFilter);
        const entriesFilteredByCategory = filterByCategory(allEntries, this.state.selectedCategory);
        let filtered = entriesFilteredByTime.filter(x => entriesFilteredByCategory.includes(x));

        // searching
        const entriesSearched = searchKeyword(allEntries, searchQuery);

        // sorting
        let sorted;
        if (!searchQuery) {
            sorted = _.orderBy(filtered, [sortColumn.path], [sortColumn.order]);
        } else {
            sorted = _.orderBy(entriesSearched, [sortColumn.path], [sortColumn.order]);
        }

        const entries = paginate(sorted, currentPage, pageSize)
        return {totalCount: sorted.length, entries, total: sorted}
    };

    render() {
        const {
            pageSize,
            currentPage,
            sortColumn,
            searchQuery,
            categories,
        } = this.state;

        if (this.state.entries.length === 0) return <h5 className="title is-5 center-text">There are no entries!</h5>
        const {totalCount, entries, total} = this.getPagedData();
        // const {history} = this.props;
        return (

            <div className="container">
                <h3 className="title is-3 center-text">Expenses</h3>
                <div className="add-more">
                    <Link to="/expenses/new" className="button is-link is-medium add-more-button">+</Link>
                </div>

                <div className="filters">

                    <SearchBox
                        value={searchQuery}
                        onChange={this.handleSearch}
                    />
                    <FilterTime onFilterChange={this.handleTimeFilterChange}
                                currentTimeFilter={this.state.currentTimeFilter}
                    />
                    <FilterCategory
                        items={this.state.categories}
                        selectedItem={this.state.selectedCategory}
                        onItemSelect={this.handleCategoryFilterChange}
                    />
                    <div className="add-category">
                        <Link to="/categories" className="button is-link is-light add-category-button">
                            <span className="mdi mdi-square-edit-outline" title="Edit categories" />
                        </Link>
                    </div>
                </div>
                <ExpensesTable
                    entries={entries}
                    categories={categories}
                    sortColumn={sortColumn}
                    onDelete={this.handleDelete}
                    onLike={this.handleLike}
                    onSort={this.handleSort}
                />

                <Pagination itemsCount={totalCount}
                            pageSize={pageSize}
                            currentPage={currentPage}
                            onPageChange={this.handlePageChange}
                />
                <h5 className="title is-5 center-text total-text">There are {totalCount} entries. <br/> Total amount of expenses is {this.totalCalculation(total)} €. </h5>
            </div>
        );
    }
}

export default Entries;
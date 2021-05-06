import React, {Component} from "react";
import Table from "./common/table";
import Link from "react-router-dom/Link";

const emptyCategory = {category_name: "", category_color: "", category_icon: ""};

class ExpensesTable extends Component {
    columns = [
        {path: 'id', label: 'Id'},
        {path: 'name', label: 'Name', content: entry => <Link to={`/expenses/${entry.id}`}>{entry.entry_name}</Link>},
        {path: 'amount', label: 'Amount'},
        {path: 'category', label: 'Category', content: entry => (
            <span style={{backgroundColor: this.getCategoryFromDb(entry.category).category_color}} className="tag">
                <span className={["mdi", this.getCategoryFromDb(entry.category).category_icon].join(" ")}/>
                {this.getCategoryFromDb(entry.category).category_name}
            </span>)},
        {path: 'date', label: 'Date', content: entry => (<span>{this.getTimeFormat(entry.entry_date)}</span>)},
        {key:'delete', content: entry => (
                // eslint-disable-next-line no-restricted-globals
                <button onClick={() => confirm("Are you sure you want to delete this entry?") ? this.props.onDelete(entry) : ""}
                        className="button is-link is-light">-
                </button>
            )
        },
    ];

    getTimeFormat = time => {
        return new Date(time).toDateString()
    };

    getCategoryFromDb = id => {
        const {categories} = this.props;
        const category = categories.filter(cat => cat.id === id);
        if (category.length === 0) return emptyCategory;
        return category[0];
    }

    render() {
        const {entries, onSort, sortColumn} = this.props;
        return (
            <Table
                columns={this.columns}
                data={entries}
                sortColumn={sortColumn}
                onSort={onSort}
            />
        );
    }
}


export default ExpensesTable;
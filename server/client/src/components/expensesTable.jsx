import React, {Component} from "react";
import Table from "./common/table";
import Link from "react-router-dom/Link";

const emptyCategory = {category_name: "", category_color: "", category_icon: ""};
const emptyUser = {user_name: "", user_color: ""}

class ExpensesTable extends Component {
    columns = [
        {path: 'id', label: 'Id',content: entry => (
                <span style={{backgroundColor: this.getUserFromDb(entry.user_id).user_color}} className="tag is-white">
                    {entry.id}
                    &nbsp;
                    <span className={["mdi", this.getShared(entry.shared)].join(" ")}/>
            </span>)},
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
                        className="button is-link is-light"> <span className="mdi mdi-trash-can-outline"/>
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

    getUserFromDb = id => {
        const {users} = this.props;
        const user = users.filter(u => u.user_id === id);
        if (user.length === 0) return emptyUser;
        return user[0];
    }
     getShared = shared => {
        if (shared === true) {
            return "mdi-account-multiple-check"
        }
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
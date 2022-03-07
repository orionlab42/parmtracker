import React from "react";
import Joi from 'joi-browser';
import Form from './common/form';
import {getCategories} from "../services/categoryService";
import {getEntry, saveEntry, duplicateEntry} from "../services/entryService";
import DatePicker from 'react-datepicker';

import "react-datepicker/dist/react-datepicker.css";

class EntryForm extends Form {
    state = {
        data: {
            entry_name: '',
            amount: '',
            category: '',
            shared: false,
            entry_date: new Date(),
        },
        categories: [],
        errors: {}
    };

    schema = {
        id: Joi.number(),
        entry_name: Joi.string()
            .required()
            .label('Title'),
        amount: Joi.number()
            .min(0)
            .required()
            .label('Amount'),
        category: Joi.number()
            .required()
            .label('Category'),
        shared: Joi
            .bool()
            .default(false)
            .label('Shared'),
        entry_date: Joi.date(),
    };

    populateCategories = async () => {
        const { data } = await getCategories();
        const categories = [...data];
        this.setState({categories});
    }

    populateEntry = async () => {
        try {
            const entryId = this.props.match.params.id;
            if (entryId === "new") return;
            const {data: entry}= await getEntry(entryId);
            this.setState({data: this.mapToViewModel(entry)});
        }
        catch (ex) {
            if (ex.response && ex.response === 404)
                this.props.history.replace("/not-found");
        }
    }

    async componentDidMount() {
        await this.populateCategories();
        await this.populateEntry();
    }

    mapToViewModel(entry) {
        return {
            id: entry.id,
            entry_name: entry.entry_name,
            amount: entry.amount,
            category: entry.category,
            shared: entry.shared,
            entry_date: Date.parse(entry.entry_date),
        };
    }

    doSubmit = async () => {
        await saveEntry(this.state.data);
        this.props.history.push("/expenses");
        // Call the server
        console.log('Submitted');
    };

    handleDateChange = (newDate) => {
        const data = {...this.state.data};
        data.entry_date = newDate;
        this.setState({data});
    };

    doDuplicate = async () => {
        await duplicateEntry(this.state.data);
        this.props.history.push("/expenses");
    };

    renderDuplicateButton = () => {
        if (this.state.data.id === undefined) {
            return <h1 className="title center-text">New entry</h1>;
        }
        return <div className="field duplicate-button">
                    <button className="button is-link center-text"
                            onClick={this.doDuplicate}>
                        Duplicate
                    </button></div>;
    }

    render() {
        return (
            <div className="container">
                {this.renderDuplicateButton()}
                <form onSubmit={this.handleSubmit}>

                    {this.renderInput('entry_name', 'Name')}
                    {this.renderInput('amount', 'Amount')}
                    {this.renderSelect('category', 'Category', this.state.categories)}
                    {this.renderCheckbox('shared', 'Shared')}
                    <div className="field">
                        <label htmlFor="Date" className="label">Date</label>
                        <div className="control picker">
                            <DatePicker
                                id="date"
                                selected={ this.state.data.entry_date}
                                onChange={this.handleDateChange}
                                name="Date"
                                dateFormat={"dd/MM/yyyy"}
                            />
                        </div>
                    </div>
                    {this.renderButton("Save")}
                </form>
            </div>
        );
    }
}

export default EntryForm;


import React from "react";
import Joi from 'joi-browser';
import Form from './common/form';
import {getCategories} from "../services/categoryService";
import {getEntry, saveEntry} from "../services/entryService";

class EntryForm extends Form {
    state = {
        data: {
            title: '',
            amount: '',
            category: '',
            shop: ''
        },
        categories: [],
        errors: {}
    };

    schema = {
        id: Joi.number(),
        title: Joi.string()
            .required()
            .label('Title'),
        amount: Joi.number()
            .min(0)
            .required()
            .label('Amount'),
        category: Joi.number()
            .required()
            .label('Category'),
        shop: Joi.string()
            .required()
            .label('Shop'),
        date: Joi.date(),
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
            title: entry.title,
            amount: entry.amount,
            category: entry.category,
            shop: entry.shop
        };
    }

    doSubmit = async () => {
        await saveEntry(this.state.data);
        this.props.history.push("/expenses");
        // Call the server
        console.log('Submitted');
    };

    render() {
        return (
            <div className="container">
                <h1 className="title center-text">New entry</h1>
                <form onSubmit={this.handleSubmit}>
                    {this.renderInput('title', 'Title')}
                    {this.renderInput('amount', 'Amount')}
                    {this.renderSelect('category', 'Category', this.state.categories)}
                    {this.renderInput('shop', 'Shop')}
                    {this.renderButton("Save")}
                </form>
            </div>
        );
    }
}

export default EntryForm;
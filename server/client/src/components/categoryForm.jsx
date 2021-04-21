import React from "react";
import Joi from 'joi-browser';
import Form from './common/form';
import { getCategoryColor, saveCategory} from "../services/categoryService";


class CategoryForm extends Form {
    state = {
        data: {
            category_name: '',
            category_color: ''
        },
        category_color: getCategoryColor(),
        errors: {}
    };

    schema = {
        id: Joi.number(),
        category_name: Joi.string()
            .required()
            .label('New category'),
        category_color: Joi
            .label('Color'),
        date: Joi.date(),
    };

    doSubmit = async () => {
        await saveCategory(this.state.data);
        this.props.history.push("/categories");
        this.props.onUpdate();
    };

    render() {
        return (
            <div className="container">
                <h1 className="title center-text">New entry</h1>
                <form onSubmit={this.handleSubmit}>
                    {this.renderInput('category_name', 'New category')}
                    {this.renderSelectColor('category_color', 'Color', this.state.category_color)}
                    {this.renderButton("Save")}
                </form>
            </div>
        );
    }
}

export default CategoryForm;
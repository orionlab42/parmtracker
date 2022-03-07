import React, {Component} from "react";
import Joi from "joi-browser";
import Input from "./input";
import Select from "./select";
import SelectColor from "./selectColor";
import Checkbox from "./checkbox";

class Form extends Component {
    state = {
        data: {},
        errors: {}
    };

    validate = () => {
        const options = {abortEarly: false};
        const {error} = Joi.validate(this.state.data, this.schema, options);
        if (!error) return null;
        const errors = {};
        for (let item of error.details) errors[item.path[0]] = item.message;
        return errors;
    };

    validateProperty = ({name, value}) => {
        const obj = {[name]: value};
        const schema = {[name]: this.schema[name]};
        const {error} = Joi.validate(obj, schema);
        return error ? error.details[0].message : null;
    };

    handleSubmit = e => {
        e.preventDefault(); // prevent the submission of the form instead call the server and save the
        // changes and redirect the user to a different page

        const errors = this.validate();
        this.setState({errors: errors || {}});
        if (errors) return;
        this.doSubmit();
    };

    handleChange = ({currentTarget: input}) => {
        const errors = {...this.state.errors};
        const errorMessage = this.validateProperty(input);
        if (errorMessage) errors[input.name] = errorMessage;
        else delete errors[input.name];

        const data = {...this.state.data};
        data[input.name] = input.value;
        this.setState({data, errors});
    };

    handleChecked = (e) => {
        const data = {...this.state.data};
        data.shared = e.target.checked;
        console.log(data.shared);
        this.setState({data});
    }

    renderButton(label) {
        return (
            <div className="field">
                <div className="control center-text">
                    <button
                        id={label}
                        disabled={this.validate()}
                        className="button is-link">{label}
                    </button>
                </div>
            </div>
        );
    };

    renderInput(name, label, type='text') {
        const {data, errors} = this.state;
        return (
            <Input
                type={type}
                name={name}
                value={data[name]}
                label={label}
                onChange={this.handleChange}
                error={errors[name]}
            />
        );
    };

    renderSelect(name, label, options) {
        const { data, errors } = this.state;
        return (
            <Select
                name={name}
                value={data[name]}
                label={label}
                options={options}
                onChange={this.handleChange}
                errors={errors[name]}
            />
        );
    };

    renderCheckbox(name, label, type='checkbox') {
        const { data, errors } = this.state;
        return (
            <Checkbox
                type={type}
                name={name}
                value={data[name]}
                label={label}
                onChange={this.handleChecked}
                errors={errors[name]}
            />
        );
    };

    renderSelectColor(name, label, options) {
        const { data, errors } = this.state;
        return (
            <SelectColor
                name={name}
                value={data[name]}
                label={label}
                options={options}
                onChange={this.handleChange}
                errors={errors[name]}
            />
        );
    };
}

export default Form;
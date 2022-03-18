import React from "react";
import Joi from 'joi-browser';
import Form from './common/form';
import { register } from '../services/userService';

class RegisterForm extends Form {
    state = {
        data: {username: '', password: '', email: ''},
        errors: {}
    };

    schema = {
        username: Joi.string()
            .required()
            .label('Username'),
        password: Joi.string()
            .required()
            .min(5)
            .label('Password'),
        email: Joi.string()
            .required()
            .email()
            .label('Email')
    };

    doSubmit = async () => {
        try {
            await register(this.state.data);
            window.location = '/client';
        }
        catch (ex) {
            if (ex.response && ex.response.status === 400) {
                const errors = {...this.state.errors};
                errors.username = ex.response.data;
                this.setState({ errors });
            }
        }
        // Call the server
        console.log('Submitted');
    };

    render() {
        return (
            <div className="container">
                <h1 className="title center">Register</h1>
                <form onSubmit={this.handleSubmit}>
                    {this.renderInput('username', 'Username')}
                    {this.renderInput('password', 'Password', 'password')}
                    {this.renderInput('email', 'Email')}
                    {this.renderButton("Register")}
                </form>
            </div>
        );
    }
}

export default RegisterForm;
import React from "react";
import Joi from 'joi-browser';
import Form from './common/form';
import {login} from '../services/userService'

class LoginForm extends Form {
    state = {
        data: {username: '', password: ''},
        errors: {}
    };

    schema = {
        username: Joi.string()
            .required()
            .label('Username'),
        password: Joi.string()
            .required()
            .label('Password')
    };

    doSubmit = async () => {
        try {
            await login(this.state.data);
            window.location = '/client/home'; // this causes a full reload of the application, so App will mount again
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
                <h1 className="title is-3 center-text">Login</h1>
                <form onSubmit={this.handleSubmit}>
                    {this.renderInput('username', 'E-mail')}
                    {this.renderInput('password', 'Password', 'password')}
                    {this.renderButton("Login")}
                </form>
            </div>
        );
    }
}

export default LoginForm;
import React, {Component} from "react";
import {Route, Switch, Redirect} from "react-router-dom";
import NavBar from "./components/navbar";
import Expenses from './components/expenses';
import Categories from './components/categories';
import Incomes from "./components/incomes";
import Overview from "./components/overview";
import NotFound from "./components/notFound";
import LoginForm from "./components/loginForm";
import RegisterForm from "./components/registerForm";
import EntryForm from "./components/entryForm";
import "./App.css";


class App extends Component {

    render() {
        return (
            <React.Fragment>
                <NavBar/>
                <main>
                    <Switch>
                        <Route path="/incomes" component={Incomes}/>
                        <Route path="/expenses/:id" component={EntryForm}/>
                        <Route path="/expenses" component={Expenses}/>
                        <Route path="/categories" component={Categories}/>
                        <Route path="/overview" component={Overview}/>
                        <Route path="/login" component={LoginForm}/>
                        <Route path="/register" component={RegisterForm}/>
                        <Redirect from="/" exact to="/expenses"/>
                        <Route path="/not-found" component={NotFound}/>
                        <Redirect to="/not-found"/>}
                    </Switch>
                </main>
            </React.Fragment>
        );
    }
}

export default App;

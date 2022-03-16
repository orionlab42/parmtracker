import React, {Component} from "react";
import {BrowserRouter, Redirect, Route, Switch} from "react-router-dom";
import jwtDecode from 'jwt-decode';
import NavBar from "./components/navbar";
import Expenses from './components/expenses';
import Categories from './components/categories';
import Home from "./components/home";
import Incomes from "./components/incomes";
import Overview from "./components/overview";
import NotFound from "./components/notFound";
import LoginForm from "./components/loginForm";
import RegisterForm from "./components/registerForm";
import EntryForm from "./components/entryForm";
import "./App.css";
import {getUser} from "./services/userService";

console.log("aaa" + process.env.REACT_APP_BASE_URL);

class App extends Component {
    state = {
        user: []
    };
    async componentDidMount() {
        try {
            const { data: user } = await getUser();
            this.setState({user});
        }
        catch (ex) {}
    }

    handleNameChange = (name) => {
        const user = {...this.state.user};
        user.user_name = name;
        this.setState({user});
    };

    render() {
        return (
            <React.Fragment>
                <BrowserRouter basename={process.env.REACT_APP_BASE_URL}>
                    <div>
                        <NavBar
                            user={this.state.user}
                            onChange={this.handleNameChange}/>
                        <main>
                            <Switch>
                                <Route path="/incomes" component={Incomes}/>
                                <Route path="/expenses/:id" component={EntryForm}/>
                                <Route path="/expenses" component={Expenses}/>
                                <Route path="/categories" component={Categories}/>
                                <Route path="/overview" component={Overview}/>
                                <Route path="/login" component={LoginForm}/>
                                <Route path="/register" component={RegisterForm}/>
                                <Route path="/logout" component={Home}/>
                                <Route path="/" component={() => <Home user={this.state.user}/>}/>
                                <Redirect from="/expense" exact to="/expenses"/>
                                <Route path="/not-found" component={NotFound}/>
                                <Redirect to="/not-found"/>}
                            </Switch>
                        </main>
                    </div>
                </BrowserRouter>
            </React.Fragment>
        );
    }
}

export default App;

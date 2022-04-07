import React, {Component} from "react";
import {BrowserRouter, Redirect, Route, Switch} from "react-router-dom";
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
import ProtectedRoute from "./components/common/protectedRoute";
import "./App.css";
import {getCurrentUser} from "./services/userService";

console.log("aaa" + process.env.REACT_APP_BASE_URL);

class App extends Component {
    state = {
        user: {user_name: "",dark_mode: 0,}
    };

    async componentDidMount() {
        const user = await getCurrentUser();
        this.setState({user});
    }

    render() {
        const user = this.state.user
        return (
            <React.Fragment>
                <BrowserRouter basename={process.env.REACT_APP_BASE_URL}>
                    <div className={user.dark_mode ? "dark-mode" : ""}>
                        <NavBar user={user}/>
                        <main>
                            <Switch>
                                <ProtectedRoute path="/incomes" user={user} component={Incomes}/>
                                {/*<ProtectedRoute path="/expenses/:id" user={user} component={EntryForm}/>*/}
                                <ProtectedRoute path="/expenses/:id" user={user} render={props => <EntryForm {...props} user={user}/>}/>
                                <ProtectedRoute path="/expenses" user={user} render={props => <Expenses {...props} user={user}/>}/>
                                <Redirect from="/expense" exact to="/expenses"/>
                                <ProtectedRoute path="/categories" user={user} component={Categories}/>
                                <ProtectedRoute path="/overview" user={user} component={Overview}/>
                                <Route path="/login" component={LoginForm}/>
                                <Route path="/register" component={RegisterForm}/>
                                <Route path="/" exact render={props => <Home {...props} user={user}/>}/>
                                <Route path="/not-found" component={NotFound}/>
                                <Redirect to="/not-found"/>
                            </Switch>
                        </main>
                    </div>
                </BrowserRouter>
            </React.Fragment>
        );
    }
}

export default App;

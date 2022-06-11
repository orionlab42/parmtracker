import React, {Component} from "react";
import {BrowserRouter, Redirect, Route, Switch} from "react-router-dom";
import NavBar from "./components/navbar";
import Expenses from './components/expenses';
import Categories from './components/categories';
import Home from "./components/home";
import Incomes from "./components/incomes";
import Overview from "./components/overview";
import Settings from "./components/settings";
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


    // needs to be reviewed with pere !!!!
    async componentDidUpdate(prevProps, prevState) {
        if (prevState.user.dark_mode !== this.state.user.dark_mode) {
            console.log('dark mode state has changed.');
        }
        // const user = await getCurrentUser();
        // this.setState({user});
    }


    HandleChange = darkMode => {
        let user = this.state.user;
        user.dark_mode = darkMode;
        this.setState({user});
        console.log("Dark mode in App js", user.dark_mode);
    };

    render() {
        console.log("User", this.state.user);
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
                                <ProtectedRoute path="/settings" user={user} render={props => <Settings {...props} user={user} onChange={this.state.HandleChange}/>}/>
                                <Route path="/login" component={LoginForm}/>
                                <Route path="/register" component={RegisterForm}/>
                                <Route path="/home" exact render={props => <Home {...props} user={user}/>}/>
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

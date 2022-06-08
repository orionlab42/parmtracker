import React from "react";
import NavLink from "react-router-dom/NavLink";
import {logout} from '../services/userService';

const NavBar =  (props) => {
    const  handleLogout = async () => {
        try {
             await logout();
             window.location = "/client/login";
        }
        catch (ex) {}
    };

    let menu, logo;
    if (props.user === "") {
       menu =  (    <React.Fragment>
                        <NavLink className="navbar-item" to="/login">
                            Login
                        </NavLink>
                        <NavLink className="navbar-item" to="/register">
                            Register
                        </NavLink>
                    </React.Fragment>)
        logo = (
            <div>
                <img className="logo-navbar" alt="ParmTracker" src={process.env.PUBLIC_URL + '/img/parm-logo-light-mode.png'}/>
            </div>
        )
    } else {
        menu = (    <React.Fragment>
                    <NavLink className="navbar-item" to="/incomes">
                        <span className="mdi mdi-home-plus-outline"/>
                        Incomes
                    </NavLink>
                    <NavLink className="navbar-item" to="/expenses">
                        <span className="mdi mdi-home-minus-outline"/>
                        Expenses
                    </NavLink>
                    <NavLink className="navbar-item" to="/overview">
                        <span className="mdi mdi-chart-bar"/>
                        Overview
                    </NavLink>
                    <NavLink className="navbar-item" to="/settings">
                        <span className="mdi mdi-account-settings-outline"/>
                        Settings
                    </NavLink>
                    <NavLink className="navbar-item" onClick={handleLogout} to="/login">
                        Logout
                    </NavLink>
                </React.Fragment>)

        props.user.dark_mode    ? logo = (<div><img className="logo-navbar" alt="ParmTracker" src={process.env.PUBLIC_URL + '/img/parm-logo-dark-mode.png'}/></div>)
                                : logo = (<div><img className="logo-navbar" alt="ParmTracker" src={process.env.PUBLIC_URL + '/img/parm-logo-light-mode.png'}/></div>)
    }
    return (
        <nav className="navbar" role="navigation" aria-label="main navigation">
            <div className="navbar-brand">
                {/*<h4 className="title is-4 center-navbar-title">ParmTracker</h4>*/}
                {logo}
                <button className="navbar-burger" aria-label="menu" aria-expanded="false"
                   data-target="navbarBasicExample">
                    <span aria-hidden="true"/>
                    <span aria-hidden="true"/>
                    <span aria-hidden="true"/>
                </button>
            </div>
            <div id="navbarBasicExample" className="navbar-menu">
                <div className="navbar-end">
                    <NavLink className="navbar-item" to="/home">
                        <span className="mdi"/>
                        Home
                    </NavLink>
                    {menu}
                </div>
            </div>
        </nav>
    );
};


export default NavBar;
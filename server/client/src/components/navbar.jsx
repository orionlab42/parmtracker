import React from "react";
import NavLink from "react-router-dom/NavLink";

const NavBar = (props) => {
    return (
        <nav className="navbar" role="navigation" aria-label="main navigation">
            <div className="navbar-brand">
                <h4 className="title is-4 center-navbar-title">ParmTracker</h4>
                <button className="navbar-burger" aria-label="menu" aria-expanded="false"
                   data-target="navbarBasicExample">
                    <span aria-hidden="true"/>
                    <span aria-hidden="true"/>
                    <span aria-hidden="true"/>
                </button>
            </div>

            <div id="navbarBasicExample" className="navbar-menu">
                { /* <div className="navbar-start center">
                    <NavLink className="navbar-item" to="/home">
                        Home
                    </NavLink>
                </div> */ }
                <div className="navbar-end">
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
                    {!props.user && <React.Fragment>
                                        <NavLink className="navbar-item" to="/login">
                                            Login
                                        </NavLink>
                                        <NavLink className="navbar-item" to="/register">
                                            Register
                                        </NavLink>
                    </React.Fragment>}
                </div>
            </div>
        </nav>
    );
};


export default NavBar;
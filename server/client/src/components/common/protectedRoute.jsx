import React from "react";
import {Redirect, Route} from "react-router-dom";

const ProtectedRoute = ({path, user, component: Component, render, ...rest}) => {
    return (
        <Route
            {...rest}
            render={props => {
            if (!user) return <Redirect to="/login"/>;
            return Component ? <Component{...props}/> : render(props);
        }}
        />
    );
}

export default ProtectedRoute;
import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import Home from "./pages/Home";
import App from "./App";

const Router: React.FC = () => {
    return (
        <BrowserRouter>
            <Switch>
                <Route path={"/app"} component={App} exact />
                <Route path={"/"} component={Home} exact />
            </Switch>
        </BrowserRouter>
    );
}

export default Router;

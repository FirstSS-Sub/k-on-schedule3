import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import HomePage from "./components/pages/HomePage";
import App from "./App";

const Router: React.FC = () => {
    return (
        <BrowserRouter>
            <Switch>
                <Route path={"/app"} component={App} exact />
                <Route path={"/"} component={HomePage} exact />
            </Switch>
        </BrowserRouter>
    );
}

export default Router;

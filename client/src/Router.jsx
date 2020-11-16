import React, {Component} from 'react';
import {BrowserRouter, Route, Switch} from "react-router-dom";

// コンテナ読み込み
import ResponsiveDrawer from './templates/ResponsiveDrawer';
// import RouteRelatedBottomNavigation from './/RouteRelatedBottomNavigation';
import Home from './pages/Home';
import Info from './pages/Info';
import Settings from './pages/Settings';

/*
// コンポーネント読み込み
import WrapMainContent from './components/WrapMainContent'
*/
// 共通スタイル読み込み
import './App.css';

// 不明なRouteは全てNotFound
const NotFound = () => {
    return (
        <h2>ページが見つかりません</h2>
    )
}


class Router extends Component {

    render() {
        return (
            <BrowserRouter>
                <div className="App">
                    <ResponsiveDrawer className="ResponsiveDrawer">
                        <Switch>
                            <Route exact path="/" component={Home}/>
                            <Route exact path="/info" component={Info}/>
                            <Route exact path="/settings" component={Settings}/>
                            <Route component={NotFound}/>
                        </Switch>
                    </ResponsiveDrawer>
                </div>
            </BrowserRouter>

        );
    }
}


// React-Router情報取得
export default Router;

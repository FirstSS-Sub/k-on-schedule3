import React, { Component } from 'react';

// コンテナ読み込み
import ResponsiveDrawer from './containers/ResponsiveDrawer';
import RouteRelatedBottomNavigation from './containers/RouteRelatedBottomNavigation';
import Notification from './containers/Notification';
import Home from './containers/Home';
import Info from './containers/Info';
import Settings from './containers/Settings';

// コンポーネント読み込み
import WrapMainContent from './components/WrapMainContent'

// 共通スタイル読み込み
import './App.css';

// Route関連
import {Route, Switch} from 'react-router-dom';
import SignUp from "./containers/SignUp";
import SignIn from "./containers/SignIn";
import Auth from "./Auth";
import Dummy from "./containers/Dummy";

// 不明なRouteは全てNotFound
const NotFound = () => {
  return(
    <h2>ページが見つかりません</h2>
  )
}


class App extends Component {

  render() {
    return (
        <div className="App">
            <Notification/>
            <ResponsiveDrawer className="ResponsiveDrawer">
                <Switch>
                    <Route exact path="/" component={WrapMainContent(Home)} />
                    <Route exact path="/signup" component={WrapMainContent(SignUp)} />
                    <Route exact path="/signin" component={WrapMainContent(SignIn)} />
                    <Route exact path="/info" component={WrapMainContent(Info)}/>
                    <Route exact path="/settings" component={WrapMainContent(Settings)}/>
                    <Auth>
                        <Switch>
                            <Route exact path="/schedule" component={WrapMainContent(Dummy)} />
                            <Route exact path="/group" component={WrapMainContent(Dummy)} />
                        </Switch>
                    </Auth>
                    {/* 上から順に確認されるため、NotFoundは一番最後 */}
                    <Route component={WrapMainContent(NotFound)}/>
                </Switch>
            </ResponsiveDrawer>
            <RouteRelatedBottomNavigation/>
        </div>
    );
  }
}


// React-Router情報取得
export default App;

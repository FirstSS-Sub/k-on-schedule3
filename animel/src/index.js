import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import registerServiceWorker from './registerServiceWorker';

// Material-UI
import {createMuiTheme, MuiThemeProvider} from '@material-ui/core/styles';
import red from '@material-ui/core/colors/red';
import blue from '@material-ui/core/colors/blue';

// bootstrap
import 'bootstrap/dist/css/bootstrap.min.css'

// Redux関連
import {applyMiddleware, compose, createStore} from 'redux';
import {Provider} from 'react-redux';
import reducers from './reducers';

// Router関連
import {BrowserRouter} from 'react-router-dom';

// Redux-Thunk関連（非同期データ取得用）
import thunk from 'redux-thunk'

import App from "./App";


// Redux設定
// const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose; // Chromeのデバック用
const composeEnhancers = compose; // 本番用
const store = createStore(
    reducers,
    composeEnhancers(
        applyMiddleware(thunk),
    )
);

// Material-UIテーマカスタマイズ
const theme = createMuiTheme({
    palette: {
        type: 'light', // light or dark
        primary: red, // primaryのカラー
        secondary: blue, // secondaryのカラー
    },
});


ReactDOM.render(
    <Provider store={store}>
        <MuiThemeProvider theme={theme}>
            <BrowserRouter>
                <App/>
            </BrowserRouter>
        </MuiThemeProvider>
    </Provider>
    , document.getElementById('root'));
registerServiceWorker();

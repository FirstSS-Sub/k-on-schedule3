import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import firebase from "firebase";

// スタイル
const styles = theme => ({
    root: {
        ...theme.mixins.gutters(),
        paddingTop: theme.spacing.unit * 2,
        paddingBottom: theme.spacing.unit * 2,
        margin: 10,
    },
    textLeft: {
        textAlign: 'left',
    },
    paragraph: {
        marginTop: 10,
        marginBottom: 10,
    },
});


class Dummy extends React.Component {

    render() {

        // Material-ui関連
        const { classes } = this.props;

        console.log(firebase.auth().currentUser.uid);
        console.log(process.env.REACT_APP_AIUEO);
        console.log(process.env.REACT_APP_API_KEY);

        return (
            <div className={classes}>
                <h1>制作中ダミー</h1>
            </div>
        );
    }
}

// Material-ui関連
Dummy.propTypes = {
    classes: PropTypes.object.isRequired,
    theme: PropTypes.object.isRequired,
};


// Material-uiのテーマ設定
export default withStyles(styles, { withTheme: true })(Dummy);

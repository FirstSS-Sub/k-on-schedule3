import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';

import {Card, Checkbox} from "@material-ui/core";
import CheckBoxOutlineBlankIcon from "@material-ui/icons/CheckBoxOutlineBlank";
import CheckBoxIcon from "@material-ui/icons/CheckBox";
import Switch from "@material-ui/core/Switch";
import CardActionArea from "@material-ui/core/CardActionArea";

// スタイル
const styles = theme => ({
    root: {
        ...theme.mixins.gutters(),
        paddingTop: theme.spacing.unit,
        paddingBottom: theme.spacing.unit,
        margin: 10,
        maxWidth: 700,
    },
    card: {
        margin: 10,
        maxWidth: 700,
    },
    textCenter: {
        textAlign: 'center',
    },
    textLeft: {
        textAlign: 'left',
        margin: 10,
    },
    textRight: {
        textAlign: 'right',
    },
    cardColor: {
        color: '#FFFFFF',
    },
    button: {
        paddingTop: theme.spacing.unit * 2,
        paddingBottom: theme.spacing.unit,
        margin: 5,
    },
    paragraph: {
        marginTop: 10,
        marginBottom: 10,
    },
});

class Schedule extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            checked: false,
            cardColor: {backgroundColor: "white"},
            data: []
        };

        this.handleChange = this.handleChange.bind(this);
    }

    handleChange() {
        if (!this.state.checked) {
            this.setState({
                checked: !this.state.checked,
                cardColor: {backgroundColor: "#EAFAF1"},
            })
        } else {
            this.setState({
                checked: !this.state.checked,
                cardColor: {backgroundColor: "white"},
            })
        }
    }

    render() {

        // Material-ui関連
        const { classes } = this.props;

        return (
            <div>
                <h2>Animelについて</h2>
                <div className={classes.textLeft}>

                    <Card className={classes.root} elevation={1} color={classes.cardColor}>
                        <div className={classes.textLeft}>20:30-22:00</div>
                        <div className={classes.textRight}>
                            <Switch />
                        </div>
                    </Card>
                    <Card className={classes.card} elevation={1} style={ this.state.cardColor } onClick={ this.handleChange }>
                        <CardActionArea>
                            <div className={classes.textLeft}>11/11(日) 20:30-22:00</div>
                            <div className={classes.textRight}>
                                <Checkbox
                                    checked={ this.state.checked }
                                    icon={<CheckBoxOutlineBlankIcon />}
                                    checkedIcon={<CheckBoxIcon />}
                                />
                            </div>
                        </CardActionArea>
                    </Card>
                    <Card className={classes.card} elevation={1} style={ this.state.cardColor } onClick={ this.handleChange }>
                        <CardActionArea>
                            <div className={classes.textLeft}>11/11(日) 20:30-22:00</div>
                            <div className={classes.textRight}>
                                <Checkbox
                                    checked={ this.state.checked }
                                    icon={<CheckBoxOutlineBlankIcon />}
                                    checkedIcon={<CheckBoxIcon />}
                                />
                            </div>
                        </CardActionArea>
                    </Card>
                </div>
            </div>
        );
    }
}

// Material-ui関連
Schedule.propTypes = {
    classes: PropTypes.object.isRequired,
    theme: PropTypes.object.isRequired,
};


// Material-uiのテーマ設定
export default withStyles(styles, { withTheme: true })(Schedule);

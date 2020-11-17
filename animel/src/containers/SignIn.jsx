import React from 'react';
import { Button, Form, FormGroup, Label, Input, FormFeedback, Spinner } from 'reactstrap';
import { Link, withRouter } from 'react-router-dom'
import { Formik } from 'formik';
import * as Yup from 'yup';
import firebase from '../Firebase';

import { withStyles } from '@material-ui/core/styles';
import FormControl from "@material-ui/core/FormControl";

// スタイル
const styles = theme => ({
    titleImage: {
        width: '100%',
        maxWidth: 700,
    },

    button: {
        marginTop: 30,
        marginBottom: 20,
        fontSize: 16,
        padding: 10,
        width: 250,
    },
    leftIcon: {
        marginRight: theme.spacing.unit,
    },
    rightIcon: {
        marginLeft: theme.spacing.unit,
    },
    root: {
    },

    // Form
    formControl: {
        margin: theme.spacing.unit,
        minWidth: 270,
    },
});

class SignIn extends React.Component {

    state = {
        loading: false, //spinner制御用
    }

    _isMounted = false;

    handleOnSubmit = (values) => {
        //spinner表示開始
        if (this._isMounted) this.setState({ loading: true })
        //サインイン（ログイン）処理
        firebase.auth().signInWithEmailAndPassword(values.email, values.password)
            .then(res => {
                //正常終了時
                this.props.history.push("/");
                if (this._isMounted) this.setState({ loading: false });
            })
            .catch(error => {
                //異常終了時
                if (this._isMounted) this.setState({ loading: false });
                alert(error);
            });

    }

    componentDidMount = () => {
        this._isMounted = true;
    }

    componentWillUnmount = () => {
        this._isMounted = false;
    }

    render() {
        // Material-ui関連
        const { classes } = this.props;

        return (
            <div className="container">
                <form autoComplete="off">
                    <FormControl className={classes.formControl}>
                        <div className="mx-auto" style={{ width: '100%', background: '#eee', padding: 10, marginTop: 20 }}>
                            <p style={{ textAlign: 'center' }}>サインイン</p>
                            <Formik
                                initialValues={{ email: '', password: '' }}
                                onSubmit={(values) => this.handleOnSubmit(values)}
                                validationSchema={Yup.object().shape({
                                    email: Yup.string().email().required(),
                                    password: Yup.string().required(),
                                })}
                            >
                                {
                                    ({ handleSubmit, handleChange, handleBlur, values, errors, touched }) => (
                                        <Form onSubmit={handleSubmit}>
                                            <FormGroup>
                                                <Label for="email">Email</Label>
                                                <Input
                                                    type="email"
                                                    name="email"
                                                    id="email"
                                                    value={values.email}
                                                    onChange={handleChange}
                                                    onBlur={handleBlur}
                                                    invalid={touched.email && errors.email ? true : false}
                                                />
                                                <FormFeedback>
                                                    {errors.email}
                                                </FormFeedback>
                                            </FormGroup>
                                            <FormGroup>
                                                <Label for="password">Password</Label>
                                                <Input
                                                    type="password"
                                                    name="password"
                                                    id="password"
                                                    value={values.password}
                                                    onChange={handleChange}
                                                    onBlur={handleBlur}
                                                    invalid={touched.password && errors.password ? true : false}
                                                />
                                                <FormFeedback>
                                                    {errors.password}
                                                </FormFeedback>
                                            </FormGroup>
                                            <div style={{ textAlign: 'center' }}>
                                                <Button color="primary" type="submit" disabled={this.state.loading}>
                                                    <Spinner size="sm" color="light" style={{ marginRight: 5 }} hidden={!this.state.loading} />
                                                    ログイン
                                                </Button>
                                            </div>
                                        </Form>
                                    )
                                }
                            </Formik>
                        </div>
                        <div className="mx-auto" style={{ width: '100%', padding: 20 }}>
                            <Link to="/signup">新規登録はこちら</Link>
                        </div>
                    </FormControl>
                </form>
            </div>
        );
    }
}

export default withRouter(
    withStyles(styles, { withTheme: true })(SignIn)
);

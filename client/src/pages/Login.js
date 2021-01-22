import React, { useState } from 'react';
import  { Redirect } from 'react-router-dom'
import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import axios from 'axios';

const useStyles = makeStyles((theme) => ({
    paper: {
        marginTop: theme.spacing(8),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    form: {
        width: '100%', // Fix IE 11 issue.
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
        backgroundColor: '#2d2e40'
    },
}));


export default function Login(props){

    // set css classes
    const classes = useStyles();

    const [loginData, setLoginData] = useState({
        username: "",
        password: "",
    })

    const handleChange = (e) => {
        const { id, value } = e.target
        setLoginData(prevLoginData => ({
            ...prevLoginData,
            [id] : value
        }))
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        const payload = {
            "username": loginData.username,
            "password": loginData.password   
        }

        axios.post("http://localhost:8888/login", payload)
            .then(response => {
                if(response.status === 200){
                    // if response is ok, redirect to admin
                    props.history.push('/admin');
                } else {
                    console.log(response);
                }
            })
            .catch(err => {
                console.log(err);
            })

    }

    return(
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <div className={classes.paper}>
                <Typography component="h1" variant="h5">
                    Login
                </Typography>
                <form className={classes.form} noValidate>
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        id="username"
                        label="Username"
                        name="username"
                        autoComplete="username"
                        autoFocus
                        value={loginData.username}
                        onChange={handleChange}
                    />
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        name="password"
                        label="Password"
                        type="password"
                        id="password"
                        autoComplete="current-password"
                        value={loginData.password}
                        onChange={handleChange}
                    />
                    <FormControlLabel
                        control={<Checkbox value="remember" color="primary" />}
                        label="Remember me"
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        color="primary"
                        className={classes.submit}
                        onClick={handleSubmit}
                    >
                        Sign In
                    </Button>
                </form>
            </div>
        </Container>
    );
}
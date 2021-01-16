import React from 'react';
import { Link } from 'react-router-dom';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
//import SearchIcon from '@material-ui/icons/Search';

const useStyles = makeStyles((theme) => ({
    root: {
        background: '#2d2e40',
    },
    title: {
        flexGrow: 1,
        [theme.breakpoints.up('sm')]: {
        display: 'block',
        },
    },
}));

export default function Navigation() {
    const classes = useStyles();

    return(
        <AppBar className={classes.root} position="fixed">
            <Toolbar>
                <Link to="/">
                    <Typography className={classes.title} variant="h5" noWrap>
                        kernel-panic.pl
                    </Typography>
                </Link>
            </Toolbar>
        </AppBar>
    );
}
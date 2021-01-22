import React, {useState, useEffect} from 'react';
import { Link } from 'react-router-dom';
import Container from '@material-ui/core/Container';
import { makeStyles } from '@material-ui/core/styles';
import TruncateString from 'react-truncate-string'
import Typography from '@material-ui/core/Typography';
import Grid from '@material-ui/core/Grid';
import axios from 'axios';

const useStyles = makeStyles((theme) => ({
    postBox: {
        paddingTop: 15,
        paddingBottom: 15,
    },
    postImage: {
        margin: 'auto',
        display: 'block'
    },
    postTitle: {
        color: 'inherit',
        textDecoration: 'none',
    },
}));

export default function Home() {
    const classes = useStyles();
    const [posts, setPosts] = useState([]);
    
    useEffect(() => {
        axios.get('http://localhost:8888/posts')
            .then(response => {
                console.log(response);
                setPosts(response.data);
            })
            .catch(err =>{
                console.log(err);
            })
    }, [])

    return (
        <Container style={{marginTop: 100, maxWidth: 1000}}>
            {posts.map(post =>(
                <Link to={{pathname: `post/${post.slug}`}} style={{color: 'inherit', textDecoration: 'none'}}>
                    <Grid container spacing={5} direction="row" className={classes.postBox}>
                        <Grid item>
                            <img src="https://via.placeholder.com/350x200" />
                        </Grid>
                        <Grid item xs={12}>
                            <Grid container direction="row">
                                <h1>{post.title}</h1>
                            </Grid>
                        </Grid>
                    </Grid>
                </Link>
            ))}
        </Container>
    );
}

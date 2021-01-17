import React, {useState, useEffect} from 'react';
import { Link } from 'react-router-dom';
import Container from '@material-ui/core/Container';
import { makeStyles } from '@material-ui/core/styles';
import TruncateString from 'react-truncate-string'
import axios from 'axios';

const useStyles = makeStyles((theme) => ({
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
        <Container style={{marginTop: 100}}>
            {posts.map(post =>(
                <div key={post.id}>
                    <Link to={{
                        pathname: `post/${post.slug}`
                    }} style={{color: 'inherit', textDecoration: 'none'}}>
                        <h1>{post.title}</h1>
                    </Link>
                    <h3>{post.subtitle}</h3>
                    <TruncateString text={post.content} truncateAt={100} />
                </div>
            ))}
        </Container>
    );
}
import React, {useState, useEffect} from 'react';
import { Link } from 'react-router-dom';
import Container from '@material-ui/core/Container';
import { makeStyles } from '@material-ui/core/styles';
import axios from 'axios';

const useStyles = makeStyles((theme) => ({
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
                    }}>
                        <h1>{post.title}</h1>
                    </Link>
                    <h3>{post.subtitle}</h3>
                    <p>{post.content}</p>
                </div>
            ))}
        </Container>
    );
}
import React, {useState, useEffect} from 'react';
import Container from '@material-ui/core/Container';
import { makeStyles } from '@material-ui/core/styles';
import axios from 'axios';

const useStyes = makeStyles((theme) => ({
    postText: {
        fontSize: 20,
    },
}));


export default function Post(props){

    const classes = useStyes();

    const [data, setData] = useState([]);
    const [comments, setComments] = useState([]);

    useEffect(() => {
        console.log(props.match.params.slug);
        //post
        axios.get(`http://localhost:8888/posts/${props.match.params.slug}`)
            .then(response => {
                setData(response.data.post);
            })
            .catch(err => {
                console.log(err);
            })
    }, [])

    return (
        <Container>
            <Container style={{marginTop: 100, maxWidth: 800}}>
                <div>
                    <h1>{data.title}</h1>
                    <h3>{data.subtitle}</h3>
                    <p className={classes.postText}>{data.content}</p>
                </div>
            </Container>
        </Container>
    );
}
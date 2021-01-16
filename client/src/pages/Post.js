import React, {useState, useEffect} from 'react';
import Container from '@material-ui/core/Container';
import axios from 'axios';


export default function Post(props){
    const [data, setData] = useState([]);

    useEffect(() => {
        console.log(props.match.params.slug);
        //post
        axios.get(`http://localhost:8888/posts/${props.match.params.slug}`)
            .then(response => {
                console.log(response.data)
                setData(response.data.post);
            })
            .catch(err => {
                console.log(err);
            })
    }, [])

    return (
        <Container style={{marginTop: 100}}>
            <div>
                <h1>{data.title}</h1>
                <h3>{data.subtitle}</h3>
                <p>{data.content}</p>
            </div>
        </Container>
    );
}
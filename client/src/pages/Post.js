import React, {useState, useEffect} from 'react';
import Container from '@material-ui/core/Container';
import TextareaAutosize from '@material-ui/core/TextareaAutosize';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import { makeStyles } from '@material-ui/core/styles';
import axios from 'axios';

const useStyes = makeStyles((theme) => ({
    postText: {
        fontSize: 20,
    },
    addCommentForm: {
        width: '100%', // Fix IE 11 issue.
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
        backgroundColor: '#2d2e40'
    },
    paper: {
        padding: theme.spacing(2),
        margin: 'auto',
        maxWidth: 500,
    },
}));


export default function Post(props){

    const classes = useStyes();

    const [data, setData] = useState([]);
    const [comments, setComments] = useState([]);
    const [commentData, setCommentData] = useState({
        post_id: null,
        content: "",
        is_admin: false,
        op_email: "",
        op_name: "",
    })

    useEffect(() => {
        //post
        axios.get(`http://localhost:8888/posts/${props.match.params.slug}`)
            .then(response => {
                setData(response.data.post);
            })
            .catch(err => {
                console.log(err);
            })

        
    }, [])
    
    const loadComments = () => {
        axios.get(`http://localhost:8888/post/${data.id}/comments`)
            .then(response => {
                console.log(response);
                setComments(response.data.comments);
            })
            .catch(err => {
                console.log(err);
            })
    }

    const handleChange = (e) => {
        const { id, value } = e.target
        setCommentData(prevCommentData => ({
            ...prevCommentData,
            [id]: value
        }));
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        const payload = {
            "post_id": data.id,
            "content": commentData.content,
            "is_admin": false,
            "op_email": commentData.op_email,
            "op_name": commentData.op_name
        }

        axios.post(`http://localhost:8888/post/${data.id}/comments`, payload)
            .then(response => {
                console.log(response);
            })
            .catch(err => {
                console.log(err);
            })
        loadComments();
    }

    return (
        <Container>
            <Container style={{marginTop: 100, maxWidth: 800}}>
                <div>
                    <h1>{data.title}</h1>
                    <h3>{data.subtitle}</h3>
                    <p className={classes.postText}>{data.content}</p>
                </div>
            </Container>
            <Container style={{maxWidth:800, paddingTop:50, paddingBottom:50}}>
                {comments.map(comment =>(
                    <div key={comment.id}>
                        <h6>{comment.content}</h6>
                    </div>
                ))}
            </Container>
            <Container style={{maxWidth:800}}>
                <form className={classes.addCommentForm} noValidate>
                    <TextField 
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        id="op_name"
                        label="Name"
                        autoFocus
                        value={commentData.op_name}
                        onChange={handleChange}
                    />
                    <TextField 
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        id="op_email"
                        label="Email"
                        autoFocus
                        value={commentData.op_email}
                        onChange={handleChange}
                    />
                    <TextareaAutosize 
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        id="content"
                        label="Content"
                        value={commentData.content}
                        onChange={handleChange}
                    />
                     <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        color="primary"
                        className={classes.submit}
                        onClick={handleSubmit}
                    >
                        Add Comment
                    </Button>
                </form>
            </Container>
        </Container>
    );
}
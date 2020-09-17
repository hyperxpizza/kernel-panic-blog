import React, { useEffect, useState } from 'react';
import PostList from './Components/PostList';
import PostListLoading from './Components/PostListLoading';


function App() {
  const ListLoading = PostListLoading(PostList);
  const [appState, setAppState] = useState({
    loading: false,
    posts: null,
  });

  useEffect(() => {
    setAppState({ loading: true });
    const apiUrl = `http://localhost:8888/posts`;
    fetch(apiUrl)
      .then((res) => res.json())
      .then((posts) => {
        setAppState({ loading: false, posts: posts });
      });
  }, [setAppState]);
  return (
    <div className='App'>
      <div className='container'>
        <h1>My Posts</h1>
      </div>
      <div className='repo-container'>
        <PostListLoading isLoading={appState.loading} repos={appState.posts} />
      </div>
    </div>
  );
}
export default App;
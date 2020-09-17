import React from 'react';
const List = (props) => {
  const { posts } = props;
  if (!posts || posts.length === 0) return <p>No repos, sorry</p>;
  return (
    <ul>
      <h2 className='list-head'>Available Posts</h2>
      {posts.map((post) => {
        return (
          <li key={post.id} className='list'>
            <span className='repo-text'>{post.title} </span>
            <span className='repo-description'>{post.content}</span>
          </li>
        );
      })}
    </ul>
  );
};
export default List;
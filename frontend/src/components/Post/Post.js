import React from 'react';
import { Link } from 'react-router-dom';
import "./Post.scss";

const Post = ({ post }) => {
  let formattedDate = new Date(post?.created_at);
  let date = `${formattedDate.getDate().toString().padStart(2, '0')}.${(formattedDate.getMonth() + 1).toString().padStart(2, '0')}.${formattedDate.getFullYear()}`;
  return (
    
    <Link to = {`/blog/${post?.id}`} key = {post?.id}>

    <div className='post-item bg-white' key={post?.id}>
      <div className='post-item-img'>
        <img className='img-cover' src={post?.image || ""} alt={post?.title} />
      </div>
      <div className='post-item-info fs-14'>
        <h3 className='title'>{post?.title}</h3>
        <p className='description'>{post?.description}</p>
        <div className='post-item-footer'>
          <span>{date}</span>
        </div>
      </div>
    </div>
    </Link>
  );
};

export default Post;

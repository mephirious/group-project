import React from 'react';
import './BlogPostListItem.scss';

const BlogPostListItem = ({ item, onSelect }) => {
  return (
    <div className="blog-post-list-item" onClick={() => onSelect(item)}>
      <div className="title">{item.title}</div>
      <div className="created-at">{new Date(item.created_at).toLocaleString()}</div>
    </div>
  );
};

export default BlogPostListItem;

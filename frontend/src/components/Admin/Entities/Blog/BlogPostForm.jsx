import React, { useState, useEffect } from 'react';
import { useDispatch } from 'react-redux';
import OperationTabs from '../../Common/OperationTabs';
import ConfirmModal from '../../Common/ConfirmModal';
import FormField from '../../Common/FormFields/FormField';
import { createBlogPost, updateBlogPost, deleteBlogPost } from '../../../../store/adminSlice';
import './BlogPostForm.scss';
import ImageField from '../../Common/FormFields/ImageField';

const BlogPostForm = ({ selectedData, onActionSuccess }) => {
  const dispatch = useDispatch();
  const [operation, setOperation] = useState('create');
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [images, setImages] = useState([]);
  const [showConfirm, setShowConfirm] = useState(false);
  const [pendingAction, setPendingAction] = useState(null);

  useEffect(() => {
    if (selectedData) {
      setOperation('update');
      setTitle(selectedData.title);
      setContent(selectedData.content);
      setImages([selectedData.image] || []);
    } else {
      setOperation('create');
      resetFields();
    }
  }, [selectedData]);

  const resetFields = () => {
    setTitle('');
    setContent('');
    setImages([]);
  };

  const handleSubmit = e => {
    e.preventDefault();
    setShowConfirm(true);
    setPendingAction(() => () => {
      const image = images[0];
      const blogPostData = { title, content, image };
      if (operation === 'create') {
        dispatch(createBlogPost(blogPostData));
      } else if (operation === 'update') {
        dispatch(updateBlogPost({ id: selectedData.id, data: blogPostData }));
      } else if (operation === 'delete') {
        dispatch(deleteBlogPost(selectedData.id));
      }
      resetFields();
    });
  };

  return (
    <div className="blog-post-form">
      <div className="form-header">
        <h2>
          {operation === 'create'
            ? 'Create Blog Post'
            : operation === 'update'
            ? 'Update Blog Post'
            : 'Delete Blog Post'}
        </h2>
        <OperationTabs currentOperation={operation} onOperationChange={setOperation} />
      </div>
      <form onSubmit={handleSubmit}>
        {operation !== 'delete' && (
          <>
            <FormField label="Title" name="title" value={title} onChange={e => setTitle(e.target.value)} required />
            <div className="form-field">
              <label>Content</label>
              <textarea name="content" value={content} onChange={e => setContent(e.target.value)} required />
            </div>
            <ImageField label="Image URL" images={images} limitImages={1} setImages={setImages}></ImageField>
          </>
        )}
        {operation === 'delete' && (
          <FormField label="ID" name="id" value={selectedData?.id || ''} onChange={() => {}} readOnly />
        )}
        <button type="submit" className="submit-btn">
          {operation === 'create'
            ? 'Create Blog Post'
            : operation === 'update'
            ? 'Update Blog Post'
            : 'Delete Blog Post'}
        </button>
      </form>
      {showConfirm && (
        <ConfirmModal
          message={`Are you sure you want to ${operation} this blog post?`}
          onConfirm={() => {
            pendingAction();
            setShowConfirm(false);
            if (onActionSuccess) onActionSuccess();
          }}
          onCancel={() => setShowConfirm(false)}
        />
      )}
    </div>
  );
};

export default BlogPostForm;

import React, { useEffect } from 'react';
import "./BlogSinglePage.scss";
import { useParams } from "react-router-dom";
import { useSelector, useDispatch } from "react-redux";
import { fetchAsyncPostSingle, getPostSingle, getSinglePostStatus } from '../../store/blogSlice';
import { STATUS } from '../../utils/status';
import Loader from "../../components/Loader/Loader";
import { formatDate } from "../../utils/helpers";

const SingleBlogPage = () => {
  const { id } = useParams();
  const dispatch = useDispatch();
  const blog = useSelector(getPostSingle);
  const blogSingleStatus = useSelector(getSinglePostStatus);

  useEffect(() => {
    dispatch(fetchAsyncPostSingle(id));
  }, [id, dispatch]);

  if (blogSingleStatus === STATUS.LOADING) {
    return <Loader />;
  }

  return (
    <main className='py-5 bg-whitesmoke'>
      <div className='blog-single'>
        <div className='container'>
          <div className='blog-single-content bg-white grid'>
            <div className='blog-single-header'>
              <h1 className='title fs-28 fw-6'>{blog?.title}</h1>
              <div className='meta-info flex align-center gap-2 my-2'>
                <span className='created-at fs-14 text-gray'>
                  Published: {formatDate(blog?.created_at)}
                </span>
                {blog?.updated_at !== blog?.created_at && (
                  <span className='updated-at fs-14 text-gray'>
                    Updated: {formatDate(blog?.updated_at)}
                  </span>
                )}
              </div>
            </div>

            {blog?.image && (
              <div className='blog-image my-3'>
                <img 
                  src={blog.image} 
                  alt={blog.title} 
                  className='img-cover' 
                />
                {blog.image_source && (
                  <div className='image-caption fs-12 text-gray mt-1'>
                    Image source: {blog.image_source}
                  </div>
                )}
              </div>
            )}

            <div className='blog-content font-manrope'>
              {blog?.content?.split('\n').map((paragraph, index) => (
                <p key={index} className='content-paragraph fs-16 lh-1-6 my-2'>
                  {paragraph}
                </p>
              ))}
            </div>
          </div>
        </div>
      </div>
    </main>
  );
};

export default SingleBlogPage;
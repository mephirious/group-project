import React, { useEffect, useState } from 'react';
import { BASE_URL } from '../../../utils/apiURL';
import Pagination from './Pagination';
import './DataList.scss';
import BlogPostListItem from '../Entities/Blog/BlogPostListItem';
import ProductListItem from '../Entities/Product/ProductListItem';

const entityEndpointMap = {
  brands: 'products/brands',
  types: 'products/types',
  categories: 'products/categories',
  'products': 'products/products',
  'blog-posts': 'blogs/blog-posts'
};

const DataList = ({ entity, onSelectData, refresh }) => {
  const [data, setData] = useState([]);
  const [status, setStatus] = useState('idle');
  const [limit] = useState(10);
  const [inputSearch, setInputSearch] = useState('');
  const [searchTerm, setSearchTerm] = useState('');
  const [page, setPage] = useState(0);
  const [totalPages, setTotalPages] = useState(1000);
  const [sortField, setSortField] = useState('created_at');
  const [sortOrder, setSortOrder] = useState('desc');

  const fetchData = () => {
    if (!entity) return;
    setStatus('loading');
    let endpoint = entityEndpointMap[entity];
    if (!endpoint) return;
    let url = `${BASE_URL}${endpoint}`;
    if (entity === 'products' || entity === 'blog-posts') {
      const skip = page * limit;
      const params = new URLSearchParams({
        limit: limit,
        sortField: sortField,
        sortOrder: sortOrder,
        skip: skip,
      });
      if (searchTerm) params.append('search', searchTerm);
      url = `${url}?${params.toString()}`;
    }
    fetch(url, { credentials: 'include' })
      .then((res) => res.json())
      .then((json) => {
        setData(json.data || json);
        if (json.total) setTotalPages(Math.ceil(json.total / limit));
        setStatus('succeeded');
      })
      .catch(() => setStatus('failed'));
  };

  useEffect(() => {
    fetchData();
  }, [entity, refresh, page, searchTerm, sortField, sortOrder]);

  const handleFetchClick = () => {
    setPage(0);
    setSearchTerm(inputSearch);
  };

  if (status === 'loading') return <div className="data-list">Loading data...</div>;
  if (status === 'failed') return <div className="data-list">Error fetching data.</div>;

  return (
    <div className="data-list">
      <h3>{entity.charAt(0).toUpperCase() + entity.slice(1)} List</h3>
      {entity === 'products' || entity === 'blog-posts' ? (
        <div className="controls">
          <input
            type="text"
            placeholder="Search..."
            value={inputSearch}
            onChange={(e) => setInputSearch(e.target.value)}
          />
          <button onClick={handleFetchClick}>Fetch</button>
          <select value={sortField} onChange={(e) => setSortField(e.target.value)}>
            <option value="created_at">Created At</option>
            <option value="price">Price</option>
            <option value="model_name">Model Name</option>
          </select>
          <select value={sortOrder} onChange={(e) => setSortOrder(e.target.value)}>
            <option value="asc">Ascending</option>
            <option value="desc">Descending</option>
          </select>
        </div>
      ) : null}
      <ul>
        {data.map((item) => (
          <li key={item.id}>
            {entity === 'blog-posts' ? (
              <BlogPostListItem item={item} onSelect={onSelectData} />
            ) : entity === 'products' ? (
              <ProductListItem item={item} onSelect={onSelectData} />
            ) : (
              <div onClick={() => onSelectData(item)}>
                {entity === 'brands' ? item.brand_name : entity === 'types' ? item.type_name : item.category_name}
              </div>
            )}
          </li>
        ))}
      </ul>
      {entity === 'products' && (
        <Pagination currentPage={page} totalPages={totalPages} onPageChange={setPage} />
      )}
    </div>
  );
};

export default DataList;

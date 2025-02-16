import React, { useEffect, useState } from 'react';
import { BASE_URL } from '../../utils/apiURL';
import './AdminDataList.scss';

const AdminDataList = ({ entity, onSelectData, refresh }) => {
  const [data, setData] = useState([]);
  const [status, setStatus] = useState('idle');
  
  // For products pagination and search
  const [limit, setLimit] = useState(50);
  const [inputSearch, setInputSearch] = useState('');
  const [searchTerm, setSearchTerm] = useState('');
  const [page, setPage] = useState(0);

  const fetchData = () => {
    if (!entity) return;
    setStatus('loading');
    let endpoint = '';
    let url = '';
    if (entity === 'brands') {
      endpoint = 'products/brands';
      url = `${BASE_URL}${endpoint}`;
    } else if (entity === 'types') {
      endpoint = 'products/types';
      url = `${BASE_URL}${endpoint}`;
    } else if (entity === 'categories') {
      endpoint = 'products/categories';
      url = `${BASE_URL}${endpoint}`;
    } else if (entity === 'products') {
      endpoint = 'products/products';
      const skip = page * limit;
      const params = new URLSearchParams({
        limit: limit,
        sortField: 'created_at',
        sortOrder: 'desc',
        skip: skip,
      });
      if (searchTerm) {
        params.append('search', searchTerm);
      }
      url = `${BASE_URL}${endpoint}?${params.toString()}`;
    }
    fetch(url, { credentials: 'include' })
      .then((res) => res.json())
      .then((json) => {
        setData(json);
        setStatus('succeeded');
      })
      .catch(() => setStatus('failed'));
  };

  useEffect(() => {
    fetchData();
  }, [entity, refresh, page, searchTerm]);

  const handleFetchClick = () => {
    setPage(0);
    setSearchTerm(inputSearch);
  };

  const renderPagination = () => {
    if (entity !== 'products') return null;
    // For simplicity, render buttons for pages 1-10.
    const pages = [];
    for (let i = 0; i < 10; i++) {
      pages.push(
        <button
          key={i}
          className={i === page ? 'active' : ''}
          onClick={() => setPage(i)}
        >
          {i + 1}
        </button>
      );
    }
    return <div className="pagination">{pages}</div>;
  };

  if (status === 'loading') return <div className="admin-data-list">Loading data...</div>;
  if (status === 'failed') return <div className="admin-data-list">Error fetching data.</div>;

  return (
    <div className="admin-data-list">
      <h3>{entity.charAt(0).toUpperCase() + entity.slice(1)} List</h3>
      {entity === 'products' && (
        <div className="search-controls">
          <input
            type="text"
            placeholder="Search products..."
            value={inputSearch}
            onChange={(e) => setInputSearch(e.target.value)}
          />
          <button onClick={handleFetchClick}>Fetch</button>
        </div>
      )}
      <ul>
        {data.map((item) => (
          <li key={item.id} onClick={() => onSelectData(item)}>
            {entity === 'brands'
              ? item.brand_name
              : entity === 'types'
              ? item.type_name
              : entity === 'categories'
              ? item.category_name
              : entity === 'products'
              ? item.model_name
              : ''}
          </li>
        ))}
      </ul>
      {renderPagination()}
    </div>
  );
};

export default AdminDataList;

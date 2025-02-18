import React, { useState } from 'react';
import './Pagination.scss';

const Pagination = ({ currentPage, totalPages, onPageChange }) => {
  const [inputPage, setInputPage] = useState(currentPage + 1);
  const maxVisibleButtons = 5;

  const handlePageChange = (page) => {
    if (page >= 0 && page < totalPages) {
      onPageChange(page);
      setInputPage(page + 1);
    }
  };

  const handleInputChange = (e) => {
    setInputPage(e.target.value);
  };

  const handleInputSubmit = () => {
    const page = parseInt(inputPage, 10) - 1;
    if (!isNaN(page)) {
      handlePageChange(page);
    }
  };

  const renderPaginationButtons = () => {
    let pages = [];
    const half = Math.floor(maxVisibleButtons / 2);

    let start = Math.max(0, currentPage - half);
    let end = Math.min(totalPages, start + maxVisibleButtons);

    if (end - start < maxVisibleButtons) {
      start = Math.max(0, end - maxVisibleButtons);
    }

    if (start > 0) pages.push(<button key="start" disabled>...</button>);

    for (let i = start; i < end; i++) {
      pages.push(
        <button
          key={i}
          className={currentPage === i ? 'active' : ''}
          onClick={() => handlePageChange(i)}
        >
          {i + 1}
        </button>
      );
    }

    if (end < totalPages) pages.push(<button key="end" disabled>...</button>);

    return pages;
  };

  return (
    <div className="pagination">
      <button disabled={currentPage === 0} onClick={() => handlePageChange(currentPage - 1)}>
        Prev
      </button>
      {renderPaginationButtons()}
      <button disabled={currentPage === totalPages - 1} onClick={() => handlePageChange(currentPage + 1)}>
        Next
      </button>
      <div className="page-input">
        <input
          type="number"
          value={inputPage}
          onChange={handleInputChange}
          onKeyDown={(e) => e.key === 'Enter' && handleInputSubmit()}
        />
        <button onClick={handleInputSubmit}>Go</button>
      </div>
    </div>
  );
};

export default Pagination;

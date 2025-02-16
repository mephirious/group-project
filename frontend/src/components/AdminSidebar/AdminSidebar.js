import React from 'react';
import './AdminSidebar.scss';

const AdminSidebar = ({ onSelectEntity, selectedEntity }) => {
  return (
    <aside className="admin-sidebar">
      <h2>Admin Edit</h2>
      <ul>
        <li className={selectedEntity === 'brands' ? 'active' : ''} onClick={() => onSelectEntity('brands')}>
          Brands
        </li>
        <li className={selectedEntity === 'types' ? 'active' : ''} onClick={() => onSelectEntity('types')}>
          Types
        </li>
        <li className={selectedEntity === 'categories' ? 'active' : ''} onClick={() => onSelectEntity('categories')}>
          Categories
        </li>
      </ul>
    </aside>
  );
};

export default AdminSidebar;

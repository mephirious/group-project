import React, { useState, useRef, useEffect } from 'react';
import AdminSidebar from '../../components/Admin/Sidebar/AdminSidebar';
import ProductForm from '../../components/Admin/Entities/Product/ProductForm';
import CrudForm from '../../components/Admin/Entities/Crud/CrudForm';
import BlogPostForm from '../../components/Admin/Entities/Blog/BlogPostForm';
import DataList from '../../components/Admin/Common/DataList';
import './AdminEditPage.scss';

const AdminEditPage = () => {
  const [selectedEntity, setSelectedEntity] = useState(null);
  const [selectedData, setSelectedData] = useState(null);
  const [refresh, setRefresh] = useState(0);
  const dataListRef = useRef(null);

  const handleRefresh = () => setRefresh(prev => prev + 1);

  useEffect(() => {
    if (dataListRef.current) {
      dataListRef.current.scrollIntoView({ behavior: 'smooth', block: 'start' });
    }
  }, [refresh]);

  return (
    <div className="admin-edit-page">
      <AdminSidebar
        onSelectEntity={(entity) => {
          setSelectedEntity(entity);
          setSelectedData(null);
        }}
        selectedEntity={selectedEntity}
      />
      <div className="admin-edit-content">
        {selectedEntity ? (
          <>
            {selectedEntity === 'products' ? (
              <ProductForm entity={selectedEntity} selectedData={selectedData} onActionSuccess={handleRefresh} />
            ) : selectedEntity === 'blog-posts' ? (
              <BlogPostForm entity={selectedEntity} selectedData={selectedData} onActionSuccess={handleRefresh} />
            ) : (
              <CrudForm entity={selectedEntity} selectedData={selectedData} onActionSuccess={handleRefresh} />
            )}
            <div ref={dataListRef}>
              <DataList entity={selectedEntity} onSelectData={setSelectedData} refresh={refresh} />
            </div>
          </>
        ) : (
          <div className="select-prompt">Please select an entity to edit.</div>
        )}
      </div>
    </div>
  );
};

export default AdminEditPage;

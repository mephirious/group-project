import React, { useState } from 'react';
import './AdminEditPage.scss';
import AdminSidebar from '../../components/AdminSidebar/AdminSidebar';
import AdminCrudForm from '../../components/AdminCrudForm/AdminCrudForm';
import AdminDataList from '../../components/AdminDataList/AdminDataList';
import AdminProductForm from '../../components/AdminProductForm/AdminProductForm';

const AdminEditPage = () => {
  const [selectedEntity, setSelectedEntity] = useState(null);
  const [selectedData, setSelectedData] = useState(null);
  const [refresh, setRefresh] = useState(0);

  const handleRefresh = () => setRefresh((prev) => prev + 1);

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
              <AdminProductForm
                selectedProduct={selectedData}
                onActionSuccess={handleRefresh}
              />
            ) : (
              <AdminCrudForm
                entity={selectedEntity}
                selectedData={selectedData}
                onActionSuccess={handleRefresh}
              />
            )}
            <AdminDataList
              entity={selectedEntity}
              onSelectData={setSelectedData}
              refresh={refresh}
            />
          </>
        ) : (
          <div className="select-prompt">Please select an entity to edit.</div>
        )}
      </div>
    </div>
  );
};

export default AdminEditPage;

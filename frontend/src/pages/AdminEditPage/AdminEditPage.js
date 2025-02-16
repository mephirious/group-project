import React, { useState } from 'react';
import './AdminEditPage.scss';
import AdminCrudForm from '../../components/AdminCrudForm/AdminCrudForm';
import AdminSidebar from '../../components/AdminSidebar/AdminSidebar';

const AdminEditPage = () => {
  const [selectedEntity, setSelectedEntity] = useState(null);
  return (
    <div className="admin-edit-page">
      <AdminSidebar onSelectEntity={setSelectedEntity} selectedEntity={selectedEntity} />
      <div className="admin-edit-content">
        {selectedEntity ? (
          <AdminCrudForm entity={selectedEntity} />
        ) : (
          <div className="select-prompt">Please select an entity to edit.</div>
        )}
      </div>
    </div>
  );
};

export default AdminEditPage;

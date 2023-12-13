import React from 'react';
import './CustomerEditForm.css';
import { formatDate } from '../func/timeconverter';

const CustomerEditForm = ({
  handleImageChange,
  editedCustomerData,
  setEditedCustomerData,
  handleSaveEdit,
}) => {
  return (
    <div className="customer-edit-container">
      {editedCustomerData.customer_picture_url && (
        <img src={editedCustomerData.customer_picture_url} alt="" className="current-preview-image" />
      )}
      <input
        type="file"
        accept="image/*"
        onChange={handleImageChange}
        className="file-input"
      />
      <label htmlFor="birthDateInput">Doğum Tarihi</label>
      <input
        type="text"
        value={editedCustomerData.birth_date}
        onChange={(e) =>
          setEditedCustomerData({
            ...editedCustomerData,
            birth_date: e.target.value,
          })
        }
        placeholder="Doğum Tarihi (gg/aa/yy)"
        className="text-input"
      />
      <button onClick={handleSaveEdit} className="save-button">Kaydet</button>
    </div>
  );
};

export default CustomerEditForm;

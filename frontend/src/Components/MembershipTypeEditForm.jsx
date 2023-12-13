import React from 'react';
import './MembershipTypeEditForm.css';

const MembershipTypeEditForm = ({
  editedMembershipTypeData,
  setEditedMembershipTypeData,
  handleSaveEdit,
}) => {
  return (
    <div className="membership-type-edit-container">
      <input
        type="text"
        value={editedMembershipTypeData.name}
        onChange={(e) =>
          setEditedMembershipTypeData({
            ...editedMembershipTypeData,
            name: e.target.value,
          })
        }
        placeholder="Üyelik Türü Adı"
        className="text-input"
      />
      <textarea
        value={editedMembershipTypeData.description}
        onChange={(e) =>
          setEditedMembershipTypeData({
            ...editedMembershipTypeData,
            description: e.target.value,
          })
        }
        placeholder="Açıklama"
        className="text-area-input"
      ></textarea>
      <input
        type="number"
        value={editedMembershipTypeData.membership_month}
        onChange={(e) =>
          setEditedMembershipTypeData({
            ...editedMembershipTypeData,
            membership_month: parseInt(e.target.value),
          })
        }
        placeholder="Üyelik Süresi (Ay)"
        className="text-input"
      />
      <input
        type="number"
        value={editedMembershipTypeData.price}
        onChange={(e) =>
          setEditedMembershipTypeData({
            ...editedMembershipTypeData,
            price: parseFloat(e.target.value),
          })
        }
        placeholder="Fiyat"
        className="text-input"
      />
      <button onClick={handleSaveEdit} className="save-button">
        Kaydet
      </button>
    </div>
  );
};

export default MembershipTypeEditForm;

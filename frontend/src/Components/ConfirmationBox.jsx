import React from 'react';

const ConfirmationBox = ({ message, onConfirm, onCancel }) => {
  return (
    <div className="confirmation-box">
      <p>{message}</p>
      <div className="confirmation-buttons">
        <button onClick={onCancel}>İptal</button>
        <button onClick={onConfirm}>Eminim</button>
      </div>
    </div>
  );
};

export default ConfirmationBox;
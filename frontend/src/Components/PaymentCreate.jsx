import React, { useState } from 'react';
import api from '../api/api';
import './PaymentCreate.css'; // CSS dosyasını burada import ediyoruz

const PaymentCreate = () => {
    const [type, setType] = useState('');
    const [description, setDescription] = useState('');
    const [amount, setAmount] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();

        const newPayment = {
            type: type,
            description: description,
            amount: parseFloat(amount),
        };

        try {
            const response = await api.post('/payment', newPayment);
            console.log('Yeni ödeme oluşturuldu:', response.data);
        } catch (error) {
            console.error('Ödeme oluşturulamadı:', error);
        }
    };

    return (
        <div className="payment-create">
            <h2>Yeni Ödeme Oluştur</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label>Tür:</label>
                    <select className="input-field" value={type} onChange={(e) => setType(e.target.value)}>
                        <option value="">Ödeme Tipini  Seçiniz</option>
                        <option value="in">Gelir</option>
                        <option value="out">Gider</option>
                    </select>
                </div>
                <div className="form-group">
    <label>Açıklama:</label>
    <textarea
        className="input-field input-field-textarea"
        value={description}
        placeholder="Açıklama girin"
        onChange={(e) => setDescription(e.target.value)}
    />
</div>
                <div className="form-group">
                    <label>Tutar:</label>
                    <input
                        className="input-field"
                        type="number"
                        value={amount}
                        placeholder="Tutar girin"
                        onChange={(e) => setAmount(e.target.value)}
                    />
                </div>
                <button className="submit-button" type="submit">
                    Ödeme Yap
                </button>
            </form>
        </div>
    );
};

export default PaymentCreate;

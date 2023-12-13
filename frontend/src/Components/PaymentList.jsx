import React, { useState, useEffect } from 'react';
import api from '../api/api';
import './PaymentList.css';

const PaymentList = () => {
    const [payments, setPayments] = useState([]);

    useEffect(() => {
        const fetchPayments = async () => {
            try {
                const response = await api.get('/payment');
                const sortedPayments = response.data.sort((a, b) => new Date(b.date_time) - new Date(a.date_time));
                setPayments(sortedPayments);
            } catch (error) {
                console.error('Ödemeler alınamadı:', error);
            }
        };

        fetchPayments();
    }, []);

    const formatType = (type) => {
        return type === 'in' ? 'Gelir' : type === 'out' ? 'Gider' : '';
    };

    const getRowStyle = (type) => {
        return type === 'in' ? { backgroundColor: 'lightgreen' } : type === 'out' ? { backgroundColor: 'lightcoral' } : {};
    };

    return (
        <div className="payment-container">
            <h2>Ödemeler</h2>
            <div className="table-container">
                <table className="payment-table">
                    <thead>
                        <tr>
                            <th className="payment-cell">Tür</th>
                            <th className="payment-cell">Açıklama</th>
                            <th className="payment-cell">Tutar</th>
                            <th className="payment-cell">Tarih ve Saat</th>
                        </tr>
                    </thead>
                    <tbody>
                        {payments.map((payment) => (
                            <tr key={payment.id} className="payment-row" style={getRowStyle(payment.type)}>
                                <td className="payment-cell">{formatType(payment.type)}</td>
                                <td className="payment-cell">{payment.description}</td>
                                <td className="payment-cell">{payment.amount} TL</td>
                                <td className="payment-cell">{new Date(payment.date_time).toLocaleString()}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default PaymentList;

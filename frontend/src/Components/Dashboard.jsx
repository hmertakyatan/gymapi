import React, { useState, useEffect, useRef } from 'react';
import api from '../api/api';
import Chart from 'chart.js/auto';
import './Dashboard.css'

const Dashboard = () => {
  const chartRef = useRef(null);
  const [profitData, setProfitData] = useState({});
  

  useEffect(() => {
    const fetchProfitData = async () => {
      try {
        const response = await api.get('/payment/profit');
        const { Profit, total_income, total_expense } = response.data;

        if (chartRef && chartRef.current) {
          const myChart = new Chart(chartRef.current, {
            type: 'bar',
            data: {
              labels: ['Kâr', 'Gelir', 'Gider'],
              datasets: [
                {
                  label: 'Kâr-Zarar Analizi',
                  data: [Profit, total_income, total_expense],
                  backgroundColor: [
                    '#358d1a',
                    '#79ea86',
                    '#e75757',
                  ],
                },
              ],
            },
            options: {
              maintainAspectRatio: false,
              scales: {
                y: {
                  beginAtZero: true,
                },
              },
            },
          });
        }

        setProfitData({ Profit, total_income, total_expense });
      } catch (error) {
        console.error('Veri alınamadı:', error);
      }
    };

    fetchProfitData();
  }, []);

  const numericProfit = parseFloat(profitData.Profit);
  const numericTotalIncome = parseFloat(profitData.total_income);
  const numericTotalExpense = parseFloat(profitData.total_expense);

  return (
    <div className="dashboard">
      <h2>Dashboard</h2>
      <div className="chart-container">
        <canvas ref={chartRef}></canvas>
      </div>
      <div className="data-values">
      <p><strong>Kâr:</strong> <span style={{ color: numericProfit > 0 ? '#79ea86' : 'red', fontWeight: 'bold' }}>{numericProfit > 0 ? '+' : '-'}{Math.abs(numericProfit).toFixed(2)} TL</span></p>
      <p><strong>Toplam Gelir:</strong> <span style={{ color: '#79ea86', fontWeight: 'bold' }}>+{Math.abs(numericTotalIncome).toFixed(2)} TL</span></p>
      <p><strong>Toplam Gider:</strong> <span style={{ color: '#e75757', fontWeight: 'bold' }}>-{Math.abs(numericTotalExpense).toFixed(2)} TL</span></p>
      </div>
    </div>
  );
};

export default Dashboard;

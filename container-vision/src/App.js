import React, { useEffect, useState } from "react";
import { Table } from "antd";

const App = () => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchPingData = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/containers/", {
        method: "GET",
        credentials: "include",  
        headers: {
        "Content-Type": "application/json"
        }
      })
      if (!response.ok) {
        throw new Error("Error featching data.");
      }
      const result = await response.json();
      setData(result);
      setError(null);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPingData();
    const interval = setInterval(fetchPingData, 5000);
    return () => clearInterval(interval);
  }, []);

  const columns = [
    { title: "IP Адрес", dataIndex: "ip_address", key: "ip_address" },
    { title: "Время пинга", dataIndex: "ping_time", key: "ping_time" },
    { title: "Дата последнего успеха", dataIndex: "last_checked", key: "last_checked" },
  ];

  if (loading) return <p>Загрузка данных...</p>;
  if (error) return <p>Ошибка: {error}</p>;

  return <Table dataSource={data} columns={columns} rowKey="ip" />;
};

export default App;

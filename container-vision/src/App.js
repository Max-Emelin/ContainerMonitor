import React, { useEffect, useState } from "react";
import ContainerTable from './components/ContainerTable'

const App = () => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchPingData = async () => {
    try {
      const SERVER_PORT = process.env.REACT_APP_SERVER_PORT;  
      const API_URL = `http://localhost:${SERVER_PORT}/api/containers/`;
      const response = await fetch(API_URL, {
        method: "GET",
        credentials: "include",  
        headers: {
          "Content-Type": "application/json"
        }
      });

      if (!response.ok) {
        throw new Error("Error fetching data.");
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

  
  return (
    <div>
      <h1 style={{ textAlign: "center" }}>Список контейнеров</h1>
      <ContainerTable data={data} loading={loading} error={error} />
    </div>
  );
};

export default App;

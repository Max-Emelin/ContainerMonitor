import React from 'react';
import { Table, ConfigProvider } from 'antd';
import dayjs from "dayjs";

const ContainerTable = ({ data, loading, error }) => {
  const formatDate = (date) => dayjs(date).format("YYYY-MM-DD HH:mm:ss");

  const columns = [
    {
      title: "IP Адрес",
      dataIndex: "ip_address",
      key: "ip_address",
      align: "center",
    },
    {
      title: "Время пинга",
      dataIndex: "ping_time",
      key: "ping_time",
      align: "center",
      render: (text) => formatDate(text),
    },
    {
      title: "Дата последней успешной попытки",
      dataIndex: "last_checked",
      key: "last_checked",
      align: "center",
      render: (text) => formatDate(text),
    },
  ];

  if (loading) return <p>Загрузка данных...</p>;
  if (error) return <p>Ошибка: {error}</p>;

  return (
    <ConfigProvider
      theme={{
        components: {
          Table: {
            borderColor: "#716561",
            headerBg: "#e0e0e0",
            headerColor: "#444441",
            headerBorderRadius: 8,
            rowHoverBg: "#f0e6e0",
            cellPaddingBlock: 10,
            cellPaddingInline: 10,
            fontSize: 14,
            colorText: "#444441",
          },
        },
      }}
    >
      <div style={{ display: "flex", justifyContent: "center", padding: "20px" }}>
        <Table
          dataSource={data}
          columns={columns}
          rowKey="ip"
          pagination={{ pageSize: 10 }}
          style={{ width: "80%", borderRadius: 8, overflow: "hidden" }}
        />
      </div>
    </ConfigProvider>
  );
};

export default ContainerTable;

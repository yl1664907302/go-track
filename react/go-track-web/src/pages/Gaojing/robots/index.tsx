import { ProCard } from '@ant-design/pro-components';
import axios from 'axios';
import { useEffect, useState } from 'react';

const DataCard = () => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchData = async (params) => {
    try {
      const response = await axios.get('http://localhost:8080/alertmanger/robot', { params }); // 替换为你的后端 API URL
      console.log('Response data:', response.data); // 调试信息
      if (Array.isArray(response.data)) {
        setData(response.data);
      } else {
        console.error('Expected an array but got:', response.data);
        setData([]); // 设置为空数组，以避免 map 错误
      }
    } catch (error) {
      setError(error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    const params = { index: 'l1' }; // 自定义参数键值对
    fetchData(params);
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <div>
      {data.map((item) => (
        <ProCard
          key={item.robot_id}
          title={`robot信息 - ${item.name}`}
          bordered
          boxShadow
          // extra={
          //   <ProFormGroup>
          //     <ProFormSwitch
          //       name={`enable-${item.id}`}
          //       noStyle
          //       checkedChildren={'启用'}
          //       unCheckedChildren={'禁用'}
          //       defaultChecked={item.enable}
          //     />
          //   </ProFormGroup>
          // }
          tooltip={item.description}
          style={{ maxWidth: 300, marginBottom: 16 }}
        >
          <div>{item.robot_class}</div>
        </ProCard>
      ))}
    </div>
  );
};

export default DataCard;

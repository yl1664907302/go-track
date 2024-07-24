// src/pages/Gaojing/RedirectToExternal.js
import { useEffect } from 'react';

const RedirectToExternal = () => {
  useEffect(() => {
    window.location.href = 'https://www.example.com'; // 替换为目标外部网站URL
  }, []);

  return <div>Redirecting...</div>;
};

export default RedirectToExternal;

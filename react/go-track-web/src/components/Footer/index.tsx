import { GithubOutlined } from '@ant-design/icons';
import { DefaultFooter } from '@ant-design/pro-components';
import React from 'react';

const Footer: React.FC = () => {
  return (
    <DefaultFooter
      style={{
        background: 'none',
      }}
      links={[
        {
          key: 'Go Track',
          title: 'Go Track',
          href: 'https://github.com/yl1664907302/go-track',
          blankTarget: true,
        },
        {
          key: 'github',
          title: <GithubOutlined />,
          href: 'https://github.com/yl1664907302/go-track',
          blankTarget: true,
        },
      ]}
    />
  );
};

export default Footer;

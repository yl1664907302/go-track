import type { ProFormInstance } from '@ant-design/pro-components';
import {
  ProCard,
  ProForm,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
  StepsForm,
} from '@ant-design/pro-components';
import { message } from 'antd';
import axios from 'axios';
import { useRef } from 'react';

const waitTime = (time: number = 100) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(true);
    }, time);
  });
};

export default () => {
  const formRef = useRef<ProFormInstance>();

  const submitFormToServer = async (values) => {
    try {
      const response = await axios.post('http://localhost:8080/alertmanger/post/dingtalk', values);

      console.log('Response from server:', response.data);
      message.success('提交成功');
    } catch (error) {
      console.error('Error submitting form:', error);
      message.error('提交失败，请检查网络连接或稍后再试');
    }
  };

  return (
    <ProCard>
      <StepsForm<{
        name: string;
        receiver: string;
        markdown: string;
        robot_class: string;
        accesstoken: string;
      }>
        formRef={formRef}
        onFinish={async (values) => {
          console.log('All values:', values);
          await submitFormToServer(values);
          return true;
        }}
        formProps={{
          validateMessages: {
            required: '此项为必填项',
          },
        }}
      >
        <StepsForm.StepForm<{
          name: string;
          receiver: string;
        }>
          name="receiver"
          title="新增消息来源"
          stepProps={{
            description: '此处填入receiver信息',
          }}
          onFinish={async () => {
            console.log('Step 1 values:', formRef.current?.getFieldsValue());
            return true;
          }}
        >
          <ProFormText
            name="name"
            label="来源昵称"
            width="md"
            tooltip="用于显示面包屑名称"
            placeholder="请输入昵称"
            rules={[{ required: true }]}
          />
          <ProFormText
            name="receiver"
            label="revicer值"
            width="md"
            tooltip="alertmanger中转发路由的“revicer”的值"
            placeholder="请输入值"
            rules={[{ required: true }]}
          />
        </StepsForm.StepForm>

        <StepsForm.StepForm<{
          markdowntemplate: string;
        }>
          name="markdown2"
          title="配置消息样式"
          stepProps={{
            description: '此处填入MarkDown模板',
          }}
          onFinish={async () => {
            console.log('Step 2 values:', formRef.current?.getFieldsValue());
            return true;
          }}
        >
          <ProForm.Group>
            <ProFormTextArea
              name="markdown"
              width="lg"
              label="markdown模板"
              rules={[{ required: true }]}
            />
          </ProForm.Group>
        </StepsForm.StepForm>

        <StepsForm.StepForm<{
          robot_class: string;
          accesstoken: string;
        }>
          name="robot"
          title="新增robot"
          stepProps={{
            description: '此处填入新增robot信息',
          }}
          onFinish={async () => {
            console.log('Step 3 values:', formRef.current?.getFieldsValue());
            return true;
          }}
        >
          <ProFormSelect
            label="选择robot类型"
            name="robot_class"
            rules={[{ required: true }]}
            initialValue="dingtalk"
            options={[
              {
                value: 'dingtalk',
                label: '钉钉',
              },
              {
                value: 'wechat',
                label: '企业微信',
              },
            ]}
          />
          <ProFormTextArea
            name="accesstoken"
            width="lg"
            label="钉钉密钥"
            rules={[{ required: true }]}
          />
        </StepsForm.StepForm>
      </StepsForm>
    </ProCard>
  );
};

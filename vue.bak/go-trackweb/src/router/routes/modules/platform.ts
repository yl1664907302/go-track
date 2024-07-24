import {DEFAULT_LAYOUT} from '../base';
import {AppRouteRecordRaw} from '../types';

const DASHBOARD: AppRouteRecordRaw = {
    path: '/platform',
    name: 'platform',
    component: DEFAULT_LAYOUT,
    meta: {
        locale: 'menu.platform',
        requiresAuth: true,
        icon: 'icon-home',
        order: 0,
    },
    children: [
        {
            path: 'dingtalk',
            name: 'dingtalk',
            component: () => import('@/views/platform/dingtalk/index.vue'),
            meta: {
                locale: 'menu.platform.dingtalk',
                requiresAuth: true,
                roles: ['*'],
            },

        },
        {
            path: 'feishu',
            name: 'feishu',
            component: () => import('@/views/platform/feishu/index.vue'),
            meta: {
                locale: 'menu.platform.feishu',
                requiresAuth: true,
                roles: ['*'],
            },
        },
        {
            path: 'wechat_robot',
            name: 'wechat_robot',
            component: () => import('@/views/platform/wechat_robot/index.vue'),
            meta: {
                locale: 'menu.platform.wechat_robot',
                requiresAuth: true,
                roles: ['*'],
            },
        }
    ],
};

export default DASHBOARD;

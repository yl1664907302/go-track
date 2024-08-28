import type { AppRouteModule } from '/@/router/types'

import { LAYOUT } from '/@/router/constant'
import { t } from '/@/hooks/web/useI18n'

const test: AppRouteModule = {
  path: '/test',
  name: 'Test',
  component: LAYOUT,
  redirect: '/test/one',
  meta: {
    orderNo: 10,
    icon: 'ion:grid-outline',
    title: t('routes.test.test'),
  },
  children: [
    {
      path: 'one',
      name: 'One',
      component: () => import('/@/views/test/one/index.vue'),
      meta: {
        // affix: true,
        title: t('routes.test.one'),
      },
    },
  ],
}

export default test

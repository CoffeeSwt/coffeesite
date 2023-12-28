export const routes = [
  {
    path: '/',
    name: 'layout',
    component: () => import('@/pages/layout/index.vue'),
    children: [
      {
        path: '',
        name: 'index',
        component: () => import('@/pages/index/index.vue'),
      },
      {
        path: 'home',
        name: 'home',
        component: () => import('@/pages/home/index.vue'),
      },
      {
        path: 'test',
        name: 'test',
        component: () => import('@/pages/test/index.vue'),
      },
    ],
  },
]

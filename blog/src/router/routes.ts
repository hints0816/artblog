import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        component: () => import('pages/home/index.vue'),
      },
      {
        path: '/posts/:id',
        component: () => import('pages/actlist/index.vue'),
      },
      {
        path: '/edit',
        component: () => import('pages/edit/index.vue'),
        children: [
          {
            path: '/edit/:id',
            component: () => import('pages/edit/index.vue'),
          }
        ]
      },
      {
        path: '/profile',
        component: () => import('pages/profile/index.vue'),
      },
      {
        path: '/repository/:id',
        component: () => import('pages/repository/index.vue'),
      }
    ],
  },
  {
    path: '/login',
    component: () => import('pages/login/index.vue')
  },
  {
    path: '/password_reset',
    component: () => import('pages/resetpassword/index.vue')
  },
  {
    path: '/signup',
    component: () => import('pages/signup/index.vue'),
    children: [
      {
        path: '/signup/confirm',
        component: () => import('pages/signup/confirmed.vue'),
      }
    ]
  },
  {
    path: '/claim',
    component: () => import('pages/signup/claim.vue')
  },
  ,
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/Error404.vue'),
  }
];

export default routes;
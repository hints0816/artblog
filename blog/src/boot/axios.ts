import { boot } from 'quasar/wrappers';
import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import { LocalStorage, Loading, QSpinnerGrid, Notify, Cookies } from 'quasar';
import Router from '../router';

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance;
  }
}

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
const api = axios.create({
  baseURL: 'http://8.134.9.37:3000',
  // baseURL: 'http://localhost:3000',
  withCredentials: true,
});
// 请求拦截import { config } from 'process';

api.interceptors.request.use((cinfig: AxiosRequestConfig) => {
  Loading.show({
    spinner: QSpinnerGrid,
    spinnerSize: 100,
  })
  // const token :string = LocalStorage.getItem('token');
  const token :string = Cookies.get('token');
  cinfig.headers.Authorization = 'Bearer '+  token;
  return cinfig;
});
//响应拦截
api.interceptors.response.use(
  (response): any => {
    Loading.hide()
    if (response.status === 200) {
      if (response.data.status === 500 && response.data.message === 'token授权已过期,请重新登录') {
        LocalStorage.set('logged_in','no')
      } else if (response.config.url !== "/api/blog/getme" && response.data.status === 500 && response.data.message === 'Token不正确,请重新登录') {
        LocalStorage.set('logged_in','no')
        Notify.create({
          message: 'Please Login ArtBlog?',
          color: 'negative',
          icon: 'report_problem',
          position: 'top',
          actions: [
            { label: 'yes', color: 'white', handler: () => {
              return Router.push('/login')
            }}
          ]
        })
      } 
      return response.data;
    } else {
      return Promise.reject(response);
    }
  },
  (err) => {
    Loading.hide()
    return Promise.reject(err);
  }
);
export default boot(({ app }) => {
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios;
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api;
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
});

export { axios, api };

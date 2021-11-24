import { boot } from 'quasar/wrappers';
import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import { LocalStorage, Loading, QSpinnerGrid } from 'quasar'; 


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
  baseURL: 'http://47.119.167.128:3000',
  withCredentials: true,
});
// 请求拦截import { config } from 'process';

api.interceptors.request.use((cinfig: AxiosRequestConfig) => {
  Loading.show({
    spinner: QSpinnerGrid,
    spinnerSize: 100,
  })
  const token :string = LocalStorage.getItem('token');
  cinfig.headers.Authorization = 'Bearer '+  token;
  return cinfig;
});
//响应拦截
api.interceptors.response.use(
  (response): any => {
    // console.log(response);
    Loading.hide()
    if (response.status === 200) {
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
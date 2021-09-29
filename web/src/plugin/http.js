import axios from 'axios'

let Url = 'http://localhost:8080/api'

axios.defaults.baseURL = Url

axios.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
    return config
})

export default axios

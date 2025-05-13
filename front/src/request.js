import axios from 'axios'
import ElementUI from "element-ui";
import router from './router';

const request = axios.create({
    baseURL: '',
    timeout: 30000,
    withCredentials: true
})




request.interceptors.request.use(config => {
    config.headers['Content-Type'] = 'application/x-www-form-urlencoded;charset=utf-8';
    return config
}, error => {
    return Promise.reject(error)
});

request.interceptors.response.use(response => {
    
    return response;
}, error => {
    if (error.response.status === 401) {
        
        router.push('/login');
    }
    return Promise.reject(error);
});

export default request


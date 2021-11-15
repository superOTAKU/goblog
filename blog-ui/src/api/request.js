import axios from 'axios'
import store from '../store'
import jwtDecode from 'jwt-decode'

let service = axios.create({
    baseURL: process.env.BASE_URL,
    timeout: 10000
})

service.interceptors.request.use(requestConfig => {
    //登录注册接口直接请求，不需要带上token
    if (requestConfig.url == '/login' || requestConfig.url == '/register'
        || requestConfig.url == '/refreshToken') {
        return requestConfig;
    }
    //其他请求，带上token访问（getter负责判断过期时间）
    let token = store.getters.token
    if (token) {
        payload = jwtDecode(jwtToken)
        if (payload.exp < Date.now()) {
            return requestConfig;
        }
        requestConfig.headers['Authentication'] = jwtToken
    }
}, error => {

})

export default service
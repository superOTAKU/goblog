import {createStore} from 'vuex'
import jwtDecode from 'jwt-decode'
import {refreshToken} from '../api/login'

let store = createStore({
    state() {
        return {
            jwtToken: null //记录登录状态的JWT Token 
        }
    },
    mutations: {
        setJwtToken(state, jwtToken) {
            state.jwtToken = jwtToken
            localStorage.setItem("authToken", jwtToken)
        }
    },
    getters: {
        async token(state) {
            if (!state.jwtToken) {
                state.jwtToken = localStorage.getItem("authToken")
            }
            if (!state.jwtToken) {
                return null
            }
            payload = jwtDecode(state.jwtToken)
            now = Date.now()
            if (now < payload.exp) {
                return state.jwtToken
            } else if (now - payload.exp < 5 * 60 * 1000) {
                try {
                    state.jwtToken = await refreshToken(state.jwtToken)
                    return state.jwtToken
                } catch(e) {
                    console.log("refresh token error", e)
                    return null
                }
            } else {
                return null
            }
        }
    }
})

export default store
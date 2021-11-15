import service from "./request";

export function refreshToken(jwtToken) {
    return service.post({
        url: '/refreshToken',
        data: {
            token: jwtToken
        }
    })
}

export function login(username, password) {
    return service.post({
        url: '/login',
        data: {username, password}
    })
}

export function register(username, password) {
    return service.post({
        url: '/register',
        data: {username, password}
    })
}
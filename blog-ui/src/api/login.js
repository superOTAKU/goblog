import service from "./request";

export function refreshToken(jwtToken) {
    return service.post({
        url: '/refreshToken',
        data: {
            token: jwtToken
        }
    })
}
import {GET, POST} from '../../logic/request'

export const fetchAuthenticatedUser = () => {
    return GET(`http://${process.env.REACT_APP_AUTHENTICATION_SERVICE_HOST}/api/web/users/me`)
}

export const loginUser = (data) => {
    return POST(    `http://${process.env.REACT_APP_AUTHENTICATION_SERVICE_HOST}/api/users/login`, data)
}

export const signUpUser = (data) => {
    return POST(`http://${process.env.REACT_APP_AUTHENTICATION_SERVICE_HOST}/api/users/sign-up`, data)
}

export const logOut = (data) => {
    return POST(`http://${process.env.REACT_APP_AUTHENTICATION_SERVICE_HOST}/api/users/log-out`, data)
}
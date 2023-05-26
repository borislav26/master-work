import {createContext, useEffect, useState} from 'react'
import {AUTHENTICATION_API} from '../../endpoints/export'

export const AuthContext = createContext({user: null})

export const AuthProvider = ({children}) => {
    const [user, setUser] = useState(null)
    const [userLoading, setUserLoading] = useState(true)
    const [error, setError] = useState([])

    useEffect(()=> {
        fetchUser()
    }, [])

    const fetchUser =  async () => {
        const res = await AUTHENTICATION_API.fetchAuthenticatedUser()
        if (res) {
            console.log(res.data)
            setUser(res.data)
        }
        setUserLoading(false)
    }

    const logIn = async (email, password) => {
        const requestData = {
            email,
            password
        }

        const res = await AUTHENTICATION_API.loginUser(requestData)
        if (res.error) {
            setError([...error, res.error])
            return
        }
        setUser(res.data)
        localStorage.setItem("_master_work_token", res.data.token)
        window.location.href = '/landing-page'

    }

    const signUp = async (firstName, lastName, email, password, passwordConfirm) => {
        let errors = []
        if (!firstName || firstName === '') {
            errors = [...errors, 'First name is not valid.']
        }
        if (!lastName || lastName === '') {
            errors = [...errors, 'Last name is not valid.']
        }
        if (!email || email === '') {
            errors = [...errors, 'Email is not valid.']
        }
        if (password.length < 8) {
            errors = [...errors, "Password doesn't have correct format."]
        }
        if (password !== passwordConfirm) {
            errors = [...errors, 'Password and password confirm dont match.']
        }
        if (errors.length > 0) {
            setError(errors)
            return
        }

        const res = await AUTHENTICATION_API.signUpUser({firstName, lastName, email, password})
        if (res.error) {
            setError([...error, res.error])
            return
        }
        setUser(res.data)
        localStorage.setItem("_master_work_token", res.data.token)
        window.location.href = '/landing-page'
    }

    const logOut = async () => {
        const data = {
            email: user.email,
        }
        const res = await AUTHENTICATION_API.logOut(data)
        if (res.data) {
            setUser(null)
            localStorage.setItem("_master_work_token", res.data)
        }
    }

    return <AuthContext.Provider value={{user, setUser, logIn, signUp, userLoading, logOut, error}}>{children}</AuthContext.Provider>
}

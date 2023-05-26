import React, {useContext, useState} from 'react'
import './style.css'
import {AuthContext} from '../../context'
import {Link} from 'react-router-dom'

const Login = () => {
    const {logIn, error} = useContext(AuthContext)
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    const handleEmailChange = (event) => {
        setEmail(event.target.value)
    }

    const handlePasswordChange = (event) => {
        setPassword(event.target.value)
    }

    return (
        <>
            {error.length > 0 &&
                <div className="alert">
                    <span className="closebtn">&times;</span>
                    <strong>Error!</strong> {error.map(e => <div>{e}</div>)}
                </div>
            }
            <div className='Login__background'>
                <div className='Login__box'>
                    <h2>Login</h2>
                    <div className='Login__inputBox'>
                        <label>Email</label>
                        <input className='Login_input' onChange={handleEmailChange} value={email}/>
                        <label>Password</label>
                        <input className='Login_input' type='password' onChange={handlePasswordChange} value={password}/>
                    </div>
                    <div className='Login_buttonBox'>
                        <button className='Login_button' onClick={() => logIn(email, password)}>Login</button>
                        <div className='Login_signupLink'>
                            <Link to='/sign-up'>
                                You don't have account? Create one.
                            </Link>
                        </div>
                    </div>
                </div>
            </div>
        </>

    )
}

export default Login
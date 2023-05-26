import React, {useContext, useState} from 'react'
import './style.css'
import {AuthContext} from '../../context'
import {Link} from 'react-router-dom'

const SignUp = () => {
    const {signUp, error} = useContext(AuthContext)
    const [firstName, setFirstName] = useState('')
    const [lastName, setLastName] = useState('')
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [passwordConfirm, setPasswordConfirm] = useState('')


    const handleFirstNameChange =  (event) => {
        setFirstName(event.target.value)
    }

    const handleLastNameChange =  (event) => {
        setLastName(event.target.value)
    }

    const handleEmailChange = (event) => {
        setEmail(event.target.value)
    }

    const handlePasswordChange = (event) => {
        setPassword(event.target.value)
    }

    const handlePasswordConfirmChange = (event) => {
        setPasswordConfirm(event.target.value)
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
                <div className='Signup__box'>
                    <h2>Sign up</h2>
                    <div className='Login__inputBox'>
                        <label>First Name</label>
                        <input className='Login_input' onChange={handleFirstNameChange} value={firstName}/>
                        <label>Last Name</label>
                        <input className='Login_input' onChange={handleLastNameChange} value={lastName}/>
                        <label>Email</label>
                        <input className='Login_input' onChange={handleEmailChange} value={email}/>
                        <label>Password</label>
                        <input className='Login_input' type='password' onChange={handlePasswordChange} value={password}/>
                        <label>Password Confirm</label>
                        <input className='Login_input' type='password' onChange={handlePasswordConfirmChange} value={passwordConfirm}/>
                    </div>
                    <div className='Login_buttonBox'>
                        <button className='Login_button' onClick={() => signUp(firstName, lastName, email, password, passwordConfirm)}>Create account</button>
                        <div className='Login_signupLink'>
                            <Link to='/login'>
                                You already have account? Login.
                            </Link>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default SignUp
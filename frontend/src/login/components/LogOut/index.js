import React, {useContext} from 'react'
import './style.css'
import {AuthContext} from '../../context'

const LogOut = () => {
    const {logOut} = useContext(AuthContext)
    return (
        <div className='Logout'>
            <button className='Logout__button' type='button' onClick={logOut}>Log out</button>
        </div>
    )
}

export default LogOut
import React from 'react'
import LogOut from '../../../login/components/LogOut'
import './style.css'

const LandingPage = () => {

    const goToCreateFunnel = () => {
        window.location.href = '/create-cv'
    }

    const goToAllCVs = () => {
        window.location.href = '/all'
    }

    return (
        <>
            <div className='LandingPage'>
                <div className='LandingPage__option' onClick={goToAllCVs}>See existing CVs</div>
                <div className='LandingPage__option' onClick={goToCreateFunnel}>Create new CV</div>
            </div>
            <LogOut/>
        </>
    )
}

export default LandingPage
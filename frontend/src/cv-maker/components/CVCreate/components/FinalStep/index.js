import React from 'react'
import './style.css'

const FinalStep = ({handleCVCreation}) => {
    return (
        <div className='FinalStep'>
            <button type='button' className='FinalStep__button' onClick={handleCVCreation}>Create your CV</button>
        </div>
    )
}

export default FinalStep
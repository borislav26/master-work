import React from 'react'
import './style.css'

const ProfileInfo = ({description, handleDescriptionChange}) => {
    return (
        <div className='ProfileInfo'>
                <textarea value={description} className='ProfileInfo__textArea' placeholder='Add some info about you...' onChange={handleDescriptionChange}>
                </textarea>
        </div>
    )
}

export default ProfileInfo
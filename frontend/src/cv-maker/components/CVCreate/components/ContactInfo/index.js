import React from 'react'
import './style.css'

const ContactInfo = ({phoneNumber, jobTitle, website, githubProfile, linkedinProfile, setPhoneNumber, setJobTitle, setWebsite, setGithubProfile, setLinkedinProfile}) => {
    return (
        <div className='ContactInfo'>
            <div className='ContactInfo__oneLine'>
                <label className='ContactInfo__label'>Job Title</label>
                <input value={jobTitle} className='ContactInfo__input' onChange={(event) => setJobTitle(event.target.value)}/>
            </div>
            <div className='ContactInfo__oneLine'>
                <label className='ContactInfo__label'>Phone number</label>
                <input value={phoneNumber} className='ContactInfo__input' onChange={(event) => setPhoneNumber(event.target.value)}/>
            </div>
            <div className='ContactInfo__oneLine'>
                <label className='ContactInfo__label'>Website</label>
                <input value={website} className='ContactInfo__input' onChange={(event) => setWebsite(event.target.value)}/>
            </div>
            <div className='ContactInfo__oneLine'>
                <label className='ContactInfo__label'>Linkedin profile</label>
                <input value={linkedinProfile} className='ContactInfo__input' onChange={(event) => setLinkedinProfile(event.target.value)}/>
            </div>
            <div className='ContactInfo__oneLine'>
                <label className='ContactInfo__label'>Github profile</label>
                <input value={githubProfile} className='ContactInfo__input' onChange={(event) => setGithubProfile(event.target.value)}/>
            </div>
        </div>
    )
}

export default ContactInfo
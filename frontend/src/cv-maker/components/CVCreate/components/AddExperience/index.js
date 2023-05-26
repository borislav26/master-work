import React from 'react'
import './style.css'
import Plus from '../../../../../icons/Plus'
import Minus from '../../../../../icons/Minus'

const AddExperience = ({experiences, setExperience, addExperience, removeExperience}) => {
    let finalHTML = []

    for (let i = 0; i < experiences.length; i++) {
        finalHTML.push(
        <div className='AddExperience' key={i}>
            {i !== 0 && <div className='AddExperience__minus' onClick={() => removeExperience(i)}>
                <Minus height={48} width={48}/>
            </div>}
            <div className='AddExperience__center'>
                <div className='AddExperience__box'>
                    <div>
                        <label className='AddExperience__label'>Company</label>
                        <input
                            value={experiences[i].company}
                            onChange={(event) => setExperience(i,
                                {
                                    company: event.target.value,
                                    jobTitle: experiences[i].jobTitle,
                                    from: experiences[i].from,
                                    to: experiences[i].to,
                                    description: experiences[i].description
                                }
                            )}
                            className='AddExperience__input'
                        />
                    </div>
                    <div>
                        <label className='AddExperience__label'>Job title</label>
                        <input
                            value={experiences[i].jobTitle}
                            onChange={(event) => setExperience(i,
                                {
                                    company: experiences[i].company,
                                    jobTitle: event.target.value,
                                    from: experiences[i].from,
                                    to: experiences[i].to,
                                    description: experiences[i].description
                                }
                            )}
                            className='AddExperience__input'
                        />
                    </div>
                </div>
                <div className='AddExperience__box'>
                    <div>
                        <label className='AddExperience__label'>From</label>
                        <input
                            value={experiences[i].from}
                            onChange={(event) => setExperience(i,
                                {
                                    company: experiences[i].company,
                                    jobTitle: experiences[i].jobTitle,
                                    from: event.target.value,
                                    to: experiences[i].to,
                                    description: experiences[i].description
                                }
                            )}
                            className='AddExperience__input'
                        />
                    </div>
                    <div>
                        <label className='AddExperience__label'>To</label>
                        <input
                            value={experiences[i].to}
                            onChange={(event) => setExperience(i,
                                {
                                    company: experiences[i].company,
                                    jobTitle: experiences[i].jobTitle,
                                    from: experiences[i].from,
                                    to: event.target.value,
                                    description: experiences[i].description
                                }
                            )}
                            className='AddExperience__input'/>
                    </div>
                </div>
                <div className='AddExperience__textAreaBox'>
                    <label className='AddExperience__label'>Description</label>
                    <textarea
                        onChange={(event) => setExperience(i,
                            {
                                company: experiences[i].company,
                                jobTitle: experiences[i].jobTitle,
                                from: experiences[i].from,
                                to: experiences[i].to,
                                description: event.target.value
                            }
                        )}
                        className='AddExperience__textArea'
                        value={experiences[i].description}
                    ></textarea>
                </div>
            </div>
            <div className='AddExperience__plus' onClick={addExperience}>
                <Plus height={48} width={48}/>
            </div>
        </div>
        )
    }

    return (
        <div style={{display: 'flex', flexDirection: 'column'}}>
            {finalHTML.map(el=> el)}
        </div>
    )
}

export default AddExperience
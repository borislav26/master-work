import React from 'react'
import Minus from '../../../../../icons/Minus'
import Plus from '../../../../../icons/Plus'
import './style.css'

const AddEducation = ({educations, addEducation, removeEducation, setEducation}) => {
    let finalHTML = []

    for (let i = 0; i < educations.length; i++) {
        finalHTML.push(
            <div className='AddEducation' key={i}>
                {i !== 0 &&
                    <div className='AddEducation__minus' onClick={() => removeEducation(i)}>
                        <Minus height={48} width={48}/>
                    </div>
                }
                <div className='AddEducation__center'>
                    <div className='AddEducation__box'>
                        <div>
                            <label className='AddEducation__label'>Faculty</label>
                            <input
                                value={educations[i].faculty}
                                onChange={(event) => setEducation(i,
                                    {
                                        faculty: event.target.value,
                                        title: educations[i].title,
                                        from: educations[i].from,
                                        to: educations[i].to,
                                    }
                                )}
                                className='AddEducation__input'
                            />
                        </div>
                        <div>
                            <label className='AddEducation__label'>Title</label>
                            <input
                                value={educations[i].title}
                                onChange={(event) => setEducation(i,
                                    {
                                        faculty: educations[i].faculty,
                                        title: event.target.value,
                                        from: educations[i].from,
                                        to: educations[i].to,
                                    }
                                )}
                                className='AddEducation__input'
                            />
                        </div>
                    </div>
                    <div className='AddEducation__box'>
                        <div>
                            <label className='AddEducation__label'>From</label>
                            <input
                                value={educations[i].from}
                                onChange={(event) => setEducation(i,
                                    {
                                        faculty: educations[i].faculty,
                                        title: educations[i].title,
                                        from: event.target.value,
                                        to: educations[i].to,
                                    }
                                )}
                                className='AddEducation__input'
                            />
                        </div>
                        <div>
                            <label className='AddEducation__label'>To</label>
                            <input
                                value={educations[i].to}
                                onChange={(event) => setEducation(i,
                                    {
                                        faculty: educations[i].faculty,
                                        title: educations[i].title,
                                        from: educations[i].from,
                                        to: event.target.value,
                                    }
                                )}
                                className='AddEducation__input'
                            />
                        </div>
                    </div>
                </div>
                <div className='AddEducation__plus' onClick={addEducation}>
                    <Plus height={48} width={48}/>
                </div>
            </div>
        )
    }
    return (
        <div style={{display: 'flex', flexDirection: 'column'}}>
            {finalHTML.map(el => el)}
        </div>
    )
}

export default AddEducation
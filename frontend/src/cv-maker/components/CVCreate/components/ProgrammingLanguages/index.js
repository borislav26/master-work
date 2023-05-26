import React from 'react'
import Chip from '../Chip'
import './style.css'
import Trash from '../../../../../icons/Trash'

const ProgrammingLanguages = ({allProgrammingLanguages, setPercentageForLanguage}) => {
    return (
        <div className='ProgrammingLanguages'>
            {allProgrammingLanguages.map((p, index) => {
                return (
                        <div key={p.name} className='ProgrammingLanguages__box'>
                            <div className='ProgrammingLanguages__trashAndName'>
                                <Trash className='ProgrammingLanguages__trash'/>
                                <div className='ProgrammingLanguages__text'>{p.name}</div>
                            </div>
                            <input type='range' className='ProgrammingLanguages__input' value={p.percentage} onChange={(event) => setPercentageForLanguage(index, event.target.value)}/>
                            <div className='ProgrammingLanguages__percentage'>
                                <span className='ProgrammingLanguages__text'>{p.percentage}%</span>
                            </div>
                        </div>
                    )
                }
            )}
        </div>

    )
}

export default ProgrammingLanguages
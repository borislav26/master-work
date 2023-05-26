import React from 'react'
import Chip from '../Chip'
import './style.css'

const Languages = ({languages, setLanguages}) => {
    const handleSelectedChip = language => {
        const newLanguages = languages.map(el => {
            if (el.id === language.id) {
                return {...el, selected: !el.selected}
            }
            return el
        })
        setLanguages(newLanguages)
    }

    return <div className='Languages'>
        <div className='Languages__chips'>
            {languages.map(pl => <Chip key={pl.id} label={pl.name} blue={pl.selected} onClick={() => handleSelectedChip(pl)}/>)}
        </div>
    </div>
}

export default Languages
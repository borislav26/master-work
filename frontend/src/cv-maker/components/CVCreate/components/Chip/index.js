import React from 'react'
import './style.css'


const Chip = ({label, onClick, blue}) => {
    return <div className={`Chip ${blue ? 'Chip__blue' : ''}`} onClick={onClick}>{label}</div>
}

export default Chip
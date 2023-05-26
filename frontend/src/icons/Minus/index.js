import React from 'react'

const Minus = ({width, height, stroke}) => {
    return (
        <svg xmlns="http://www.w3.org/2000/svg" width={width || 24} height={height || 24} fill="none" viewBox="0 0 24 24">
            <path stroke={stroke || 'black'} strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5" d="M8 12h8m5 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"/>
        </svg>
    )
}

export default Minus
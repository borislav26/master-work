import React from 'react'
import './style.css'


const ChooseTemplate = ({pdfData, setChosenTemplate, setAllPdfData}) => {
    const handleTemplateSelect = fileName => {
        const newPdfData = pdfData.map(el => el.fileName === fileName ? ({...el, selected: true}) : ({...el, selected: false}))
        setAllPdfData(newPdfData)
        setChosenTemplate(fileName)
    }

    return (
        <div className='ChooseTemplate'>
            {pdfData.map(p => <img onClick={() => handleTemplateSelect(p.fileName)} key={p.fileName} className={`ChooseTemplate__img ${p.selected ? 'ChooseTemplate__imgSelected' : ''}`} src={`data:image/png;base64,${p.data}`} alt={"something"}/>)}
        </div>
    )

}

export default ChooseTemplate
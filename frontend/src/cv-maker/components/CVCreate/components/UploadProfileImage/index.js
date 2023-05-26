import React, {useRef, useState} from 'react'
import UploadCloud from '../../../../../icons/UploadCloud'
import './style.css'
import Cancel from '../../../../../icons/Cancel'


const UploadProfileImage = ({profilePhoto, setProfilePhoto}) => {
    const hiddenInput = useRef()
    const [fileNameVisible, setFileNameVisible] = useState(false)
    const [uploadedFileName, setUploadedFileName] = useState('')

    const handleImageUploadChange = (event) => {
        if (event.target.files.length > 0) {
            const reader = new FileReader()
            reader.onload = (e, pe) => {
                const newBytes = e.target.result.split(",").pop()
                setProfilePhoto(newBytes)
                setUploadedFileName(event.target.files[0].name)
            }
            reader.readAsDataURL(event.target.files[0])
        }

    }

    const handleChooseFileClick = () => {
        hiddenInput.current.click()
    }

    const removeImage = () => {
        setProfilePhoto(null)
    }

    const setFileNameSectionVisible = () => {
        setFileNameVisible(true)
    }

    const setFileNameSectionInvisible = () => {
        setFileNameVisible(false)
    }

    return (
        <div className='UploadProfileImage'>
            <div className='UploadProfileImage__wrapper' onMouseOver={setFileNameSectionVisible} onMouseLeave={setFileNameSectionInvisible}>
                <div className='UploadProfileImage__image'>
                    <img src={`data:image/png;base64,${profilePhoto}`} alt=''/>
                </div>
                <div className='UploadProfileImage__content'>
                    <div className='UploadProfileImage__icon'><UploadCloud width={96} height={96}/></div>
                    <div className='UploadProfileImage__noFileChosen'>No file chosen</div>
                </div>
                <div className='UploadProfileImage__cancel' onClick={removeImage}><Cancel/></div>
                {fileNameVisible && profilePhoto && <div className='UploadProfileImage__uploadedFileName'>{uploadedFileName}</div>}
            </div>
            <input accept='image/*' ref={hiddenInput} type='file' hidden onChange={handleImageUploadChange}/>
            <button type='button' className='UploadProfileImage__uploadButton' onClick={handleChooseFileClick}>Choose file</button>
        </div>
    )
}

export default UploadProfileImage
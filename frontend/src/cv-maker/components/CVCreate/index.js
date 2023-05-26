import React, {useEffect, useReducer, useState} from 'react'
import {cvCreateReducer, initialState, actions, MIN_STEP, MAX_STEP} from './reducer'
import LeftArrow from '../../../icons/LeftArrow'
import RightArrow from '../../../icons/RightArrow'
import './style.css'
import AddExperience from './components/AddExperience'
import ProgrammingLanguages from './components/ProgrammingLanguages'
import {CV_MAKER_API, PDF_GENERATOR_API} from '../../../endpoints/export'
import ChooseTemplate from './components/ChooseTemplate'
import ContactInfo from './components/ContactInfo'
import ProfileInfo from './components/ProfileInfo'
import Languages from './components/Languages'
import AddEducation from './components/AddEducation'
import UploadProfileImage from './components/UploadProfileImage'
import FinalStep from './components/FinalStep'
import Confetti from 'react-confetti'

const CVCreate = () => {
    const [currentState, dispatchFunction] = useReducer(cvCreateReducer, initialState)
    const [showConfetti, setShowConfetti] = useState(false)


    useEffect(() => {
        fetchAllProgrammingLanguages()
        fetchAllPDFExamples()
        fetchAllLanguages()
    }, [])

    const fetchAllProgrammingLanguages = async () => {
        const {data} = await CV_MAKER_API.fetchProgrammingLanguages()
        if (data) {
            const programmingLanguagesWithFlags = data.map(el => ({...el, percentage: 50}))
            dispatchFunction({type: actions.SET_ALL_PROGRAMMING_LANGUAGES, payload: programmingLanguagesWithFlags})
        }
    }

    const fetchAllPDFExamples =  async () => {
        const {data} = await PDF_GENERATOR_API.fetchAllPDFExamples()
        if (data) {
            const pdfDataWithFlags = data.map(el => ({...el, selected: false}))
            dispatchFunction({type: actions.SET_PDF_EXAMPLE_DATA, payload: pdfDataWithFlags})
        }
    }

    const fetchAllLanguages = async () => {
        const {data} = await CV_MAKER_API.fetchLanguages()
        if (data) {
            const languagesWithFlags = data.map(el => ({...el, selected: false}))
            dispatchFunction({type: actions.SET_ALL_LANGUAGES, payload: languagesWithFlags})
        }
    }

    const handleCVCreation = async () => {
        const copiedLanguages = [...currentState.allLanguages]
        const copiedProgrammingLanguages = [...currentState.allProgrammingLanguages]
        const filteredLanguages = copiedLanguages.filter(l => l.selected).map(l => ({id: l.id, name: l.name}))
        const filteredProgrammingLanguages = copiedProgrammingLanguages.map((pl) => ({id: pl.id, name: pl.name, percentage: +pl.percentage}))

        const reqData = {
            description: currentState.description,
            jobTitle: currentState.jobTitle,
            phoneNumber: currentState.phoneNumber,
            website: currentState.website,
            linkedinProfile: currentState.linkedinProfile,
            githubProfile: currentState.githubProfile,
            templateName: currentState.chosenTemplate,
            profilePhotoBytes: currentState.uploadedProfilePhotoBytes,
            experiences: currentState.experiences,
            languages: filteredLanguages,
            educations: currentState.educations,
            programmingLanguages:filteredProgrammingLanguages,

        }
       const {error} =  await CV_MAKER_API.createCV(reqData)
        if (error) {
            dispatchFunction({ type: actions.SET_ERRORS, payload: error })
            return
        }
        setShowConfetti(true)
        dispatchFunction({type: actions.SET_TITLE, payload: 'You successfully created your CV! You can check your email.'})
    }

    const moveLeft = () => {
        dispatchFunction({ type: actions.DECREASE_STEP })
    }

    const moveRight = () => {
        dispatchFunction({ type: actions.INCREASE_STEP })
    }

    const handleDescriptionChange = event => {
        dispatchFunction({ type: actions.SET_DESCRIPTION, payload: event.target.value})
    }

    const addExperience = () => {
        dispatchFunction({type: actions.ADD_EXPERIENCE})
    }

    const removeExperience = i => {
        dispatchFunction({type: actions.REMOVE_EXPERIENCE, index: i})
    }

    const setProgrammingLanguages = languages => {
        dispatchFunction({type: actions.SET_ALL_PROGRAMMING_LANGUAGES, payload: languages})
    }

    const setChosenTemplate = template => {
        dispatchFunction({type: actions.SET_CHOSEN_TEMPLATE, payload: template})
    }

    const setAllPdfData = data => {
        dispatchFunction({type: actions.SET_PDF_EXAMPLE_DATA, payload: data})
    }

    const setJobTitle = value => {
        dispatchFunction({type: actions.SET_JOB_TITLE, payload: value})
    }

    const setPhoneNumber = value => {
        dispatchFunction({type: actions.SET_PHONE_NUMBER, payload: value})
    }

    const setWebsite = value => {
        dispatchFunction({type: actions.SET_WEBSITE, payload: value})
    }

    const setLinkedinProfile = value => {
        dispatchFunction({type: actions.SET_LINKEDIN_PROFILE, payload: value})
    }

    const setGithubProfile = value => {
        dispatchFunction({type: actions.SET_GITHUB_PROFILE, payload: value})
    }

    const setExperience = (index, val) => {
        dispatchFunction({type: actions.SET_EXPERIENCES, index, payload: val})
    }

    const setLanguages = (value) => {
        dispatchFunction({type: actions.SET_ALL_LANGUAGES, payload: value})
    }

    const setPercentageForLanguage = (index, val) => {
        dispatchFunction({type: actions.SET_PERCENTAGE_FOR_PROGRAMMING_LANGUAGE, index, payload: val})
    }

    const addEducation = () => {
        dispatchFunction({type: actions.ADD_EDUCATION})
    }

    const removeEducation = i => {
        dispatchFunction({type: actions.REMOVE_EDUCATION, index: i})
    }

    const setEducation = (index, val) => {
        dispatchFunction({type: actions.SET_EDUCATION, index, payload: val})
    }

    const setProfilePhoto = (val) => {
        dispatchFunction({type: actions.SET_UPLOADED_PROFILE_PHOTO_BYTES, payload: val})
    }

    const cleanErrorMessages = () => {
        dispatchFunction({type: actions.SET_ERRORS, payload: []})
    }

    const goToLandingPage = () => {
        window.location.href = '/landing-page'
    }

    return (
        <div style={{position: 'relative'}}>
            {showConfetti && <Confetti width={window.innerWidth} height={window.innerHeight}/>}
            <div className='CVCreate__title'>
                <h1>{currentState.title}</h1>
            </div>
            {!showConfetti &&
                <div className='CVCreate__arrows'>
                    <div className={currentState.step > MIN_STEP ? 'CVCreate__leftArrow' : 'CVCreate__invisible_left'} onClick={moveLeft}>
                        <LeftArrow width={96} height={96}/>
                    </div>
                    <div className={currentState.step < MAX_STEP ? 'CVCreate__rightArrow' : 'CVCreate__invisible_right'} onClick={moveRight}>
                        <RightArrow width={96} height={96}/>
                    </div>
                </div>
            }
            <div className='CVCreate__title'>
                {currentState.step === 1 &&
                    <ChooseTemplate
                        pdfData={currentState.pdfExampleData}
                        setChosenTemplate={setChosenTemplate}
                        setAllPdfData={setAllPdfData}
                    />
                }
                {currentState.step === 2 &&
                    <UploadProfileImage profilePhoto={currentState.uploadedProfilePhotoBytes} setProfilePhoto={setProfilePhoto}/>
                }
                {currentState.step === 3 &&
                    <ContactInfo
                        jobTitle={currentState.jobTitle}
                        phoneNumber={currentState.phoneNumber}
                        website={currentState.website}
                        linkedinProfile={currentState.linkedinProfile}
                        githubProfile={currentState.githubProfile}
                        setJobTitle={setJobTitle}
                        setPhoneNumber={setPhoneNumber}
                        setWebsite={setWebsite}
                        setLinkedinProfile={setLinkedinProfile}
                        setGithubProfile={setGithubProfile}
                    />
                }
                {currentState.step === 4 &&
                    <ProfileInfo
                        description={currentState.description}
                        handleDescriptionChange={handleDescriptionChange}
                    />
                }
                {currentState.step === 5 &&
                    <AddExperience
                        experiences={currentState.experiences}
                        setExperience={setExperience}
                        addExperience={addExperience}
                        removeExperience={removeExperience}
                    />
                }
                {currentState.step === 6 &&
                    <AddEducation
                        addEducation={addEducation}
                        removeEducation={removeEducation}
                        educations={currentState.educations}
                        setEducation={setEducation}
                    />
                }
                {currentState.step === 7 &&
                    <Languages
                        languages={currentState.allLanguages}
                        setLanguages={setLanguages}
                    />
                }
                {currentState.step === 8 &&
                    <ProgrammingLanguages
                        allProgrammingLanguages={currentState.allProgrammingLanguages}
                        setPercentageForLanguage={setPercentageForLanguage}
                    />
                }
                {currentState.step === 9 && !showConfetti &&
                    <FinalStep handleCVCreation={handleCVCreation}/>
                }
                <div className='CVCreate__logout'>
                    <button className='CVCreate__logoutButton' type='button' onClick={goToLandingPage}>Go to Landing Page</button>
                </div>
            </div>
            {currentState.errors.length > 0 &&
                <div className="alert">
                    <span className="closebtn" onClick={cleanErrorMessages}>&times;</span>
                    <strong>Error!</strong> {currentState.errors.map(e => <div>{e}</div>)}
                </div>
            }
        </div>
    )
}

export default CVCreate
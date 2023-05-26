const actions = {
    INCREASE_STEP: 'INCREASE_STEP',
    DECREASE_STEP: 'DECREASE_STEP',
    SET_DESCRIPTION: 'SET_DESCRIPTION',
    ADD_EDUCATION: 'ADD_EDUCATION',
    REMOVE_EDUCATION: 'REMOVE_EDUCATION',
    SET_ALL_PROGRAMMING_LANGUAGES: 'SET_ALL_PROGRAMMING_LANGUAGES',
    SET_PDF_EXAMPLE_DATA: 'SET_PDF_EXAMPLE_DATA',
    SET_CHOSEN_TEMPLATE: 'SET_CHOSEN_TEMPLATE',
    SET_JOB_TITLE: 'SET_JOB_TITLE',
    SET_PHONE_NUMBER: 'SET_PHONE_NUMBER',
    SET_WEBSITE: 'SET_WEBSITE',
    SET_LINKEDIN_PROFILE: 'SET_LINKEDIN_PROFILE',
    SET_GITHUB_PROFILE: 'SET_GITHUB_PROFILE',
    SET_EXPERIENCES: 'SET_EXPERIENCES',
    ADD_EXPERIENCE: 'ADD_EXPERIENCE',
    REMOVE_EXPERIENCE: 'REMOVE_EXPERIENCE',
    SET_ALL_LANGUAGES: 'SET_ALL_LANGUAGES',
    SET_PERCENTAGE_FOR_PROGRAMMING_LANGUAGE: 'SET_PERCENTAGE_FOR_PROGRAMMING_LANGUAGE',
    SET_EDUCATION: 'SET_EDUCATION',
    SET_UPLOADED_PROFILE_PHOTO_BYTES: 'SET_UPLOADED_PROFILE_PHOTO_BYTES',
    SET_ERRORS: 'SET_ERRORS',
    SET_TITLE: 'SET_TITLE',
}

const initialState = {
    step: 1,
    description: "",
    title: "Choose your template",
    numberOfEducationItems: 1,
    allProgrammingLanguages: [],
    selectedProgrammingLanguages: [],
    pdfExampleData: [],
    chosenTemplate: "",
    jobTitle: "",
    phoneNumber: "",
    website: "",
    linkedinProfile: "",
    githubProfile: "",
    experiences: [{company: '', jobTitle: '', from: '', to: '', description: ''}],
    allLanguages: [],
    educations: [{faculty: '', from: '', to: '', title: ''}],
    uploadedProfilePhotoBytes: null,
    errors: [],
}

const MAX_STEP = 9
const MIN_STEP = 1

const cvCreateReducer = (state, action) => {
    let newState
    switch (action.type) {
        case actions.INCREASE_STEP:
            if (state.step === MAX_STEP) {
                return {...state}
            }
            const err = validateDataPerStep(state)
            if (err) {
                return {...state, errors: [...state.errors, err]}
            }
            newState = {...state, step: state.step + 1, errors: []}
            return handleTitleChangeOnSteps(newState)
        case actions.DECREASE_STEP:
            if (state.step === MIN_STEP) {
                return {...state}
            }
            newState = {...state, step: state.step - 1}
            return handleTitleChangeOnSteps(newState)
        case actions.SET_DESCRIPTION:
            return {...state, description: action.payload}
        case actions.ADD_EDUCATION:
            return {...state, educations: [...state.educations, {faculty: '', from: '', to: '', title: ''}]}
        case actions.REMOVE_EDUCATION:
            const filteredEdu = state.experiences.filter((val, index) => index !== action.index)
            return {...state, educations: filteredEdu}
        case actions.SET_ALL_PROGRAMMING_LANGUAGES:
            return {...state, allProgrammingLanguages: action.payload}
        case actions.SET_PDF_EXAMPLE_DATA:
            return {...state, pdfExampleData: action.payload}
        case actions.SET_CHOSEN_TEMPLATE:
            return {...state, chosenTemplate: action.payload}
        case actions.SET_JOB_TITLE:
            return {...state, jobTitle: action.payload}
        case actions.SET_PHONE_NUMBER:
            return {...state, phoneNumber: action.payload}
        case actions.SET_WEBSITE:
            return {...state, website: action.payload}
        case actions.SET_LINKEDIN_PROFILE:
            return {...state, linkedinProfile: action.payload}
        case actions.SET_GITHUB_PROFILE:
            return {...state, githubProfile: action.payload}
        case actions.ADD_EXPERIENCE:
            return {...state, experiences: [...state.experiences, {company: '', jobTitle: '', from: '', to: '', description: ''}]}
        case actions.REMOVE_EXPERIENCE:
            const filteredExp = state.experiences.filter((val, index) => index !== action.index)
            return {...state, experiences: filteredExp}
        case actions.SET_EXPERIENCES:
            const newExperiences = [...state.experiences].map((e, index) => index === action.index ? action.payload : e)
            return {...state, experiences: newExperiences}
        case actions.SET_ALL_LANGUAGES:
            return {...state, allLanguages: action.payload}
        case actions.SET_PERCENTAGE_FOR_PROGRAMMING_LANGUAGE:
            const mappedProgrammingLanguages = state.allProgrammingLanguages.map(
                (p, i) => i === action.index ? ({...p, percentage: action.payload}) : ({...p}))
            return {...state, allProgrammingLanguages: mappedProgrammingLanguages}
        case actions.SET_EDUCATION:
            const newEducations = [...state.educations].map((e, index) => index === action.index ? action.payload : e)
            return {...state, educations: newEducations}
        case actions.SET_UPLOADED_PROFILE_PHOTO_BYTES:
            return {...state, uploadedProfilePhotoBytes: action.payload}
        case actions.SET_ERRORS:
            return {...state, errors: action.payload}
        case actions.SET_TITLE:
            return {...state, title: action.payload}
        default:
            return {...state}
    }
}

const handleTitleChangeOnSteps = state => {
    if (state.step === 1) {
        return {...state, title: initialState.title}
    }
    if (state.step === 2) {
        return {...state, title: "Upload profile photo"}
    }
    if (state.step === 3) {
        return {...state, title: "Write some contact information"}
    }
    if (state.step === 4) {
        return {...state, title: "Add some description about you and your background"}
    }
    if (state.step === 5) {
        return {...state, title: "Add experience"}
    }
    if (state.step === 6) {
        return {...state, title: "Add education"}
    }
    if (state.step === 7) {
        return {...state, title: "Pick some languages"}
    }
    if (state.step === 8) {
        return {...state, title: "Set your knowledge about programming languages"}
    }
    return state
}

const validateDataPerStep = state => {
    if (state.step === 1 && !state.chosenTemplate) {
        return 'You need to select template.'
    }
    if (state.step === 2 && !state.uploadedProfilePhotoBytes) {
        return 'You need to choose profile photo.'
    }
    if (state.step === 3 && (!state.jobTitle || !state.phoneNumber)) {
        return 'You need to add at least job title and phone number.'
    }
    if (state.step === 4 && !state.description) {
        return 'You need to add profile description.'
    }
    if (state.step === 5) {
        if (state.experiences.length === 0) {
            return 'There is no experiences.'
        }
        for (const experience of state.experiences) {
            if (!experience.company || !experience.jobTitle || !experience.description) {
                return 'You need to add job title, company and description.'
            }
            if (!experience.from || isNaN(experience.from)) {
                return 'From is not a number.'
            }
            if (!experience.to || isNaN(experience.to)) {
                return 'To is not a number.'
            }
        }
    }
    if (state.step === 6) {
        if (state.educations.length === 0) {
            return 'There is no educations.'
        }
        for (const education of state.educations) {
            if (!education.faculty || !education.title) {
                return 'You need to add faculty and title.'
            }
            if (!education.from || isNaN(education.from)) {
                return 'From is not a number.'
            }
            if (!education.to || isNaN(education.to)) {
                return 'To is not a number.'
            }
        }
    }
    if (state.step === 7) {
        for(const l of state.allLanguages) {
            if (l.selected) {
                return
            }
        }
        return 'You need select at least one language.'
    }
}

export {
    cvCreateReducer,
    initialState,
    actions,
    MIN_STEP,
    MAX_STEP
}
import {GET, POST} from '../../logic/request'

export const fetchProgrammingLanguages = () => {
    return GET(`http://${process.env.REACT_APP_CV_MAKER_SERVICE_HOST}/api/programming_languages/all`)
}

export const fetchLanguages = () => {
    return GET(`http://${process.env.REACT_APP_CV_MAKER_SERVICE_HOST}/api/languages/all`)
}

export const createCV = (data) => {
    return POST(`http://${process.env.REACT_APP_CV_MAKER_SERVICE_HOST}/api/maker/create-cv`, data)
}

export const allGeneratedCVsForUser = () => {
    return GET(`http://${process.env.REACT_APP_CV_MAKER_SERVICE_HOST}/api/cv-data-byte/all-for-user`)
}

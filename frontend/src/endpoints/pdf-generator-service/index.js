import {GET} from '../../logic/request'

export const fetchAllPDFExamples = () => {
    return GET(`http://${process.env.REACT_APP_PDF_GENERATOR_SERVICE_HOST}/api/web/generator/examples`)
}


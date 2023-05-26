import React, {useEffect, useState} from 'react'
import {CV_MAKER_API} from '../../../endpoints/export'
import './style.css'

const AllUserGenerated = () => {
    const [allCVs, setAllCVs] = useState([])

    const fetchAllGeneratedCVs = async () => {
        const {data} = await CV_MAKER_API.allGeneratedCVsForUser()
        if (data) {
            setAllCVs(data)
            console.log(data)
        }
    }

    const goToLandingPage = () => {
        window.location.href = '/landing-page'
    }

    useEffect(() => {
        fetchAllGeneratedCVs()
    }, [])

    return (
        <div className='AllUserGenerated'>
            {allCVs.map(p => <iframe style={{width: '600px', height: '800px'}} key={p.id} src={`data:application/pdf;base64,${p.bytes}`} type="application/pdf"></iframe>)}
            <div className='CVCreate__logout'>
                <button className='CVCreate__logoutButton' type='button' onClick={goToLandingPage}>Go to Landing Page</button>
            </div>
        </div>
    )
}

export default AllUserGenerated

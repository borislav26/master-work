import React from 'react'
import axios from 'axios'

const userMessageOrDefault = (error, defaultErrorMessage = 'Request failed for an unknown reason') => {
    return error?.response?.data?.userMessage
        ? error.response.data.userMessage
        : defaultErrorMessage
}

const makeGETRequest = async (url) => {
    try {
        return await Axios.get(url)
    } catch (e) {
        return {
            error: userMessageOrDefault(e),
            status: e.status
        }
    }
}

const makePOSTRequest = async (url, data) => {
    try {
        return await Axios.post(url, data)
    } catch (e) {
        return {
            error: userMessageOrDefault(e),
            status: e.status
        }
    }
}

const Axios = axios.create({
    headers: {
        Authorization: localStorage.getItem("_master_work_token")
    },
    withCredentials: true
})

export const GET = async (url) => {
    return makeGETRequest(url)
}

export const POST = async (url, data) => {
    return makePOSTRequest(url, data)
}

export const PATCH = async (url, data) => {
    return await Axios.patch(url, data)
}

export const PUT = async (url, data) => {
    return await Axios.put(url, data)
}
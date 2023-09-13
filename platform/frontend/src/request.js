import axios from "axios";

const devBaseURL = "http://localhost:8080/api"; //开发环境
const proBaseURL = "http://localhost:8080/api"; //生产环境
export const BASE_URL = process.env.NODE_ENV === 'development' ? devBaseURL: proBaseURL;
export const TIMEOUT = 50000;

export const request=axios.create({
    baseURL:BASE_URL,
    timeout:TIMEOUT
})
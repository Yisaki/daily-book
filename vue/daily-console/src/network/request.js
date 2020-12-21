import axios from 'axios'

function request(config) {
    const instance = axios.create({
        baseURL: 'http://localhost:8963/api',
        //baseURL: 'http://81.69.59.137:80/api',
        timeout: 5000
    });

    return instance(config);
};

export default request;

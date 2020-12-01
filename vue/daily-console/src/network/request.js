import axios from 'axios'

function request(config) {
    const instance = axios.create({
        baseURL: 'http://localhost:8963',
        timeout: 5000
    });

    return instance(config);
};

export default request;

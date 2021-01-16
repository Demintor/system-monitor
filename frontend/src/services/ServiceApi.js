import axios from 'axios';

class ServiceApi {
    static createApi() {
        let api = axios.create({
            baseURL: 'http://localhost:8080/v2/testapi/',
        });
        return api;
    }
    static async get(route = "", params = {}) {
        let api = this.createApi();
        try {
            let response = await api.get(route, params);
            return response;
        }
        catch(error) {
            throw new Error(error);
        }
    }
    static async post(route = "", postData = {}) {
        let api = this.createApi();
        try {
            let response = await api.post(route, postData);
            return response;
        }
        catch(error) {
            throw new Error(error);
        }
    }
}

export default ServiceApi;
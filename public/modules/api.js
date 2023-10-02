import {Requests} from "./requests.js";
import {restEndpoints} from "../configs/rest_config.js";
import {backendUrl} from "../configs/rest_config.js";

export class Api extends Requests {
    async getUserProfile(id) {
        const endpoint = restEndpoints['getProfile'];
        const url = backendUrl + endpoint.url.replace("{id}", id);
        return this.make_request(url, endpoint.method);
    }

    async getUserPosts(id) {
        const endpoint = restEndpoints["getPosts"];
        const url =  backendUrl + endpoint.url.replace("{id}", id);
        return  this.make_request(url, endpoint.method);
    }
}
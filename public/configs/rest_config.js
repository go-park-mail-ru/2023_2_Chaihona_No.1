const GET_METHOD = 'GET';
const POST_METHOD = 'POST';

export const backendUrl = "https://127.0.0.1:5000/api/v1"

export const restEndpoints = {
    getPosts:  {
        url: "/profile/{id}/post",
        method: GET_METHOD
    },
    getProfile: {
        url: "/profile/{id}",
        method: GET_METHOD
    },
}
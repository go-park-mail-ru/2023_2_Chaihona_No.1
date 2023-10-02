const noBodyRequests = ['GET', 'HEAD']

export class Requests {
    async make_request(url, method, data=null) {
        const params = {
            method,
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json',
            },
        };

        if (!noBodyRequests.includes(method)) {
            params.body = JSON.stringify(data);
        }

        const response = await fetch(url, params);

        const responseJson = await response.json();

        return {
            status: response.status,
            responseJson
        };
    }
}
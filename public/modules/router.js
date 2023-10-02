import {routes} from "../configs/link_config.js";
import {profile1} from "../test_data/test_data.js";

export class Router {
    constructor() {
        window.addEventListener('click', (e) => {
            const target = e.target.getAttribute("href");
            if (target !== null) {
                e.preventDefault();
                this.goTo(target);
            }
        });

        window.addEventListener('popstate', () => {
            const path = window.location.pathname;
            this.goTo(path);
        });
    }

    goTo(path) {
        const renderer = path.replace(/[/0-9]*/g, "");
        let route = routes[renderer];
        window.history.pushState(null, null, path);
        if (route === undefined) {
            route = routes[404]
        }
        route.render();
    }
}
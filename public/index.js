import {Router} from "./modules/router.js";
import Navbar from "./components/Navbar/Navbar.js";
import navbar from "./components/Navbar/Navbar.js";

navbar();
const router = new Router();
window.router = router;
router.goTo('/login');
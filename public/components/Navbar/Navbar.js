export default async() => {
    const navbarElement = document.querySelector('#navbar');
    navbarElement.innerHTML = '';
    navbarElement.innerHTML = Handlebars.templates.navbar();
    const loginBtn = document.querySelector("#login-btn");

    loginBtn.addEventListener("click", function (e) {
        window.router.goTo('/login');
    });
    const regBtn = document.querySelector("#reg-btn")
    regBtn.addEventListener("click", function (e) {
        window.router.goTo('/register');
    });


}

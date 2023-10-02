export default async() => {
    const rootElement = document.querySelector('#root');
    rootElement.innerHTML = '';
    rootElement.innerHTML = Handlebars.templates.register();

    const regBtn = document.querySelector(".regbutton")
    regBtn.addEventListener("click", function (e) {
       const pass = document.querySelector("#password");
       const repeatPass = document.querySelector("#repeat-password");
       if (pass.value !== repeatPass.value) {
           e.preventDefault();
           const err = document.querySelector(".error")
           err.textContent = "Пароли не совпадают"
       }
    });
}

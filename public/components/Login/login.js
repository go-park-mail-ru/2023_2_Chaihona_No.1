


export default async() => {
    const rootElement = document.querySelector('#root');
    rootElement.innerHTML = '';
    rootElement.innerHTML = Handlebars.templates.login();
}
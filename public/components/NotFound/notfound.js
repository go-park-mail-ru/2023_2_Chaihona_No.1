/**
 * Функция отрисовки страницы 404
 */
export default () => {
    const rootElement = document.querySelector('#root');
    rootElement.innerHTML = Handlebars.templates.notfound();
};
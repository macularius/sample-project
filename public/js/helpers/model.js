// вспомогательный класс для совершения запросов
export default class Model {
    // метод для совершения get запроса
    get(url) {
        return new Promise(function(resolve, reject) {
            let xhr = new XMLHttpRequest();
            xhr.open('GET', url, false);
            xhr.send()
            if (xhr.status != 200) {
                webix.message(xhr.status + ': ' + xhr.statusText, 'error');
                reject()
            } else {
                resolve(xhr.responseText);
            }
        })
    }

    // метод для совершения post запроса
    post(url, params) {
        let xhr = new XMLHttpRequest();
        xhr.open('POST', url, false);
        xhr.send(params)
        if (xhr.status != 200) {
            webix.message(xhr.status + ': ' + xhr.statusText, 'error');
            reject()
        } else {
            resolve(xhr.responseText);
        }
    }
}
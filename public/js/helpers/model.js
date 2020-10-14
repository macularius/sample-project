// вспомогательный класс для совершения запросов
export default class Model {
    // метод для совершения get запроса
    get(url) {
        return new Promise(function (resolve, reject) {
            let xhr = new XMLHttpRequest();
            xhr.open('GET', url, false);
            xhr.send()
            xhr.onload = () => {
                // проверка статуса HTTP запроса
                if (xhr.status != 200) {
                    webix.message(xhr.status + ': ' + xhr.statusText, 'error');
                    reject()
                } else {
                    console.log(xhr.responseText)

                    // валидация статуса ответа сервера
                    if (!xhr.responseText.status) {
                        webix.message('Не удалось совершить запрос', 'error');
                        console.error(`GET xhr.responseText.status is ${xhr.responseText.status}`);
                        reject()
                    }

                    // проверка статуса ответа сервера
                    switch (xhr.responseText.status) {
                        case RESULT_STATE.SUCCESS: // положительный результат запроса
                            resolve(xhr.responseText.data);
                            return;
                        case RESULT_STATE.FAILED: // отрицательный результат запроса
                            webix.message('Не удалось совершить запрос', 'error');
                            console.error(`GET ${xhr.responseText.error}`);
                            reject();
                            return;
                        default: // ошибка при получении результата запроса
                            webix.message('Не удалось совершить запрос', 'error');
                            console.error(`GET Статус ответа сервера не соответствует ожидаемым значениям, xhr.responseText.status is ${xhr.responseText.status}`);
                            reject();
                            return;
                    }
                }
            }
        })
    }

    // метод для совершения post запроса
    post(url, params) {
        return new Promise(function (resolve, reject) {
            let xhr = new XMLHttpRequest();
            xhr.open('POST', url, false);
            xhr.send(params)
            xhr.onload = () => {
                // проверка статуса HTTP запроса
                if (xhr.status != 200) {
                    webix.message(xhr.status + ': ' + xhr.statusText, 'error');
                    reject()
                } else {
                    // валидация статуса ответа сервера
                    if (!xhr.responseText.status) {
                        webix.message('Не удалось совершить запрос', 'error');
                        console.error(`GET xhr.responseText.status is ${xhr.responseText.status}`);
                        reject()
                    }

                    // проверка статуса ответа сервера
                    switch (xhr.responseText.status) {
                        case RESULT_STATE.SUCCESS: // положительный результат запроса
                            resolve(xhr.responseText.data);
                            return;
                        case RESULT_STATE.FAILED: // отрицательный результат запроса
                            webix.message('Не удалось совершить запрос', 'error');
                            console.error(`GET ${xhr.responseText.error}`);
                            reject();
                            return;
                        default: // ошибка при получении результата запроса
                            webix.message('Не удалось совершить запрос', 'error');
                            console.error(`GET Статус ответа сервера не соответствует ожидаемым значениям, xhr.responseText.status is ${xhr.responseText.status}`);
                            reject();
                            return;
                    }
                }
            }
        })
    }
}

// константы допустимых значений статуса ответа
const RESULT_STATE = {
    SUCCESS: 'succes',
    FAILED: 'failed',
}
import authModel from "../src/models/authModel.js"

// функция проверяет авторизованность клиента
// вызывает callback функцию с boolean аргументом - признаком авторизованности
export function checkAuth(callback) {
    authModel.check().then((isAuthorize) => {
        callback(isAuthorize)
    })
}
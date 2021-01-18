import Model from '../../helpers/model.js';

// AuthModel объект для авторизации
class AuthModel extends Model {
    constructor() {
        super()
    }

    // вход в учетную запись
    login(user) {
        return this.post('/user/login', user)
    }

    // выход из учетной записи
    logout() {
        return this.get('/user/logout')
    }

    // получение сотруднкиа текущего пользователя
    getCurrentEmployee() {
        return this.get('/user/employee')
    }

    // проверка авторизации
    check() {
        return this.post('/user/authorize/check')
    }
}
const authModel  = new AuthModel();
export default authModel
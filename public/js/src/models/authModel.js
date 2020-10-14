import Model from "../../helpers/model.js";

// AuthModel объект для авторизации
class AuthModel extends Model {
    constructor() {
        super()
    }

    // выход из учетной записи
    logout() {
        return this.get('/user/logout')
    }

    // плучение сотруднкиа текущего пользователя
    getCurrentEmployee() {
        return this.get('/user/employee')
    }
}
const authModel  = new AuthModel();
export default authModel
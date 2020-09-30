import Model from "../../helpers/model.js";

// AuthModel объект для авторизации
class AuthModel extends Model {
    constructor() {}

    // выход из учетной записи
    logout() {
        return this.get('/user/logout')
    }
}
const authModel  = new AuthModel();
export default authModel
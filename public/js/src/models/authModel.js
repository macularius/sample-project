import Model from "../../helpers/model.js";
import { Employee } from "./entities/employee.js";

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
        return new Promise((resolve, reject) => {
            setTimeout(() => {
                resolve(new Employee(1, 'Коваценко', 'Игорь', 'Николаевич', 'инженер-программист', '', ''))
            }, 5000);
        })
    }
}
const authModel  = new AuthModel();
export default authModel
import Model from '../../helpers/model.js';

/// EmployeeModel объект для работы(CRUD) с данными
class EmployeeModel extends Model {
    constructor() {
        super()
    }

    // поолучение всех сотрудник
    getEmployees() {
        return this.get('/employee/all')
    }

    // поолучение сотрудника по его ID
    getEmployeeByID(id) {
        return this.get(`/employee/${id}`)
    }

    // создание сотрудника
    createEmployee(employee) {
        return this.post('/employee/create', employee)
    }

    // изменение сотрудника
    updateEmployee(employee) {
        return this.post('/employee/update', employee)
    }

    // удаление сотрудника
    deleteEmployee(employee) {
        return this.post('/employee/delete', employee)
    }

    // получение книг читательского билета сотрудника по ID
    getCardByEmployeeID(id) {
        return this.get(`/employee/${id}/card`)
    }
}
const employeeModel  = new EmployeeModel();
export default employeeModel
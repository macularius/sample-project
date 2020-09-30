// EmployeeModel объект для работы(CRUD) с данными
class EmployeeModel extends Model {
    constructor() {}

    // поолучение всех сотрудник
    getEmployees() {
        return new Promise(function(resolve, reject) {

        })
    }

    // поолучение сотрудника по его ID
    getEmployeeByID(id) {
        return new Promise(function(resolve, reject) {

        })
    }

    // создание сотрудника
    createEmployee(employee) {
        return new Promise(function(resolve, reject) {

        })
    }

    // удаление сотрудника
    deleteEmployee(employee) {
        return new Promise(function(resolve, reject) {

        })
    }
}
const employeeModel  = new EmployeeModel();
export default employeeModel
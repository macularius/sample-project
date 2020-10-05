import Model from "../../helpers/model.js";
import { Employee } from "./entities/employee.js";
import { Position } from "./entities/position.js";

/// EmployeeModel объект для работы(CRUD) с данными
class EmployeeModel extends Model {
    constructor() {
        super()

        this.data = new Map();
        this.data.set(1, new Employee(1, 'lastname1', 'firstname1', 'middlename1', 'position1', 'number1', 'email1'));
        this.data.set(2, new Employee(2, 'lastname2', 'firstname2', 'middlename2', 'position2', 'number2', 'email2'));
        this.data.set(3, new Employee(3, 'lastname3', 'firstname3', 'middlename3', 'position3', 'number3', 'email3'));
        this.data.set(4, new Employee(4, 'lastname4', 'firstname4', 'middlename4', 'position4', 'number4', 'email4'));
    }

    // поолучение всех сотрудник
    getEmployees() {
        return new Promise((resolve, reject) => {
            let employees = []

            for (let employee of this.data.values()) {
                employees.push(employee)
            }

            resolve(employees)
        })
    }

    // поолучение сотрудника по его ID
    getEmployeeByID(id) {
        return new Promise((resolve, reject) => {
            resolve(this.data.get(id))
        })
    }

    // создание сотрудника
    createEmployee(employee) {
        return new Promise((resolve, reject) => {
            let id

            for (let key of this.data.keys()) {
                id = key
            }
            id++

            employee.id = id
            this.data.set(id, employee)
            resolve(this.data.get(employee.id))
        })
    }

    // изменение сотрудника
    updateEmployee(employee) {
        return new Promise((resolve, reject) => {
            this.data.set(employee.id, employee)
            resolve(this.data.get(employee.id))
        })
    }

    // удаление сотрудника
    deleteEmployee(employee) {
        return new Promise((resolve, reject) => {
            this.data.delete(employee.id)
            resolve()
        })
    }

    // получение должностей
    getPositions() {
        return new Promise((resolve, reject) => {
            resolve([
                new Position(1, 'Прогарммист'),
                new Position(2, 'Старший программист')
            ])
        })
    }
}
const employeeModel  = new EmployeeModel();
export default employeeModel
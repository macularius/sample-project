// Employee класс для представления сущности сотрудника
export class Employee {
    constructor(id, lastname, firstname, middlename, position, phoneNumber, email) {
        this.ID = id
        this.lastname = lastname
        this.firstname = firstname
        this.middlename = middlename
        this.position = position
        this.phoneNumber = phoneNumber
        this.email = email
    }
}
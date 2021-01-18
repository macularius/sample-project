// Event класс для представления сущности события выдачи/сдачи книги
export class Event {
    constructor(id, book, employee, type, date) {
        this.ID = id
        this.book = book
        this.employee = employee
        this.type = type
        this.date = date
    }
}

// допустимые типы события
export let EVENT_TYPE = {
    give: 'выдача',
    take: 'сдача',
}
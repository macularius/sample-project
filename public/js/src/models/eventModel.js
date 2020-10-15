import Model from "../../helpers/model.js";
import { Book } from "./entities/book.js";
import { Employee } from "./entities/employee.js";
import { Event, EVENT_TYPE } from "./entities/event.js";

/// EventModel объект для работы(CRUD) с данными
class EventModel extends Model {
    constructor() {
        super()
    }

    // получение всех событий
    getEvents() {
        return this.get('/event/all')
    }

    // получение событий по id книги
    getEventsByBook(bookID) {
        return this.get(`/event/byBook/${bookID}`)
    }

    // поолучение событий по id сотрудника
    getEventsByEmployee(employeeID) {
        return this.get(`/event/byEmployee/${employeeID}`)
    }

    // создание события выдачи
    createGiveEvent(bookID, employeeID) {
        // формирование объекта события для запроа
        let event = new Event()
        event.book = new Book()
        event.book.ID = bookID
        event.employee = new Employee()
        event.employee.ID = employeeID
        event.type = EVENT_TYPE.give

        return this.post('/event/create', event)
    }

    // создание события сдачи
    createTakeEvent(bookID) {
        // формирование объекта события для запроа
        let event = new Event()
        event.book = new Book()
        event.book.ID = bookID
        event.type = EVENT_TYPE.take

        return this.post('/event/create', event)
    }

    // получение последнего события выдачи по сотруднику и книге
    getLastGiveEventByBookID(bookID, employeeID) {
        // формирование объекта события для запроа
        let event = new Event()
        event.book = new Book()
        event.book.ID = bookID
        event.type = EVENT_TYPE.take

        return new Promise((resolve, reject) => {
            resolve(new Event(
                1,
                new Book(1, 'isbn1', 'name1', 'author1', 'publisher1', '2020', 'status'),
                new Employee(1, 'lastname1', 'firstname1', 'middlename1', 'position1', 'number1', 'email1'),
                'выдача',
                new Date(Date.now())
            ))
        })
        return this.post('/event/give/last', event)
    }
}
const eventModel = new EventModel();
export default eventModel
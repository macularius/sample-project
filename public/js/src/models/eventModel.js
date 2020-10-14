import Model from "../../helpers/model.js";
import { Event } from "../../models/entities/event.js";

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

    // создание события
    createEvent(bookID, employeeID) {
        // формирование объекта события для запроа
        let event = new type(Event)
        event.book.ID = bookID
        event.employee.ID = employeeID

        return this.get('/event/create', event)
    }
}
const eventModel = new EventModel();
export default eventModel
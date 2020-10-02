import Model from "../../helpers/model.js";
import { Event } from "./entities/event.js";
import { Book } from "./entities/book.js";
import { Employee } from "./entities/employee.js";

/// EventModel объект для работы(CRUD) с данными
class EventModel extends Model {
    constructor() {
        super()

        this.data = new Map();
        this.data.set(
            1,
            new Event(1, new Book(1, 'isbn1', 'name1', 'author1', 'publisher1', '2020', 'status'),
                new Employee(1, 'lastname1', 'firstname1', 'middlename1', 'position1', 'number1', 'email1'),
                'type',
                new Date(Date.now()))
        );
    }

    // получение всех событий
    getEvents() {
        return new Promise((resolve, reject) => {
            let events = []

            for (let event of this.data.values()) {
                events.push(event)
            }

            resolve(events)
        })
    }

    // получение событий по id книги
    getEventsByBook(bookID) {
        return new Promise((resolve, reject) => {
            let events = []

            this.data.values().map((value) => {
                if (value.book.ID === bookID) {
                    events.push(value)
                }
            })

            resolve(events)
        })
    }

    // поолучение событий по id сотрудника
    getEventsByEmployee(employeeID) {
        return new Promise((resolve, reject) => {
            let events = []

            this.data.values().map((value) => {
                if (value.employee.ID === employeeID) {
                    events.push(value)
                }
            })

            resolve(events)
        })
    }

    // создание события
    createEvent(event) {
        return new Promise((resolve, reject) => {
            let id

            for (let key of this.data.keys()) {
                id = key
            }
            id++

            event.id = id
            this.data.set(id, event)
            resolve(this.data.get(event.id))
        })
    }
}
const eventModel = new EventModel();
export default eventModel
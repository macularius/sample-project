// EventModel объект для работы(CRUD) с данными
class EventModel extends Model {
    constructor() {}

    // поолучение всех событий
    getEvents() {
        return new Promise(function(resolve, reject) {

        })
    }

    // поолучение событий по id книги
    getEventsByBook(bookID) {
        return new Promise(function(resolve, reject) {

        })
    }

    // поолучение событий по id сотрудника
    getEventsByEmployee(employeeID) {
        return new Promise(function(resolve, reject) {

        })
    }

    // создание события
    createEvent(event) {
        return new Promise(function(resolve, reject) {

        })
    }
}
const eventModel  = new EventModel();
export default eventModel
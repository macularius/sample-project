import { BookTabView, BookTabContextMenu } from "./BookTabView.js"
import { CBookWindow, BOOK_WINDOW_TYPE } from "./bookWindow/CBookWindow.js";
import bookModel from "../../models/bookModel.js";
import eventModel from "../../models/eventModel.js";
import employeeModel from "../../models/employeeModel.js";
import { Book } from "../../models/entities/book.js";

// класс таба "Книги"
export class CBookTab {
    constructor() {
        this.view       // объект для быстрого доступа к представлениям
        this.window     // экземпляр окна для работы с книгами
    }

    // метод инициализации компонента
    init() {
        this.window = new CBookWindow(); // инициализация компонента окна
        this.window.init(
            () => { this.refreshTable() }
        ) // вызова инициализации компонента окна
    }

    // метод получения webix конфигурации компонента
    config() {
        // т.к. window и popup расположены не в дереве приложения, а поверх слоев, его нужно отрисовывать отдельно
        webix.ui(this.window.config())
        // отложенное заполнение массива сотрудников в сабменю
        let names = []
        employeeModel.getEmployees().then((employees) => {
            employees.map((employee) => {
                names.push({ ID: employee.ID, value: `${employee.lastname} ${employee.firstname}` })
            })
        })
        webix.ui(BookTabContextMenu(names))

        // вызов функции представления
        return BookTabView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {
        // инициализация используемых представлений
        this.view = {
            datatable: $$('bookTabDatatable'),
            datatableContextMenu: $$('bookTabDatatableContextMenu'),
        }

        // инициализация обработчиков событий модального окна
        this.window.attachEvents()

        // прикрепление контекстного меню к таблице
        this.view.datatableContextMenu.attachTo(this.view.datatable)

        // загрузка первичных данных в таблицу
        this.refreshTable()

        // обработка события нажатия на пункт контекстного меню
        this.view.datatableContextMenu.attachEvent('onMenuItemClick', (itemID) => {
            // получение значения пункта, на которое произошло нажатие
            

            // проверка вложенности выбранного пункта меню
            if (!this.view.datatableContextMenu.getItem(itemID)) {
                this.handleSubMenu(itemID)
            } else {
                let item = this.view.datatableContextMenu.getItem(itemID).value
                this.handleContextMenu(item)
            }

        });
    }

    // обработка выбора в контекстном меню
    handleContextMenu(item) {
        // получение выделенного элемента
        let selected = this.view.datatable.getSelectedItem()

        switch (item) {
            case BOOK_CONTEXT_MENU.add: // добавление книги
                this.window.parse(new Book())
                this.window.switch(BOOK_WINDOW_TYPE.create)
                break;
            case BOOK_CONTEXT_MENU.edit: // редактирование выделленой книги
                if (!selected) {
                    webix.message('Выделите строку')
                    return
                }
                if (!selected.ID) {
                    console.error('Incorrect ID of item:', selected.ID)
                    return
                }
                bookModel.getBookByID(selected.ID).then((book) => {
                    this.window.parse(book)
                    this.window.switch(BOOK_WINDOW_TYPE.update)
                })
                break;
            case BOOK_CONTEXT_MENU.remove: // удаление выделенной книги
                if (!selected) {
                    webix.message('Выделите строку')
                    return
                }
                if (!selected.ID) {
                    console.error('id of item is ', selected.ID)
                    return
                }
                bookModel.getBookByID(selected.ID).then((book) => {
                    this.window.parse(book)
                    this.window.switch(BOOK_WINDOW_TYPE.delete)
                })
                break;
            default:
                console.error(`Неизвестное значение пункта меню: ${item}.`);
                break;
        }
    }

    // обработка выбора в submenu
    handleSubMenu(empItem) {
        // получения сотрудника из submenu
        let submenu = $$(this.view.datatableContextMenu.getItem(BOOK_CONTEXT_MENU.give).submenu)
        let employee = submenu.getItem(empItem)

        // получение выделенного элемента
        let book = this.view.datatable.getSelectedItem()
        if (!book) {
            webix.message('Выделите строку')
            return
        }
        if (!book.ID) {
            console.error('Incorrect ID of item:', book.ID)
            return
        }

        eventModel.createEvent(book.ID, employee.ID)
    }

    // функция обновления таблицы книг
    refreshTable(books) {
        if (books) {
            this.view.datatable.clearAll()
            this.view.datatable.parse(books)
            return
        } else {
            bookModel.getBooks().then((books) => {
                this.view.datatable.clearAll()
                this.view.datatable.parse(books)
            })
        }
    }

    // метод отображения таба с фильтрацией по книге
    showByBookID(bookID) {
        bookModel.getBookByID(bookID).then((book) => {
            // применение фильтров
            this.view.datatable.getFilter('ISBN').value = book.ISBN;
            this.view.datatable.filterByAll();

            // выделение нужной строки
            for (let rowID = 0; rowID < this.view.datatable.serialize().length; rowID++) {
                let item = this.view.datatable.serialize()[rowID]

                if (item.ID === bookID) {
                    this.view.datatable.select(item.id)
                    break
                }
            }
        })
    }
}

// допустимые значения пунктов контекстного меню таба Книги
export const BOOK_CONTEXT_MENU = {
    give: 'Выдать',
    add: 'Добавить',
    edit: 'Изменить',
    remove: 'Удалить'
}
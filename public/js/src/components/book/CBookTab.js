import { BookTabView, BookTabContextMenu } from "./BookTabView.js"
import { CBookWindow, BOOK_WINDOW_TYPE } from "./bookWindow/CBookWindow.js";
import bookModel from "../../models/bookModel.js";
import { Book } from "../../models/entities/book.js";

// класс таба "Книги"
export class CBookTab {
    constructor() {
        this.view // объект для быстрого доступа к представлениям
        this.window // экземпляр окна для работы с книгами
        this.toEvents // функция перехода к событиям
        this.toEmployee // функция перехода к сотруднику
    }

    // метод инициализации компонента
    init() {
        this.window = new CBookWindow();
        this.window.onChange = () => { this.refreshTable }
        this.window.init()
    }

    // метод получения webix конфигурации компонента
    config() {
        webix.ui(this.window.config())
        webix.ui(BookTabContextMenu())

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
        this.view.datatableContextMenu.attachEvent('onMenuItemClick', (id) => {
            let item = this.view.datatableContextMenu.getItem(id).value
            let selected = this.view.datatable.getSelectedItem()

            switch (item) {
                case BOOK_CONTEXT_MENU.toEvent:
                    this.toEvents()
                    break;
                case BOOK_CONTEXT_MENU.toEmployee:
                    this.toEmployee()
                    break;
                case BOOK_CONTEXT_MENU.add:
                    this.window.parse(new Book())
                    this.window.switch(BOOK_WINDOW_TYPE.create)
                    break;
                case BOOK_CONTEXT_MENU.edit:
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
                case BOOK_CONTEXT_MENU.remove:
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
        });
    }

    // функция обновления таблицы книг
    refreshTable() {
        bookModel.getBooks().then((books) => {
            this.view.datatable.clearAll()
            this.view.datatable.define('data', books)
            this.view.datatable.refresh()
        })
    }
}

// допустимые значения пунктов контекстного меню таба Книги
export const BOOK_CONTEXT_MENU = {
    toEvent: 'К событию',
    toEmployee: 'К сотруднику',
    add: 'Добавить',
    edit: 'Изменить',
    remove: 'Удалить'
}
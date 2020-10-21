import { BookTabView, BookTabContextMenu, TabControllsView } from './BookTabView.js'
import { CBookWindow, BOOK_WINDOW_TYPE } from './bookWindow/CBookWindow.js'
import bookModel from '../../models/bookModel.js'
import eventModel from '../../models/eventModel.js'
import employeeModel from '../../models/employeeModel.js'
import { Book, BOOK_STATUS } from '../../models/entities/book.js'

// класс таба 'Книги'
export class CBookTab {
    constructor() {
        this.refreshControlls       // функция обновления элементов управления в header'е
        this.view                   // объект для быстрого доступа к представлениям
        this.window                 // экземпляр окна для работы с книгами
        this.updateEventsDatatable  // функция обновления таблицы событий
        this.names                  // массив сотрудников в сабменю
    }

    // метод инициализации компонента
    init(updateEventsDatatable, refreshControlls) {
        this.updateEventsDatatable = updateEventsDatatable  // функция обновления таблицы событий
        this.refreshControlls = refreshControlls            // функция обновления элементов управления в header'е

        this.window = new CBookWindow() // инициализация компонента окна
        this.window.init(
            () => { this.refreshTable() }
        ) // вызова инициализации компонента окна

        this.names = [] // массив сотрудников в сабменю
    }

    // метод получения webix конфигурации компонента
    config() {
        // т.к. window и popup расположены не в дереве приложения, а поверх слоев, его нужно отрисовывать отдельно
        webix.ui(this.window.config())
        webix.ui(BookTabContextMenu(this.names))

        // вызов функции представления
        return BookTabView()
    }

    // метод получения webix конфигурации элементов управления таба
    configTabControlls() {
        return TabControllsView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {
        // инициализация используемых представлений
        this.view = {
            datatable: $$('bookTabDatatable'),
            datatableContextMenu: $$('bookTabDatatableContextMenu'),
            controlls: $$('booktab-controlls'),
            btns: {
                createBtn: $$('booktab-add-btn'),
                updateBtn: $$('booktab-edit-btn'),
                deleteBtn: $$('booktab-remove-btn'),
            }
        }

        // создание книги
        this.view.btns.createBtn.attachEvent('onItemClick', () => {
            this.createBook()
        })

        // изменение книги
        this.view.btns.updateBtn.attachEvent('onItemClick', () => {
            this.updateBook()
        })

        // удаление книги
        this.view.btns.deleteBtn.attachEvent('onItemClick', () => {
            this.deleteBook()
        })

        // отложенное заполнение массива сотрудников в сабменю
        employeeModel.getEmployees().then((employees) => {
            // проверка наличия данных
            if (!employees) {
                return
            }

            employees.map((employee) => {
                this.names.push({ ID: employee.ID, value: `${employee.lastname} ${employee.firstname}` })
            })
        })

        // инициализация обработчиков событий модального окна
        this.window.attachEvents()

        // прикрепление контекстного меню к таблице
        this.view.datatableContextMenu.attachTo(this.view.datatable)

        // загрузка первичных данных в таблицу
        this.refreshTable()

        // обработка события нажатия на пункт контекстного меню
        this.view.datatableContextMenu.attachEvent('onMenuItemClick', (itemID) => {
            // проверка вложенности выбранного пункта меню
            if (!this.view.datatableContextMenu.getItem(itemID)) {
                this.handleSubMenu(itemID)
            } else {
                // получение значения пункта, на которое произошло нажатие
                let item = this.view.datatableContextMenu.getItem(itemID).value
                this.handleContextMenu(item)
            }

        })
    }

    // обработка выбора в контекстном меню
    handleContextMenu(item) {
        switch (item) {
            case BOOK_CONTEXT_MENU.edit: // редактирование выделленой книгиbreak
                this.updateBook()
                break
            case BOOK_CONTEXT_MENU.remove: // удаление выделенной книгиbreak
                this.deleteBook()
                break
            case BOOK_CONTEXT_MENU.take: // добавление книги
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

                // проверка статуса книги
                if (book.status === BOOK_STATUS.available) {
                    webix.message('Книга не выдана')
                    return
                }

                eventModel.createTakeEvent(book.ID).then(() => {
                    this.refreshTable()
                })
                break
            default:
                console.error(`Неизвестное значение пункта меню: ${item}.`)
                break
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

        // проверка статуса книги
        if (book.status === BOOK_STATUS.notAvailable) {
            webix.message('Книга уже выдана')
            return
        }

        eventModel.createGiveEvent(book.ID, employee.ID).then(() => {
            this.refreshTable()
            this.updateEventsDatatable()
        })
    }

    // функция обновления таблицы книг
    refreshTable(books) {
        if (books) {
            this.view.datatable.clearAll()
            this.view.datatable.parse(books)
            return
        } else {
            bookModel.getBooks().then((books) => {
                // проверка наличия данных
                if (books) {
                    // преобразование даты издания
                    books.map((book) => {
                        book.year = new Date(book.year)
                    })
                }

                // заполнение таблицы окна данными книги
                this.view.datatable.clearAll()
                this.view.datatable.parse(books)
            })
        }
    }

    // метод отображения таба с фильтрацией по книге
    showByBookID(bookID) {
        bookModel.getBookByID(bookID).then((book) => {
            // проверка наличия данных
            if (!book) {
                return
            }

            // применение фильтров
            this.view.datatable.getFilter('ISBN').value = book.ISBN
            this.view.datatable.filterByAll()

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

    // функция переключения оторбажения элементов управления таба
    switchControlls() {
        switch (this.view.controlls.isVisible()) {
            case true:
                this.hideControlls()
                break;
            case false:
                this.showControlls()
                break;
        }
    }

    // функция отображения элементов управления таба
    showControlls() {
        this.view.controlls.show()
    }

    // функция сокрытия элементов управления таба
    hideControlls() {
        this.view.controlls.hide()
    }

    // функция создания книги
    createBook() {
        this.window.parse(new Book())
        this.window.switch(BOOK_WINDOW_TYPE.create)
    }

    // функция изменения книги
    updateBook() {
        // получение выделенного элемента
        let selected = this.view.datatable.getSelectedItem()

        // проверка выделенного элемента
        if (!selected) {
            webix.message('Выделите строку')
            return
        }
        // проверка наличия поля ID у выделенного элемента
        if (!selected.ID) {
            console.error('Incorrect ID of item:', selected.ID)
            return
        }
        bookModel.getBookByID(selected.ID).then((book) => {
            // проверка наличия данных
            if (!book) {
                return
            }

            // преобразование даты издания
            let time = new Date(book.year)
            book.year = time.getFullYear()

            // заполнение полей окна данными книги
            this.window.parse(book)
            this.window.switch(BOOK_WINDOW_TYPE.update)
        })
    }

    // функция удаления книги
    deleteBook() {
        // получение выделенного элемента
        let selected = this.view.datatable.getSelectedItem()

        if (!selected) {
            webix.message('Выделите строку')
            return
        }
        if (!selected.ID) {
            console.error('id of item is ', selected.ID)
            return
        }
        bookModel.getBookByID(selected.ID).then((book) => {
            // проверка наличия данных
            if (!book) {
                return
            }
            // проверка выданности книги
            if (book.status === BOOK_STATUS.notAvailable) {
                webix.message('Нельзя удалить выданную книгу')
                return
            }

            // преобразование даты издания
            let time = new Date(book.year)
            book.year = time.getFullYear()

            // заполнение полей окна данными книги
            this.window.parse(book)
            this.window.switch(BOOK_WINDOW_TYPE.delete)
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
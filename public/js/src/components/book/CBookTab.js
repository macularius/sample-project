import { BookTabView, BookTabContextMenu } from "./BookTabView.js"
import { CBookWindow, BOOK_WINDOW_TYPE } from "./bookWindow/CBookWindow.js";

// класс таба "Книги"
export class CBookTab {
    constructor() {
        this.view // объект для быстрого доступа к представлениям
        this.window // экземпляр окна для работы с книгами
    }

    // метод инициализации компонента
    init() {
        this.window = new CBookWindow();
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

        // обработка события нажатия на пункт контекстного меню
        this.view.datatableContextMenu.attachEvent('onMenuItemClick', (id) => {
            let item = this.view.datatableContextMenu.getItem(id).value
            switch (item) {
                case BOOK_CONTEXT_MENU.toEvent:
                    console.log(BOOK_CONTEXT_MENU.toEvent)
                    break;
                case BOOK_CONTEXT_MENU.toEmployee:
                    console.log(BOOK_CONTEXT_MENU.toEmployee)
                    break;
                case BOOK_CONTEXT_MENU.add:
                    if (!this.view.datatable.getSelectedItem()) {
                        webix.message('Выделите строку')
                        return
                    }
                    this.window.switch(BOOK_WINDOW_TYPE.create)
                    break;
                case BOOK_CONTEXT_MENU.edit:
                    if (!this.view.datatable.getSelectedItem()) {
                        webix.message('Выделите строку')
                        return
                    }
                    this.window.switch(BOOK_WINDOW_TYPE.update)
                    break;
                case BOOK_CONTEXT_MENU.remove:
                    if (!this.view.datatable.getSelectedItem()) {
                        webix.message('Выделите строку')
                        return
                    }
                    this.window.switch(BOOK_WINDOW_TYPE.delete)
                    break;
                default:
                    console.error(`Неизвестное значение пункта меню: ${item}.`);
                    break;
            }
        });
    }
    
    // метод вызова модального окна
    showByID(id) {}
}

// допустимые значения пунктов контекстного меню таба Книги
export const BOOK_CONTEXT_MENU = {
    toEvent: 'К событию',
    toEmployee: 'К сотруднику',
    add: 'Добавить',
    edit: 'Изменить',
    remove: 'Удалить'
}
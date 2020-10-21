import { JournalTabView, EventTabContextMenu, TabControllsView } from './JournalTabView.js'
import eventModel from '../../models/eventModel.js'
import { GetDate } from '../../../helpers/dateFormatter.js'

// класс таба 'Журнал событий'
export class CJournalTab {
    constructor() {
        this.refreshControlls   // функция обновления элементов управления в header'е
        this.view               // объект для быстрого доступа к представлениям
        this.toBookTab          // функция перехода к книге
        this.toEmployeeTab      // функция перехода к сотруднику
    }

    // метод инициализации компонента
    init(toBookTab, toEmployeeTab, refreshControlls) {
        this.refreshControlls = refreshControlls    // функция обновления элементов управления в header'е
        this.toBookTab = toBookTab                  // функция перехода к книге, возвращает ссылку на CBookTab
        this.toEmployeeTab = toEmployeeTab          // функция перехода к сотруднику, возвращает ссылку на CEmployeeTab
    }

    // метод получения webix конфигурации компонента
    config() {
        // т.к. window и popup расположены не в дереве приложения, а поверх слоев, его нужно отрисовывать отдельно
        webix.ui(EventTabContextMenu())

        // вызов функции представления
        return JournalTabView()
    }

    // метод получения webix конфигурации элементов управления таба
    configTabControlls() {
        return TabControllsView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {
        // инициализация используемых представлений
        this.view = {
            datatable: $$('eventTabDatatable'),
            datatableContextMenu: $$('eventTabDatatableContextMenu'),
            controlls: $$('journaltab-controlls'),
            btns: {
                toEmployeeBtn: $$('journaltab-to-employee-btn'),
                toBookBtn: $$('journaltab-to-book-btn'),
            }
        }

        // переход к сотруднику
        this.view.btns.toEmployeeBtn.attachEvent('onItemClick', () => {
            this.toEmployee()
        })

        // переход к книге
        this.view.btns.toBookBtn.attachEvent('onItemClick', () => {
            this.toBook()
        })

        // прикрепление контекстного меню к таблице
        this.view.datatableContextMenu.attachTo(this.view.datatable)

        // загрузка первичных данных в таблицу
        this.refreshTable()

        // обработка события нажатия на пункт контекстного меню
        this.view.datatableContextMenu.attachEvent('onMenuItemClick', (id) => {
            // получение значения пункта, на которое произошло нажатие
            let item = this.view.datatableContextMenu.getItem(id).value
            this.handleContextMenu(item)
        });
    }

    // обработка выбора в контекстном меню
    handleContextMenu(item) {
        switch (item) {
            case EVENT_CONTEXT_MENU.toBook: // переход к книге
                this.toBook()
                break;
            case EVENT_CONTEXT_MENU.toEmployee: // переход к сотруднику
                this.toEmployee()
                break;
            default:
                console.error(`Неизвестное значение пункта меню: ${item}.`);
                break;
        }
    }

    // функция обновления таблицы событий
    refreshTable() {
        eventModel.getEvents().then((events) => {
            // проверка наличия данных
            if (!events) {
                return
            }

            // преобразование даты к строке
            events.map((event) => {
                if (!event) {
                    return
                }

                event.bookString = `[${event.book.ISBN}] ${event.book.name}`
                event.employeeString = `${event.employee.lastname} ${event.employee.firstname}`
                event.date = GetDate(new Date(event.date))
            })

            this.view.datatable.clearAll()
            this.view.datatable.define('data', events)
            this.view.datatable.refresh()
        })
    }

    // метод отображения таба с фильтрацией по книге и сотруднику
    showLastEventByBookID(bookID) {
        eventModel.getLastGiveEventByBookID(bookID).then((event) => {
            // проверка наличия данных
            if (!event) {
                return
            }

            // применение фильтров
            this.view.datatable.getFilter('bookString').value = `[${event.book.ISBN}] ${event.book.name}`
            this.view.datatable.getFilter('employeeString').value = `${event.employee.lastname} ${event.employee.firstname}`
            this.view.datatable.getFilter('type').value = event.type
            this.view.datatable.getFilter('date').value = GetDate(new Date(event.date))
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

    // функция перехода к книге
    toBook() {
        // получение выделенного элемента
        let selected = this.view.datatable.getSelectedItem()
        
        if (!selected) {
            webix.message('Выделите строку')
            return
        }
        if (!selected.ID) {
            console.error('Incorrect ID of item:', selected.ID)
            return
        }
        let cBookTab = this.toBookTab()
        cBookTab.showByBookID(selected.book.ID)
    }

    // функция перехода к сотруднику
    toEmployee() {
        // получение выделенного элемента
        let selected = this.view.datatable.getSelectedItem()

        if (!selected) {
            webix.message('Выделите строку')
            return
        }
        if (!selected.ID) {
            console.error('Incorrect ID of item:', selected.ID)
            return
        }

        let cEmployeeTab = this.toEmployeeTab()
        cEmployeeTab.showByEmployeeID(selected.employee.ID)
    }
}

// допустимые значения пунктов контекстного меню таба Книги
export const EVENT_CONTEXT_MENU = {
    toBook: 'К книге',
    toEmployee: 'К сотруднику'
}
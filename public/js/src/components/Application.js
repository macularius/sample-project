import { CUserInfo } from './userInfo/CUserInfo.js';
import { CBookTab } from './book/CBookTab.js';
import { CEmployeeTab } from './employee/CEmployeeTab.js';
import { CJournalTab } from './journal/CJournalTab.js';

// галвный компонент приложения
export class Application {
    constructor() {
        this.view // быстрый доступ к объектам представлений
        this.userInfo = new CUserInfo(); // экземпляр контроллера пользовательской информации 
        this.bookTab = new CBookTab(); // экземпляр контроллера книг
        this.employeeTab = new CEmployeeTab(); // экземпляр контроллера сотрудников
        this.journalTab = new CJournalTab(); // экземпляр контроллера событий
    }

    // метод инициализации главного компонента
    init() {
        this.userInfo.init()
        this.bookTab.init(
            () => { return this.journalTab.refreshTable() }, // updateEvent
        )
        this.employeeTab.init(
            () => { return this.dispatch(APP_TAB.booksTab) }, // toBook
            () => { return this.dispatch(APP_TAB.journalTab) }, // toEvent
            () => { return this.journalTab.refreshTable() }, // updateEventsDatatable
            () => { return this.bookTab.refreshTable() }, // updateBooksDatatable
        )
        this.journalTab.init(
            () => { return this.dispatch(APP_TAB.booksTab) }, // toBook
            () => { return this.dispatch(APP_TAB.employeesTab) }, // toEvent
        )
    }

    // метод вызова обработки событий
    attachEvents() {
        this.view = {
            tabbar: $$('main-tabbar'),
            multiviews: $$('main-views'),
        }

        this.dispatch(APP_TAB.booksTab)

        this.userInfo.attachEvents()
        this.bookTab.attachEvents()
        this.employeeTab.attachEvents()
        this.journalTab.attachEvents()
    }

    // метод отрисовки главной конфигурации представления
    config() {
        return {
            rows: [
                // header
                {
                    cols: [
                        {
                            view: 'tabbar',
                            id: 'main-tabbar',
                            value: 'listView',
                            width: 700,
                            multiview: true,
                            options: [
                                { id: 'bookTab', value: 'Книги' },
                                { id: 'employeeTab', value: 'Сотрудники' },
                                { id: 'journalTab', value: 'Журнал событий' },
                            ]

                        },
                        {},
                        this.userInfo.config(),
                    ],
                },
                // содержимое табов
                {
                    view: "multiview",
                    id: "main-views",
                    cells: [
                        this.bookTab.config(),
                        this.employeeTab.config(),
                        this.journalTab.config(),
                    ]
                },
            ],
        }
    }

    // метод диспетчеризации переключения между табами
    dispatch(tab) {
        let tabObj

        switch (tab) {
            case APP_TAB.booksTab:
                tabObj = this.bookTab
                break;
            case APP_TAB.employeesTab:
                tabObj = this.employeeTab
                break;
            case APP_TAB.journalTab:
                tabObj = this.journalTab
                break;
            default:
                console.error('Incorrect tab: ', tab);                
                return;
        }
        
        this.view.tabbar.setValue(tab)
        this.view.multiviews.setValue(tab)

        return tabObj
    }
}

// константы перечисления табов(id представления)
export const APP_TAB = {
    booksTab: 'bookTab',
    employeesTab: 'employeeTab',
    journalTab: 'journalTab',
}
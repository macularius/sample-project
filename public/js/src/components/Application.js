import { CUserInfo } from './userInfo/CUserInfo.js'
import { CBookTab } from './book/CBookTab.js'
import { CEmployeeTab } from './employee/CEmployeeTab.js'
import { CJournalTab } from './journal/CJournalTab.js'
import WorkedPlaceView from './ApplicationView.js'
import { CMainWindow } from './mainWindow/CMainWindow.js'
import { deleteCookie } from '../../helpers/cookies.js'
import { checkAuth } from '../../helpers/checkAuth.js'

// галвный компонент приложения
export class Application {
    constructor() {
        this.view                               // быстрый доступ к объектам представлений
        this.userInfo = new CUserInfo()         // экземпляр контроллера пользовательской информации 
        this.bookTab = new CBookTab()           // экземпляр контроллера книг
        this.employeeTab = new CEmployeeTab()   // экземпляр контроллера сотрудников
        this.journalTab = new CJournalTab()     // экземпляр контроллера событий
        this.mainWindow = new CMainWindow()     // окно входа в приложение
    }

    // метод инициализации главного компонента
    init() {
        // инициализация компонента информации о пользователе
        this.userInfo.init(
            () => {
                deleteCookie('auth-token')
                location.replace('/user/logout')
            }, // onLogout
        )
        // инициализация компонента вкладки книг
        this.bookTab.init(
            () => { return this.journalTab.refreshTable() }, // updateEvent
        )
        // инициализация компонента вкладки сотрудников
        this.employeeTab.init(
            () => { return this.dispatch(APP_TAB.booksTab) }, // toBook
            () => { return this.dispatch(APP_TAB.journalTab) }, // toEvent
            () => { return this.journalTab.refreshTable() }, // updateEventsDatatable
            () => { return this.bookTab.refreshTable() }, // updateBooksDatatable
        )
        // инициализация компонента вкладки событий
        this.journalTab.init(
            () => { return this.dispatch(APP_TAB.booksTab) }, // toBook
            () => { return this.dispatch(APP_TAB.employeesTab) }, // toEvent
        )
        // инициализация компонента окна входа в приложение
        this.mainWindow.init(
            () => { location.replace('/') }, // onLogin
        )
    }

    // метод вызова обработки событий
    attachEvents() {
        this.view = {
            tabbar: $$('main-tabbar'),
            multiviews: $$('main-views'),
            workedPlace: $$('workedPlace'),
        }

        // компоненты требующие авторизации
        // вызываются через проверку авторизации
        // если клиент не авторизован, то эти
        // компоненты не будут отрисованы
        checkAuth((isAuth) => {
            if (isAuth) {
                this.view.workedPlace.show()

                this.dispatch(APP_TAB.booksTab)

                this.userInfo.attachEvents()
                this.bookTab.attachEvents()
                this.employeeTab.attachEvents()
                this.journalTab.attachEvents()
            } else {
                this.view.workedPlace.hide()
            }
        })

        // вызов обработки событий окна входа в приложение
        this.mainWindow.attachEvents()

        // первоночальное состояние приложения
        this.view.workedPlace.hide()
        this.mainWindow.switch()
    }

    // метод отрисовки главной конфигурации представления
    config() {
        webix.ui(this.mainWindow.config())

        return WorkedPlaceView(this.bookTab, this.employeeTab, this.journalTab, this.userInfo)
    }

    // метод диспетчеризации переключения между табами
    dispatch(tab) {
        let tabObj

        switch (tab) {
            case APP_TAB.booksTab:
                tabObj = this.bookTab
                break
            case APP_TAB.employeesTab:
                tabObj = this.employeeTab
                break
            case APP_TAB.journalTab:
                tabObj = this.journalTab
                break
            default:
                console.error('Incorrect tab: ', tab)
                return
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
import { CToolbar } from './toolbar/CToolbar.js'
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
        this.userInfo = new CToolbar()         // экземпляр контроллера пользовательской информации 
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
            (config) => { this.refreshControlls(config) }, // refreshControlls
        )
        // инициализация компонента вкладки сотрудников
        this.employeeTab.init(
            () => { return this.dispatch(APP_TAB.booksTab) }, // toBook
            () => { return this.dispatch(APP_TAB.journalTab) }, // toEvent
            () => { return this.journalTab.refreshTable() }, // updateEventsDatatable
            () => { return this.bookTab.refreshTable() }, // updateBooksDatatable
            (config) => { this.refreshControlls(config) }, // refreshControlls
        )
        // инициализация компонента вкладки событий
        this.journalTab.init(
            () => { return this.dispatch(APP_TAB.booksTab) }, // toBook
            () => { return this.dispatch(APP_TAB.employeesTab) }, // toEvent
            (config) => { this.refreshControlls(config) }, // refreshControlls
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
            tabControllsContainer: $$('main'),
        }

        // компоненты требующие авторизации
        // вызываются через проверку авторизации
        // если клиент не авторизован, то эти
        // компоненты не будут отрисованы
        checkAuth((isAuth) => {
            if (isAuth) {
                // переключение таба
                this.view.tabbar.attachEvent('onItemClick', () => {
                    this.dispatch(this.view.tabbar.getValue())
                })

                // отрисовать рабочее пространство
                this.view.workedPlace.show()

                // обработчики событий компонентов
                this.userInfo.attachEvents()
                this.bookTab.attachEvents()
                this.employeeTab.attachEvents()
                this.journalTab.attachEvents()

                // выделить таб книг
                this.dispatch(APP_TAB.booksTab)
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

        // определение объекта таба
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

        // переключение таба
        this.view.tabbar.setValue(tab)
        this.view.multiviews.setValue(tab)

        // замена элементов управления в header'е
        this.bookTab.hideControlls()
        this.employeeTab.hideControlls()
        this.journalTab.hideControlls()
        tabObj.switchControlls()

        return tabObj
    }

    refreshControlls(config) {
        webix.ui(config, this.view.tabControllsContainer, $$('tab-controlls'))
    }
}

// константы перечисления табов(id представления)
export const APP_TAB = {
    booksTab: 'bookTab',
    employeesTab: 'employeeTab',
    journalTab: 'journalTab',
}
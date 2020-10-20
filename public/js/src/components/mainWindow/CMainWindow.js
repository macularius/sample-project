import authModel from '../../models/authModel.js'
import MainWindowView from './MainWindowView.js'
import { User } from './../../models/entities/user.js'
import { checkAuth } from '../../../helpers/checkAuth.js'

// компонент окна авторизации
export class CMainWindow {
    constructor() {
        this.view       // объект для быстрого доступа к представлениям
        this.onLogin    // callback функция авторизации пользователя
    }

    // метод инициализации компонента
    init(onLogin, onLogout) {
        this.onLogin = onLogin
    }

    // метод получения webix конфигурации компонента
    config() {
        return MainWindowView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {
        // инициализация используемых представлений
        this.view = {
            window: $$('mainWindow'),
            form: {
                form: $$('mainWindowForm'),
                login: $$('mainWindowFormLogin'),
                password: $$('mainWindowFormPassword'),
                confirm: $$('mainWindowConfirmBtn'),
            }
        }

        // событие входа в приложение
        this.view.form.confirm.attachEvent('onItemClick', () => {
            let user = new User()
            user.login = this.view.form.login.getValue()
            user.password = this.view.form.password.getValue()

            // авторизация пользователя
            authModel.login(user).then((isValid) => {
                if (isValid === true) {
                    this.switch()
                    this.onLogin()
                }
            })
        })
    }

    // метод переключения состояния окна входа в приложение
    // результат выполнения функции зависит от авторизованности пользователя
    switch() {
        checkAuth((isAuthorize) => {
            if (isAuthorize) {
                this.hide()
            } else {
                this.show()
            }
        })
    }

    // метод отображения окна
    show() {
        this.view.window.show()
    }

    // метод сокрытия окна
    hide() {
        this.view.window.hide()
    }

    // функция валидации формы
    validate() {
        let isValid = false

        // удаление пробелов в полях формы

        // валидация webix

        return isValid
    }
}

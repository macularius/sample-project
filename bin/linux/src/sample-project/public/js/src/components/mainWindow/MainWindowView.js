// возвращает webix конфигурацию окна авторизации
export default function MainWindowView() {
    return {
        view: 'window',
        id: 'mainWindow',
        fullscreen: true,
        modal: true,
        headHeight: 0,
        position: 'center',
        css: 'main_window',
        body: {
            rows: [
                {},
                {
                    cols: [
                        {},
                        {
                            view: 'form',
                            id: 'mainWindowForm',
                            elements: [
                                {
                                    view: 'text',
                                    id: 'mainWindowFormLogin',
                                    label: 'Логин',
                                    name: 'login',
                                    required: true,
                                    labelWidth: 150,
                                },
                                {
                                    view: 'text',
                                    id: 'mainWindowFormPassword',
                                    label: 'Пароль',
                                    type: 'password',
                                    name: 'password',
                                    required: true,
                                    labelWidth: 150,
                                },
                                {
                                    view: 'button',
                                    id: 'mainWindowConfirmBtn',
                                    value: 'Войти',
                                },
                            ]
                        },
                        {},
                    ]
                },
                {},
            ]
        }
    }
}
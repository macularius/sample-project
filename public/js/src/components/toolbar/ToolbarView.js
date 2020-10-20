// возвращает webix конфигурации таба для работы с событиями
export default function ToolbarView() {
    return {
        view: 'toolbar',
        cols: [
            {
                view: 'label',
                label: 'Электронная библиотека',
            },
            {},
            // фио
            {
                view: 'label',
                id: 'userInfoLabel',
                label: 'загрузка...',
                width: 200,
            },
            // кнопка выхода
            {
                view: 'button',
                id: 'logoutBtn',
                css: 'webix_secondary',
                label: 'Выход',
                width: 150,
            },
        ]
    }
}
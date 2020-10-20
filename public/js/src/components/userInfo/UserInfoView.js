// возвращает webix конфигурации таба для работы с событиями
export default function UserInfoView() {
    return {
        cols: [
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
                label: 'Выход',
                width: 150,
            },
        ]
    }
}
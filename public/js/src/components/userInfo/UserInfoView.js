// возвращает webix конфигурации таба для работы с событиями
export default function UserInfoView(label) {
    return {
        cols: [
            // фио
            {
                view: 'label',
                id: 'userInfoLabel',
                label: label,
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
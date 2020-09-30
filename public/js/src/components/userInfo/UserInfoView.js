// возвращает webix конфигурации таба для работы с событиями
export default function UserInfoView() {
    return {
        cols: [
            // фио
            {
                template: 'Фамилия Имя Отчество',
                width: 200,
            },
            // кнопка выхода
            {
                view: 'button',
                label: 'Выход',
                width: 150,
            },
        ]
    }
}
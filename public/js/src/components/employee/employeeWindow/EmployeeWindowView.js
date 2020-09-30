// возвращает webix конфигурации окна для работы с сущностью сотрудника
export default function EmployeeWindowView() {
    let headText = 'Сотрудник'

    return { 
        view: 'window',
        id: 'employeeWindow',
        head: headText,
        modal: true,
        body: {
            view: 'form',
            id: 'bookWindowForm',
            elements: {
                rows: [
                    {
                        view: 'text',
                        label: 'ISBN',
                    },
                    {
                        view: 'text',
                        label: 'Название',
                    },
                    {
                        view: 'text',
                        label: 'Автор',
                    },
                    {
                        view: 'text',
                        label: 'Издательство',
                    },
                    {
                        view: 'text',
                        label: 'Статус',
                    },
                    {
                        view: 'text',
                        label: 'Год издания',
                    },
                ],
            }
        }
    }
}
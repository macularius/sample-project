// возвращает webix конфигурации таба для работы с сотрудникамиы
export default function EmployeeTabView(windowConfig) {
    return {
        id: 'employeeTab',
        rows: [
            {
                view: 'datatable',
                columns: [
                    { id: 'lastname', header: 'Фамилия', fillspace: true, },
                    { id: 'firstname', header: 'Имя', fillspace: true, },
                    { id: 'middlename', header: 'Отчество', fillspace: true, },
                    { id: 'position', header: 'Должность', fillspace: true, },
                    { id: 'phoneNumber', header: 'Телефонный номер', fillspace: true, },
                    { id: 'email', header: 'Email', fillspace: true, },
                ],
                data: []
            },
        ]
    }
}
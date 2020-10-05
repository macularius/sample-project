import { BOOK_WINDOW_TYPE } from './CBookWindow.js';

// возвращает webix конфигурацию окна для работы с сущностью книги
export default function BookWindowView(type) {
    let headText = 'Книга' // текст заголовка модального окна

    switch (type) {
        case BOOK_WINDOW_TYPE.create:
            headText = 'Добавление книги'
            break;
        case BOOK_WINDOW_TYPE.update:
            headText = 'Редактирование книги'
            break;
        case BOOK_WINDOW_TYPE.delete:
            headText = 'Удаление книги'
            break;
    }

    return {
        view: 'window',
        id: 'bookWindow',
        head: {
            view: 'template',
            id: 'bookWindowLabel',
            template: headText,
            css: 'webix_template'
        },
        modal: true,
        position: "center",
        width: 400,
        body: {
            view: 'form',
            id: 'bookWindowForm',
            elements: [
                {
                    view: 'text',
                    label: 'ISBN',
                    name: 'ISBN',
                    required: true,
                    labelWidth: 150,
                },
                {
                    view: 'text',
                    label: 'Название',
                    name: 'name',
                    required: true,
                    labelWidth: 150,
                },
                {
                    view: 'text',
                    label: 'Автор',
                    name: 'author',
                    required: true,
                    labelWidth: 150,
                },
                {
                    view: 'text',
                    label: 'Издательство',
                    name: 'publisher',
                    required: true,
                    labelWidth: 150,
                },
                {
                    view: 'text',
                    label: 'Статус',
                    name: 'status',
                    required: true,
                    labelWidth: 150,
                },
                {
                    view: 'text',
                    label: 'Год издания',
                    name: 'year',
                    required: true,
                    labelWidth: 150,
                },
                {
                    cols: [
                        {
                            view: 'button',
                            id: 'bookWindowConfirmBtn',
                            value: 'Применить',
                        },
                        {
                            view: 'button',
                            id: 'bookWindowCancelBtn',
                            value: 'Отмена',
                        },
                    ]
                },
            ]
        }
    }
}
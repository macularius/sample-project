import { BOOK_WINDOW_TYPE } from './CBookWindow.js';

// возвращает webix конфигурации окна для работы с сущностью книги
export default function BookWindowView(type) {
    let headText = 'Книга'

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
        position:"center",
        body: {
            view: 'form',
            id: 'bookWindowForm',
            elements: [
                {
                    view: 'text',
                    label: 'ISBN',
                    name: 'isbn',
                },
                {
                    view: 'text',
                    label: 'Название',
                    name: 'name',
                },
                {
                    view: 'text',
                    label: 'Автор',
                    name: 'author',
                },
                {
                    view: 'text',
                    label: 'Издательство',
                    name: 'publisher',
                },
                {
                    view: 'text',
                    label: 'Статус',
                    name: 'status',
                },
                {
                    view: 'text',
                    label: 'Год издания',
                    name: 'year',
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
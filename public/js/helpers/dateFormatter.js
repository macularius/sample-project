// функция форматирования даты в человеко-читаемую строку
// date: Date
export function FormatDate(date) {
    var format = webix.Date.dateToStr("%Y-%m-%d %H:%i")
    
    return format(new Date(date.getTime() + date.getTimezoneOffset() * 60000))
}

// функция форматирования даты с учетом часового пояса
// date: Date
export function GetDate(date) {    
    return new Date(date.getTime() + date.getTimezoneOffset() * 60000)
}
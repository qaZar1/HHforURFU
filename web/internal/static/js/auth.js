// auth.js
function isTokenExpired() {
    // Получаем куки
    const cookies = document.cookie.split('; ');
    
    // Ищем куки с именем 'token'
    const tokenCookie = cookies.find(cookie => cookie.startsWith('token='));

    // Проверяем, существует ли куки
    if (!tokenCookie) {
        return true; // Токен не существует, считается, что истек
    }

    // Получаем куки с временем истечения
    const expiresCookie = cookies.find(cookie => cookie.startsWith('token_expires='));

    // Проверяем, существует ли куки с временем истечения
    if (!expiresCookie) {
        return true; // Время истечения не задано, считаем, что истекло
    }

    // Получаем дату истечения
    const expiresAt = new Date(expiresCookie.split('=')[1]).getTime();

    // Проверяем, истек ли токен
    return Date.now() > expiresAt; // Если текущее время больше даты истечения, токен истек
}


async function checkToken() {
    if (isTokenExpired()) {
        window.location.href = '/login';
    }
}

// Экспортируем функцию для использования на других страницах
export { checkToken };

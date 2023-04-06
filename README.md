## Micorservices labs (Команда 9)
- ***Вавринюк Максим*** - movie-service
- ***Тітов Єгор*** - admin-service
- ***Юдаков Олександр*** - user-service

***Тема:*** Система рекомендацій фільмів.
- Сервіс пошуку даних про фільми (movie-service): Цей сервіс відповідатиме за отримання даних про фільми з різних джерел, наприклад IMDB.
- Сервіс рекомендацій (user-service): Цей сервіс відповідатиме за аналіз історії переглядів та вподобань користувача і генерувати рекомендації фільмів на основі його інтересів.
- Сервіс адміністрування (admin-service): Цей сервіс відповідатиме за адміністрування юзерів та перегляд статистики.
- Клієнт.


apply command:

```kubectl apply -f k8s```
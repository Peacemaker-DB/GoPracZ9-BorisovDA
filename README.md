# Практическое задание № 9 Борисов Денис Александрович ЭФМО-01-25

Цели:

1.	Научиться безопасно хранить пароли (bcrypt), валидировать вход и обрабатывать ошибки.
2.	Реализовать эндпоинты POST /auth/register и POST /auth/login.
3.	Закрепить работу с БД (PostgreSQL + GORM или database/sql) и валидацией ввода.
4.	Подготовить основу для JWT-аутентификации в следующем ПЗ (№10).

# Выполнение практического задания.

1. Структура проекта

  Для выполнения практической работы была сделана следующая структура проекта

<img width="295" height="337" alt="image" src="https://github.com/user-attachments/assets/7da6532b-11b7-4e9f-a94f-ee119684305b" />

   Так же были установлены все необходимые расширения для выполнения практики

<img width="636" height="428" alt="Снимок экрана 2025-11-16 211228" src="https://github.com/user-attachments/assets/b5cd150a-d1da-4a26-a1b4-1160c0c4f303" />

2.	Конфигурация и подключение к БД (PostgreSQL).
  Для реализации подключения к PostgreSQL был создан файл config.go

<img width="357" height="467" alt="Снимок экрана 2025-11-16 212403" src="https://github.com/user-attachments/assets/cca676ab-bfc8-4bcc-a21b-a28726a29d1e" />

  После был создан файл postgres.go

  <img width="428" height="226" alt="Снимок экрана 2025-11-16 212513" src="https://github.com/user-attachments/assets/2aa4f5a2-3bfe-493c-b990-738c614bbe23" />

3. Модель пользователя и миграция.
 Для выполнения практики была написана модель пользователя в файле user.go, а так же миграция данных в файле user_repo.go. В котором реализован CRUD для пользователя.

   Файл user.go

<img width="565" height="242" alt="Снимок экрана 2025-11-16 212624" src="https://github.com/user-attachments/assets/a8a4bcb2-3853-489e-976a-3d8b6f25e341" />


   Файл user_repo.go

<img width="602" height="685" alt="Снимок экрана 2025-11-16 212715" src="https://github.com/user-attachments/assets/e0dd24d3-bb14-4447-9c97-89f2fe0adf67" />


5. HTTP-слой: валидация, bcrypt-хэширование и логин.

   Затем был создан файл auth.go, где происходит валидация, bcrypt-хэширование и логин

<img width="542" height="959" alt="Снимок экрана 2025-11-16 212938" src="https://github.com/user-attachments/assets/08ff344e-36c0-4b38-a3aa-2f1784e99f32" />

6. Точка входа и запуск сервера

   Для запуска сервера был написан main.go

<img width="461" height="535" alt="Снимок экрана 2025-11-16 213024" src="https://github.com/user-attachments/assets/8616d908-eb9f-4c38-8a2a-162b25ffeb28" />

# Проверка работоспособности

  Для проверки работоспособности был запущен сервер, посел в Postman были проверено:

Создание пользователя

<img width="686" height="432" alt="image" src="https://github.com/user-attachments/assets/014bdbe7-ab65-4041-9e90-17568eefe6ae" />

Повторная регистрация 

<img width="690" height="402" alt="image" src="https://github.com/user-attachments/assets/796347de-f751-4e41-a5b2-5825bbde6b01" />

Вход (верный)

<img width="687" height="382" alt="image" src="https://github.com/user-attachments/assets/16859b18-6935-43b0-924d-45ccc2fedf4d" />

Вход (неверный) → 401

<img width="687" height="337" alt="image" src="https://github.com/user-attachments/assets/ff7c39f3-ecfd-42a7-9ac1-4ca4deede77f" />

Данные пользователя в БД

<img width="1098" height="135" alt="image" src="https://github.com/user-attachments/assets/d9b4941a-4863-4db8-8211-18ceb7dd99d9" />


-- db service

-- Создание таблицы пользователей
CREATE TABLE users (
                       id UUID PRIMARY KEY
);

-- Вставка демонстрационных пользователей
INSERT INTO users (id) VALUES
                           ('81f4b2ae-4af4-4e15-8c85-7a324b6f0c58'),
                           ('c4be2cbc-bf65-4a82-860f-c3e8aa8a6769'),
                           ('492a3ce6-9998-45f2-9443-f8bbaca87360'),
                           ('b6702108-f752-47a9-b14e-15a7422341ab');

-- Создание таблицы подписок (сервисов)
CREATE TABLE services (
                          id UUID PRIMARY KEY,
                          name TEXT,
                          price INTEGER,
                          user_id UUID,
                          start_date TIMESTAMP,
                          end_date TIMESTAMP,
                          created_at TIMESTAMP
);

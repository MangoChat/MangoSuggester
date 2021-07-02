# MangoSuggester

Простейший предложка бот для Telegram. Сделан для каналов MangoChat Communities, там же и найдете применение. 

## Принцип работы

1. Пользователь присылает сообщение
2. Бот копирует его в личку админу. Следующим сообщением присылает ссылку на автора.

## Установка (Linux)

    git clone https://github.com/MangoChat/MangoSuggester
    nano settings.env 
   *В этом файле вы указываете Telegram token, Id админа, ваш URL (longpoll пока не поддерживается), приветственные сообщения
   

    docker build --tag mangosuggesterimage:latest .
    docker run -p 2010:2010 --name mangosuggester --restart always -d mangosuggesterimage
Порт только свой не забудьте подставить


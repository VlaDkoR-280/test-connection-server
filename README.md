
## Запуск приложения на сервере
>[!WARNING]
> Необходимо иметь установленный **docker**
```
docker run -d \
  --name health-monitor \
  --restart unless-stopped \
  -p {{YOUR_PORT}}:8080 \
  vladkor280/health-monitor:latest
```

>[!TIP] 
> `{{YOUR_PORT}}` - ваш порт, который будет открыт для работы сервера
>
> **Docker** открывает порт автоматически

### Для переустановки или обновления необходимо
- Остановить контейнер
```
docker stop health-monitor
```
- Удалить старый контейнер
```
docker rm health-monitor
```
- [Выполнить команду для запуска](#запуск-приложения-на-сервере)

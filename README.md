# GoGeoserver

Сервер для обработки geoTiff изображений и отдачи его по тайлам <br>
Для работы требует предустоновленную библиотеку GDAL на устройстве

## Настройка окружений
- Установите GDAL
```bash
sudo apt update
```

```bash 
sudo apt install gdal-bin libgdal-dev
```

## Развертывание
- Положите файл карты в resource/map/geo_map.tif
- запустите cmd/main.go

## Проверка работы
- Перейдите по адресу http://localhost:8080/
- карта должна была инициализироваться

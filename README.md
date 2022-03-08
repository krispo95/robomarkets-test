Разработать REST API для получения координат пользователя по его IP-адресу и получения списка местоположений по названию города.

Требование к архитектуре и исходному коду:
1. API должно быть спроектировано и разработано с расчетом на 100 000 000 запросов в день.
2. Исходный код должен быть оформлен в едином стиле и содержать необходимые комментарии.
3. Аккуратность исходного кода будет оцениваться наряду с функциональностью приложения.

Техническое описание приложения:

База данных хранится в файле geobase.dat, которые содержится в прикрепленном к письму архиве. База данных не будет изменятся она предназначена только для чтения.
База данных имеет бинарный формат. В файле последовательно хранятся:

60 байт — заголовок
```
int   version;           // версия база данных
sbyte name[32];          // название/префикс для базы данных
ulong timestamp;         // время создания базы данных
int   records;           // общее количество записей
uint  offset_ranges;     // смещение относительно начала файла до начала списка записей с геоинформацией
uint  offset_cities;     // смещение относительно начала файла до начала индекса с сортировкой по названию городов
uint  offset_locations;  // смещение относительно начала файла до начала списка записей о местоположении
```

12 байт * Header.records (количество записей) — cписок записей с информацией об интервалах IP адресов, отсортированный по полям ip_from и ip_to
```
uint ip_from;           // начало диапазона IP адресов
uint ip_to;             // конец диапазона IP адресов
uint location_index;    // индекс записи о местоположении
```

96 байт * Header.records (количество записей) — cписок записей с информацией о местоположении с координатами (долгота и широта)
```
sbyte country[8];        // название страны (случайная строка с префиксом "cou_")
sbyte region[12];        // название области (случайная строка с префиксом "reg_")
sbyte postal[12];        // почтовый индекс (случайная строка с префиксом "pos_")
sbyte city[24];          // название города (случайная строка с префиксом "cit_")
sbyte organization[32];  // название организации (случайная строка с префиксом "org_")
float latitude;          // широта
float longitude;         // долгота
```


4 байта * Header.records (количество записей) — список индексов записей местоположения отсортированный по названию города, каждый индекс — это адрес записи в файле относительно Header.offset_locations
База данных грузится полностью в память при старте приложения.
Необходимо реализовать быстрый (бинарный) поиск по загруженной базе по IP адресу и по точному совпадению названия города с учетом регистра.

В приложении должны быть реализованы два метода HTTP API:
```
GET /ip/location?ip=123.234.123.234
GET /city/locations?city=cit_Gbqw4
ответ сервера на каждый из запросов должен быть представлен в формате JSON.
```

Дополнительно:
1. Проявление инициативы сверх основного задания приветствуется.
2. Любые комментарии к коду приветствуются

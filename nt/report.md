# Отчет о выполнении нагрузочного тестирования

## Конфигруация сервера - 2 ГБ RAM 40 ГБ SSD 2 CPU

## До добавления индекса
### Запрос - ```select * from user where name like  'мих%' and surname like 'шев%'```
### Результат выполнения команды Explain:
<table border="1" style="border-collapse:collapse">
<tr><th>id</th><th>select_type</th><th>table</th><th>partitions</th><th>type</th><th>possible_keys</th><th>key</th><th>key_len</th><th>ref</th><th>rows</th><th>filtered</th><th>Extra</th></tr>
<tr><td>1</td><td>SIMPLE</td><td>user</td><td>NULL</td><td>ALL</td><td>NULL</td><td>NULL</td><td>NULL</td><td>NULL</td><td>950258</td><td>1.23</td><td>Using where</td></tr>
</table>

### Результаты нагрузочного тестирования с использованием wrk, значение timeout - 5s
| Threads | Connections | Latency  | Req/Sec | Total Request | Broken Request % |
|---------|-------------|----------|---------|---------------|------------------|
| 1       | 1           | 868.58ms | 1,26    | 38            | 0                |
| 10      | 10          | 3,2s     | 3       | 90            | 0                |
| 12      | 100         | 1,15s    | 150,09  | 4510          | 100%             |
| 12      | 1000        | 4,91s    | 193,09  | 5807          | 100%             |

#### При установлении значения connections >=100 все 100% запросов завершились с ошибками.

## После добавления индекса
### Запрос на добавление индекса ```create index user_surname_name on user (surname, name);```
По итогам тестов был выбран комбинированный индекс по колонкам surname и name, на первом месте используется колонка surname, т.к. этот вариант показал лучшее время выполнения запросов.
### Запрос - ```select * from user where name like  'мих%' and surname like 'шев%'```
### Результат выполнения команды Explain:
<table border="1" style="border-collapse:collapse">
    <tr><th>id</th><th>select_type</th><th>table</th><th>partitions</th><th>type</th><th>possible_keys</th><th>key</th><th>key_len</th><th>ref</th><th>rows</th><th>filtered</th><th>Extra</th></tr>
    <tr><td>1</td><td>SIMPLE</td><td>user</td><td>NULL</td><td>range</td><td>user_surname_name</td><td>user_surname_name</td><td>2044</td><td>NULL</td><td>1</td><td>11.11</td><td>Using index condition</td></tr>
</table>

### Результаты нагрузочного тестирования с использованием wrk, значение timeout - 5s

| Threads | Connections | Latency | Req/Sec | Total Request | Broken Request % |
|---------|-------------|---------|---------|---------------|------------------|
| 1       | 1           | 13.26ms | 136.53  | 4100          | 0                |
| 10      | 10          | 29.32ms | 716.19  | 21518         | 0                |
| 12      | 100         | 96.62ms | 775.98  | 23314         | 0.17%            |
| 12      | 1000        | 3.54s   | 1370.03 | 41143         | 4,16%             |

#### После добавления индекса скорость обработки запросов увеличилась многократно, например показатель Req/Sec вырос более чем в 100 раз.
#### При значении connections=100 начинают наблюдаться ошибки при выполнении запросов, но это значение составляет менее 1%.
#### При значении connections=1000 метрика latency выросла в 36 раз, а количество запросов с ошибками составило 4,16%.



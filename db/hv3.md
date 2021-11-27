# Отчет о выполнении домашнего задания "Полусинхронная репликация"
Выполнение ДЗ производилось на локальной машине с использованием docker-compose
## Часть №1 Асинхронная репликация

Стандартными средствами Mysql была настроена асинхронная репликация master и 2 slave  

Для разделения записи и чтения использовался ProxySQL.  

На реплики было переведено два запроса на чтение:
1. Поисковый запрос пользователей
2. Выбор интересов пользователя
```
mysql_query_rules:
(
    {
        rule_id=1
        active=1
        match_pattern="^select .* from user where id=."
        destination_hostgroup=0
        apply=1
    },
    {
        rule_id=2
        active=1
        match_pattern="^select .* from user"
        destination_hostgroup=1
        apply=1
    },
    {
        rule_id=3
        active=1
        match_pattern="^select .* from interest"
        destination_hostgroup=1
        apply=1
    },
)
```
Нагрузочное тестирование проводилось с помощью утилиты wrk и скрипта, подготовленного в рамках выполнения ДЗ по нагрузочному тестированию.

Т.к. мастер и слейвы запускались на локальной машине в docker-compose, то проверка, что запросы на чтения переключены на реплику, выполнялась через анализ general.log мастера и реплик.

## Часть №2 Настройка полусинхронной репликация
1. Включение row-based репликации:
   - на мастере и слейвах в my.cnf добавлено: ```binlog_format=ROW```
2.  Включение GTID:
   - на мастере и слейвах в my.cnf добавлено: ```gtid_mode=ON```
   - на мастере и слейвах my.cnf добавлено: ```enforce-gtid-consistency=true```
   - на слейвах выполнены команды ```STOP REPLICA; CHANGE MASTER TO MASTER_AUTO_POSITION = 1; START REPLICA```
3. Настройка полуcинхронной репликации:
   - на мастер установлен плагин ```rpl_semi_sync_source```
   - на мастере выполнена команда ```set global rpl_semi_sync_source_enabled = 1```
   - на мастере выполнена команда ```set global rpl_semi_sync_source_wait_for_replica_count = 1``` // ждет подтверждение от одной реплики
   - на слейвах установлен плагин ```rpl_semi_replica_source```
   - на слейвах выполнена команды ```STOP SLAVE IO_THREAD; set global rpl_semi_sync_replica_enabled = 1; START SLAVE IO_THREAD;```
   - проверка статуса полуcинхронной репликации на слейвах ```SHOW STATUS LIKE 'Rpl_semi_sync%'```
   - проверка статуса полуcинхронной репликации на мастере ```SHOW STATUS LIKE 'Rpl_semi_sync%'```

## Часть №3 Проверка потери транзакций
1. Нагрузка создана скриптом по созданию пользователей (см. front/faker/index.mjs)
2. Мастер выключен путем остановки контейнера: ```docker stop social-db -t 0```
3. Промоутим слейв db-replica-1 до мастера:
   - ```STOP SLAVE IO_THREAD;```
   - ```SHOW PROCESSLIST;``` // проверяем, что слейв закончил чтение relay лог
   - ```STOP REPLICA;``` // останавливаем реплику
   - ```RESET MASTER``` // очистка бин лога
4. На слейве db-replica-2:
   - ```STOP SLAVE;``` // останавливаем репликацию 
   - ```CHANGE MASTER TO MASTER_HOST='db-replica-1', master_auto_position = 1;``` // переключаем на master
   - ```START replica;``` // запускаем репликацию
   - ```SHOW REPLICA STATUS;``` // проверяем статус репликации
5. На ProxySQL:
   - ```UPDATE mysql_servers SET hostgroup_id = 3 WHERE hostname = 'db';``` // устанавливаем для master hostgroup_id, который не используется в правилах proxysql
   - ```UPDATE mysql_servers SET hostgroup_id = 0 WHERE hostname = 'db-replica-1';``` // устанавливаем для нового мастера hostgroup_id, для корректного разделения записи и чтения
   - ```LOAD MYSQL SERVERS TO RUNTIME;``` // применяем изменения
6. Проверка потери транзакций: 
   - запускаем контейнер с мастером, проверяем последних созданных пользователей и их общее количество
   - проверяем последних созданных пользователей на новом мастере и сравниваем их общее количество с данными на станом мастере
   - потери транзакций не обнаружено
   - снова даем нагрузку через скрипт создания пользователей и убеждаемся, что вся нагрузка идет на новый мастер и данные реплицируются в db-replica-2

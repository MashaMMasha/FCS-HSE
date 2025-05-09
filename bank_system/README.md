# КПО ИДЗ 1
## Задание
Перед вами поставлена задача разработать классы доменной модели ключевого модуля
будущего приложения – модуль «Учет финансов». Доменная модель классов должна быть
реализована с соблюдением принципов SOLID, ключевыми идеями GRASP: High Cohesion и
Low Coupling, а также рядом паттернов GoF: порождающих; структурных и поведенческих.

## Описание решения
Для решения задачи были созданы следующие классы:
1. `BankAccount` - класс банковского счета, содержит информацию о счете.
2. `Category` - класс пользовательских категорий расходов/доходов.
3. `Operation` - класс операции, содержит информацию о совершенной банковской операции.

### Паттерны

1. Для создания объектов классов `BankAccount`, `Category` и `Operation` был использован паттерн **фабрика + реестр**. Классы которые реализуют комбинацию этих паттернов инкаспулируют создание объекта (со скрытием некотрых вещей таких как генерация uuid) и хранение объектов (а так же управление их состоянием и жизненным циклом).
2. Для управления банковской системой применен паттерн **фасад**. Класс представляет упрощенный интерфейс для работы со сложной банковской системой. Клиент работает с фасадом а не напрямую с тремя фабриками.
3. Паттерн **команда** инкапсулирует пользовательские запросы в виде объектов и позваляет добавлять дополнительную логику для работы с операциями.
4. В качестве дополнения к паттерну команда применен паттерн **декоратор** для логирования операций. Он добавляет дополнительную функциональность к объекту без изменения его структуры.
5. Для экспорта данных в разных форматах используются сразу несколько паттернов. Это **надблюдатель** и **шаблонный метод**. Наблюдатель используется для уведомления наблюдателей о событиях (начало экспорта, конец, ошибка во время выполнения), а шаблонный метод для определения алгоритма сериализации данных.

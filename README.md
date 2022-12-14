Необходимо разработать приложение **EWallet** реализующее систему обработки транзакций платёжной системы.
Приложение должно быть реализовано в виде **gRPC** сервера и клиента.
Сервер ewalletd должен иметь два метода: 
1.  Send, который отправляет деньги с одного из кошельков на указанный кошелек пользователя. Метод имеет следующие параметры: 
    * from - адрес кошелька, откуда нужно отправить деньги. Например: e240d825d255af751f5f55af8d9671beabdf2236c0a3b4e2639b3e182d994c88
    * to - адрес кошелька, куда нужно отправить деньги. 
    * amount - сумма перевода. Например: 3.50.
2.  GetLast, который возвращает массив объектов, содержащий информацию о последних поступлениях на кошельки. Последними считаются поступления, которые ещё не запрашивались данным методом. Объекты массива имеют следующие поля: 
    * datetime - дата и время поступления. Например: 
    * address - адрес, на который был произведен перевод 
    * amount - сумма перевода. 
    ____
Приложение должно использовать СУБД CouchDB, в которой должны храниться следующие данные: 

    Информация об имеющихся кошельках с актуальным балансом. 
    Информация обо всех транзакциях на вывод с информацией о том, с какого кошелька был произведен перевод. 
    Информация обо всех транзакциях на ввод
    
Клиент ewallet должен иметь интерфейс командной строки с двумя командами send, getlast.

Информация должна выводиться в формате JSON.

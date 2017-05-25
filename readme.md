https://github.com/fzipp/gocyclo
https://hackernoon.com/dancing-with-go-s-mutexes-92407ae927bf
# Melhorias #

* usar Getenv
* ler info de arquivos yml
* fazer tratamento para serviço subir mesmo com BD ou Redis fora do AR
* Podemos validar o canal pelo owner, LogicNumber, pin, etc

https://smartystreets.com/blog/2015/02/go-testing-part-1-vanillla
https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742

API
* N - POST Criação de Ordem
* N - PUT Cancelamento de ordem = /api/v2/orders/<id>?operation=cancel
* N - PUT Liberando para pagamento = /api/v2/orders/<id>?operation=PLACE
* N - PUT Fechar pagamento = /api/v2/orders/<id>?operation=CLOSE
* - GET Buscar ordem = /api/v2/orders/<id>
* - GET Buscar Ordens = /api/v2/orders?<parameters>
* - DEL Deletar Ordem (deleted_at) = /api/v2/orders/<id>
* N - POST Add Item = /api/v2/orders/<id>/items
* N - PUT Editar Item = /api/v2/orders/<id>/items/<item_id>
* N - DEL Deletar Item (deleted_at) = /api/v2/orders/<id>/items/<item_id>
* - GET Consultar Items = /api/v2/orders/<id>/items
* - GET Consultar Transac = /api/v2/orders/<id>/transactions

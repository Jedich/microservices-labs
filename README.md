## Micorservices lab 1
- ***Вавринюк Максим*** - max-sevice
- ***Тітов Єгор*** - titov-service
- ***Юдаков Олександр*** - golang-service

apply command:

```kubectl apply -f titov_service/deployment -f golang_service/k8s -f max_service/k8s -f client/k8s```
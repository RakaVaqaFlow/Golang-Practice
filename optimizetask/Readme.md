1) оформляем мерж реквест к данному репо с реализацией
2) в мерж реквесте пишем в ридми(здесь) результаты бенчмарков

Задача: Не меняя сигнатуру функции и бенчмарка ускорить функцию(написана не оптимально). Можно использовать любые алгоритмы хэширования

результаты базовых бенчмарков

goos: darwin
goarch: arm64
pkg: gitlab.ozon.dev/OptimizeTask
BenchmarkHashUser
BenchmarkHashUser/old_hash,_parse_by_user_id
BenchmarkHashUser/old_hash,_parse_by_user_id-8         	 1724763	       622.2 ns/op
PASS


goos: darwin
goarch: arm64
pkg: gitlab.ozon.dev/OptimizeTask
BenchmarkHashUser
BenchmarkHashUser/old_hash,_parse_by_user_age
BenchmarkHashUser/old_hash,_parse_by_user_age-8         	 1939216	       589.4 ns/op
PASS

Process finished with the exit code 0

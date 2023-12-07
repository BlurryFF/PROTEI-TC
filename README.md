# PROTEI-TC
Приложение принимает и обрабатывает запросы на расширение/модификацию данных о пользоавтеле
# Как запустить?
```go
go run cmd/uDES/main.go
```
____
# Config
```json
{
  "grpc": {
    "ip":         "127.0.0.1", // ip адрес gRPC сервера
    "port":       ":8081", // port gRPC сервера
    "queueSize":  10, // максимальный размер очереди
    "numWorkers": 5 // количество обработчиков
  },
  "http": {
    "ip":   "localhost", // ip адрес http сервера
    "port": ":8080", // port http сервера
    "auth": {
      "username": "admin", // login для аутентификации на http сервере
      "pass":     "qwerty" // password для аутентификации на http сервере
    },
    "absencesFileName": "absences.json", // название файла, хранящего данные об отсутствиях
    "employeesFileName": "employees.json", // название файла, хранящего данные о сотрудниках
    "baseStoragePath": "internal/http/storage/" // стандартный путь до файлов с данными
  }
}
```
```grpc``` описывает конфигурацию gRPC сервера

```http``` описывает конфигурацию http сервера
___

# grpcserver

```go
func NewUserManagementServer() *UserManagementServer
```
NewUserManagementServer создает новый gRPC сервер

```go
func (server *UserManagementServer) Run() error
```
Run запускает gRPC сервер, прослушивая указанный адрес и порт, и обслуживает запросы

```go
func (server *UserManagementServer) GetEmployee(_ context.Context, in *pb.GetEmployeeParams) (*pb.EmployeesList, error)
```
GetEmployee обрабатывает запрос на получение информации о сотруднике

```go
func (server *UserManagementServer) GetAbsences(_ context.Context, in *pb.GetAbsencesParams) (*pb.AbsencesList, error)
```
GetAbsences обрабатывает запрос на получение информации об отсутствиях

___

# services

```go
func doRequest(in interface{}, endpoint string) ([]byte, error)
```
doRequest выполняет запрос на http сервер, для получения данных об отсутвии/сотруднике

```go
func checkHealth() error
```
checkHealth() проверяет доступен ли http сервер

```go
func HandleGetEmployee(in *pb.GetEmployeeParams) (*pb.EmployeesList, error)
```
HandleGetEmployee обрабатывает запрос на получение информации о сотрудниках. Возвращает список сотрудников в формате protobuf

```go
func HandleGetAbsences(in *pb.GetAbsencesParams) (*pb.AbsencesList, error)
```
HandleGetAbsences обрабатывает запрос на получение информации об отсутствиях. Возвращает список отсутствий в формате protobuf

```go
func getNameById(id int32) string
```
getNameById получает displayName сотрудника по id, полученному из информации о найденном отсутствии

```go
func getSmile(reason int) string
```
getSmile возвращает emoji, которое нужно добавить к displayName

___

# httpserver

```go
func healthCheckHandler(w http.ResponseWriter, _ *http.Request)
```
healthCheckHandler обрабатывает запрос на проверку состояния сервера

```go
func employeeInfoHandler(w http.ResponseWriter, r *http.Request)
```
employeeInfoHandler обрабатывает запрос на получение информации о сотруднике

```go
func absencesInfoHandler(w http.ResponseWriter, r *http.Request)
```
absencesInfoHandler обрабатывает запрос на получание отсутствия

```go
func loadAbsencesData() ([]models.Absences, error)
```
loadAbsencesData() выполняет загрузку данных об отсутствии из json

```go
func loadEmployeesData() ([]models.Employee, error) {
```
loadEmployeesData() выполняет загрузку данных о сотруднике из json

```go
func absenceFinder(request models.AbsencesRequest) (absences []models.Absences)
```
absenceFinder выполняет поиск отсутствия по полученным параметрам. Поиск производится по датам указанным в запросе. При незаполненном PersonID поле, вернутся все отсутствия за указанный период

```go
func employeeFinder(request models.EmployeeRequest) (employees []models.Employee)
```
employeeFinder выполняет поиск отсутствия по полученным параметрам. Поиск производится в порядке приоритета Id, Email, ФИО, workPhone

```go
func verifyUser(r *http.Request) bool
```
verifyUser проверяет логин и пароль из заголовка на соответствие с указанным в конфиге

___

# workers

```go
type Pool interface
```
Интерфейс для управления пулом воркеров

```go
type Task interface
```
Интерфейс задачи, которую выполняют воркеры

```go
type WorkerPool struct
```
Реализация интерфейса Pool

```go
func NewWorkerPool(numWorkers int, channelSize int) (Pool, error)
```
NewWorkerPool создает новый экземпляр WorkerPool.

```go
func (p *WorkerPool) Start()
```
Запускает workerPool

```go
func (p *WorkerPool) Stop()
```
Останавливает workerPool

```go
func (p *WorkerPool) AddWork(t Task)
```
Добавление задачи в очередь

```go
func (p *WorkerPool) startWorkers()
```
Запускает воркеров для выполнения задач из очереди

```go
type TaskResult struct
```
Представляет результат выполнения задачи

```go
type GetTask struct
```
Представляет задачу, которую выполняет воркер, для получения данных

```go
func (r TaskResult) GetData() interface{}
```
Возвращает данные результата задачи

```go
func (r TaskResult) GetErr() error
```
Возвращает ошибку, возникшую при выполнении задачи

```go
func (t *GetTask) Execute() error
```
Выполняет задачу в зависимости от параметров, результат отправляется в канал Result
___



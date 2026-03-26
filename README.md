<h1 align="center"> Привет! Я <a target="_blank"> Кармеев Артур из группы ЭФМО-01-25 </a> 
<img src="https://github.com/blackcater/blackcater/raw/main/images/Hi.gif" height="32"/></h1>
<h3 align="center"> Данная практика была выполнена с божьей помощью! :dizzy_face: </h3>

Структура проекта:

    └ pz2-grpc/
        ├── go.mod
        ├── go.sum
        ├── README.md
        ├── proto/
        │   └── student.proto
        ├── .idea/
        │   ├── .gitignore
        │   ├── modules.xml
        │   ├── pz2-grpc.iml
        │   ├── vcs.xml
        │   └── workspace.xml
        ├── 
        ├── internal/
        │   └── student/
        │       ├── data.go
        │       └── service.go
        ├── gen/
        │   └── studentpb/
        │       ├── student.pb.go
        │       └── student_grpc.pb.go
        └── cmd/
            ├── server/
            │   └── main.go
            └── client/
                └── main.go

## 1. Начало работы

Установили зависимости и инструменты для генерации кода.
Установили protoc, закинули путь в PATH
Сгенерили код, скопировали код из практики, смотрим как это работает:

### 1.1 Запуск сервера

<img width="945" height="342" alt="image" src="https://github.com/user-attachments/assets/3b13f004-d159-45eb-80d2-c202c932d126" />

### 1.2 Запуск клиента

<img width="974" height="279" alt="image" src="https://github.com/user-attachments/assets/c436a295-b3f9-4ff8-882a-e97f94c99068" />

### 1.3 Проверка сценария ошибки (id меняем с 1 на 999)

<img width="974" height="518" alt="image" src="https://github.com/user-attachments/assets/fd861dd9-f738-485b-8c0e-7b7ffe44eac3" />

Всё сработало, всё понятно. Что в итоге было сделано:

- описал контракт сервиса в `.proto`;
- сгенерировал код клиента и сервера;
- реализовал gRPC-сервер на Go;
- реализовал gRPC-клиент на Go;
- выполнил вызовы методов сервиса;
- получил структурированные ответы;
- обработал типовую ошибку.

## 2. Доп задание (1 вариант) 😞

<img width="633" height="169" alt="image" src="https://github.com/user-attachments/assets/3d15a940-e86d-413f-856c-8fc9ae4677c8" />

### 2.1 Добавление приколов в протокол

Добавил в `student.proto`   `import "google/protobuf/empty.proto"`
`google.protobuf.Empty` – пустой тип, используется, когда метод не требует входных параметров.
И бахнул новое сообщение в этом же `student.proto` 

```go
message ListStudentsResponse {
  repeated Student students = 1;
}

service StudentService {
  rpc Ping(PingRequest) returns (PingResponse);
  rpc GetStudentByID(GetStudentRequest) returns (GetStudentResponse);
  rpc ListStudents(google.protobuf.Empty) returns (ListStudentsResponse); // ← новый метод
}
```
Прописал новый метод в `student.proto`
Сгенерировали новый код

### 2.2 Добавление нового метода в internal/student/data.go

<img width="974" height="282" alt="image" src="https://github.com/user-attachments/assets/6dba8d0d-33f3-43ad-aa6f-b5329790b8d2" />

GetAll возвращает всех студентов из внутреннего хранилища

### 2.3 Реализация gRPC-метода в сервисе (internal/student/service.go)

<img width="974" height="179" alt="image" src="https://github.com/user-attachments/assets/af0d4be0-8377-4d46-8949-61c7085083e7" />

`ListStudents` реализует интерфейс `StudentServiceServer`
Пояснение:
-	Метод принимает `emptypb.Empty` и возвращает `ListStudentsResponse`.
-	Внутри вызывается репозиторий для получения всех студентов.
-	Создаётся ответ со списком студентов и возвращается.

### 2.4 Обновление клиента (cmd/client/main.go)

<img width="974" height="452" alt="image" src="https://github.com/user-attachments/assets/13831540-dfff-447e-ba1d-7ddc1c6d2ea3" />

Клиент передаёт пустой объект `emptypb.Empty`, получает ответ и выводит всех студентов.

### 2.5 Проверка

<img width="974" height="210" alt="image" src="https://github.com/user-attachments/assets/01005eac-1d59-44e1-93e8-7fee5dd2a2ce" />

Всё 😃






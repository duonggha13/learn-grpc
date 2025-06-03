# learn-grpc
Learn + Practice gRPC with Go

## Khái niệm
- gRPC (Google Remote Procedure Call) là một framework mã nguồn mở do Google phát triển, dùng để giao tiếp giữa các dịch vụ (service to service communication) theo mô hình RPC (Remote Procedure Call - gọi hàm từ xa)<br>
- Thường được sử dụng trong hệ thống phân tán hoặc microservices
- gRPC là một giao thức truyền thông hiệu suất cao, nhẹ, sử dụng HTTP/2 làm nền tảng
- Hỗ trợ đa ngôn ngữ: Golang, Python, Java, C++...
- gRPC định nghĩa dịch vụ thông qua file .proto (protocol buffers)
- Nó tự động sinh mã client/server từ file .proto

## Ưu điểm
- Nhanh, hiệu quả (dữ liệu nhị phân, HTTP/2)
- Hỗ trợ nhiều ngôn ngữ
- Hỗ trợ streaming tốt
- Có sẵn tính năng authentication, load balancing, deadline, retry, metadata, ...
## Hạn chế
- Không thân thiện với trình duyệt (do dùng HTTP/2)
- Debug khó hơn vì dùng nhị phân thay vì json
- Cần công cụ tích hợp với RESTful ecosystem

## Khi nào nên sử dụng gRPC
- Giao tiếp giữa các microservices nội bộ
- Hệ thống yêu cầu hiệu suất cao, truyền tải dữ liệu lớn, tần suất cao
- Cần streaming 2 chiều
- Dự án lớn, muốn sinh code client/server tự động
- Không cần tương tác trực tiếp từ trình duyệt

## So sánh với REST

|Tiêu chí| gRPC                                       | REST                                                    |
|--|--------------------------------------------|---------------------------------------------------------|
|Giao thức| HTTP/2                                     | Thường là HTTP/1.1, cũng có HTTP/2 nhưng không phổ biến |
|Định dạng dữ liệu| Protocol Buffers (binary, nhỏ gọn, nhanh)  | JSON(text-based, dễ đọc, nặng hơn)|
|Tốc độ truyền dữ liệu| Rất nhanh (binary + HTTP/2 multiplexing)   | Chậm hơn (JSON + HTTP/1.1)|
|Streaming| Hỗ trợ bidirectional streaming             | Không hỗ trợ, phải dùng WebSocket|
|Định nghĩa API| Sử dụng file .proto (có sinh code tự động) | Không chuẩn chung, thường dùng OpenAPI (Swagger)|
|Tự động sinh mã client| Có sẵn                                     | Cần dùng tool ngoài như swagger codegen|
|Debug dễ dàng| Khó hơn (binary, khó đọc)                  | Dễ hơn (JSON)|
|Trình duyệt hỗ trợ|Không hỗ trợ trực tiếp|Hỗ trợ trực tiếp|
|Hệ sinh thái hỗ trợ|Mạnh mẽ trong microservices, cloud-native|Rộng rãi, phổ biến trong Web APIs|
|Hỗ trợ đa ngôn ngữ| Có | Có|
|Quản lý version API|Khó hơn, cần thiết thế Protobuf cẩn thận | Dễ hơn thông qua URL versioning|

Nên kết hợp REST và gRPC với gRPC cho nội bộ (internal communicate) và REST cho public API hoặc frontend


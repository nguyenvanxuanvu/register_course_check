# install and run
  - Go version 1.18
  - setup gopath for environment
  - Install : make install
  - sql run dtb : change name and pw in config/local.yaml
  - Run: make run



# register_course_check

The backend for allow user send request check course register

# Document
## Squence diagram
### Check
```plantuml
autonumber
participant "Unitime UI" as ui
participant "Check Service" as service
participant "Redis" as redis
participant "Core Service" as core



ui -> service: check register
alt student status exist in redis
service -> redis: get student status
redis --> service: response

else not exist
service -> core: get student status
core --> service: response
end










loop for each register course
service -> service: get course
service -> service: check course condition
alt course condition is "TQ","HT"
service -> service: check done course list of student
service -> dtb: get done course list
dtb --> service: response


else course condition is "SH"
service -> service: check course condition exist in register course list of student
end

service -> service: check min credit
service -> dtb: get min credit config
dtb --> service: response

service --> ui: response

```



## Database 


### course
| Field           | Type         | Description                            | Example |
|-----------------|--------------|----------------------------------------|---------|
| id (PK)       | varchar(45)      |  MSMH                           | CO3001       |
| subject_name          | VARCHAR(45) |  Tên môn học  |  Kiểm tra phần mềm   |
| num_credits | int      |  Số tín chỉ môn học      | 3    |
| faculty | varchar(45)      | Khoa     | KHMT    |

### course_condition
| Field           | Type    | Description               | Example        |
|-----------------|---------|---------------------------|----------------|
| id (PK)      | int | Id            | 1              |
| subject_id     | varchar(45)    | MSMH     | CO3002 |
| subject_des_id | varchar(45)    | MSMH của môn điều kiện | CO3001 |
| condition_type | int    |  Loại điều kiện | 1: Tiên quyết  2:  Học trước   3: Song hành |
### min_max_credit
| Field           | Type      | Description         | Example             |
|-----------------|-----------|---------------------|---------------------|
| id (PK)  | bigint  | Id        | 1                   |
| student_id   | int   | MSSV        | 1915982                 |
| subject        | varchar(45) | MSMH   | CO3001 |
| result           | int   |  Kết quả môn học của sinh viên    | 1: Đạt  2: Chưa đạt  |


### Min_Credit
| Field           | Type         | Description                            | Example |
|-----------------|--------------|----------------------------------------|---------|
| id (PK)       | int      |  Id                           | 1       |
| academic_program          | varchar(45) | Chương trình đào tạo | DT: Đại trà  CLC: Chất lượng cao  |
| semester | int      | Kì     | 202    |
| min_credit | int     | Số tín chỉ tối thiểu     | 11   |

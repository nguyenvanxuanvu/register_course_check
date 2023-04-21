## Database 


### course: Thông tin về môn học
| Field           | Type         | Settings |                 Description                      | Example | Default & Notes |
|-----------------|----------|----------|----------------------------------------|------|---|
| id (PK)       | varchar(45)      |  PK, not null | Mã số môn học                           | CO3001       | |
| course_name          | VARCHAR(45) |  not null | Tên môn học  |  Kiểm tra phần mềm   | |
| num_credits | int      |  not null | Số tín chỉ của môn học      | 3    | |
| faculty | varchar(45)   |   | Khoa     | KHMT    | NULL |

### course_condition: Thông tin về điều kiện đăng ký môn học (tiên quyết, song hành, học trước)
| Field           | Type    | Settings | Description               | Example        |  Default & Notes |
|-----------------|--------|--------|---------------------------|------|----------|
| id (PK)      | int | PK, not null | Id của điều kiện       | 1              | |
| course_id     | varchar(45)  | not null  | Mã số môn học     | CO3002 | |
| conditon     | varchar(45)  | not null  | Điều kiện của môn học     | {"course": {"type": 1, "courseDesId": "ENGL"}} | "type": 1 - Tiên quyết, 2 - Song hành, 3 - Học trước


### min_max_credit: Thông tin cấu hình về số tín chỉ tối thiểu, tối đa
| Field           | Type      | Settings | Description         | Example             |  Default & Notes |
|-----------------|------|-----|---------------------|--------------|-------|
| id (PK)  | int  | PK, increment, not null | Id của cấu hình      | 1            |       |
| academic_program   | varchar(45)  | not null | Chương trình đào tạo        | DT           |      |
| semester        | int | not null | Kỳ học   | 191 | |
| min_credit          | int   |  not null | Số tín chỉ tối thiểu    | 14  | -1 |
| max_credit          | int   |  not null | Số tín chỉ tối đa    | 21  | -1 |


### teaching_plan: Thông tin về các môn học theo chương trình đào tạo
| Field           | Type        | Settings | Description                            | Example |  Default & Notes |
|-----------------|----------|----|----------------------------------------|---|------|
| id (PK)       | int      |  PK, increment, not null | Id của thông tin                        | 1       | |
| falcuty          | varchar(45) | not null | Khoa | KHMT  |
| speciality | varchar(45)     | not null | Chuyên ngành     | CNPM    | |
| academic_program | varchar(45)  |  not null  | Chương trình đào tạo     | CLC   | |
| semester_order | int    | not null | Số thứ tự của kỳ học     | 3   | |
| course_list | longtext   | not null   | Danh sách môn học     | ["CS", "CALC"]   | |
| free_credit_info | longtext     | | Danh sách môn học tự chọn theo nhóm   | [{"nums": 3,"group": "C"},{"nums": 6,"group": "D"} ]   | NULL |

### white_list: Thông tin sinh viên với cấu hình số tín chỉ tối thiểu, tối đa đặc biệt 
| Field           | Type      | Settings   | Description                            | Example |  Default & Notes |
|-----------------|--------|------|----------------------------------------|-----|----|
| id (PK)       | int   | PK, increment, not null  |  Id của thông tin                       | 1       | |
| student_id         | varchar(45) | not null | Mã số sinh viên | 1915982  | |
| semester | int    | not null  | Kì học    | 202    | |
| min_credit | int  | not null   | Số tín chỉ tối thiểu     | 11   | |
| max_credit | int  | not null   | Số tín chỉ tối đa     | 21   | |
| description | varchar(45)   |  | Mô tả     | Đặc biệt   | NULL |


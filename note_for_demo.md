# Map data from Unitime data
  - Student 1 : 1915982
  - Student 2: 1915983
  - Student 3: 1915984  # Ko đc phép đk môn 

# Map data semester 
  - Kì đang chạy nơi Unitime -> 191
# Map data môn học 


# Request body description

```
{
    "studentId": 1915982,
    "academicProgram": "DT",
    "semester": 191,
    "registerSubjects": [
        
        {
            "subjectId": "CO3"
        },
        {
            "subjectId": "CO2"
        },
        {
            "subjectId": "CO1"
        }

        
    ]
        
}
```
- studentId : map từ tên trên Unitime UI ra mssv
- academicProgram: Tạm thời để: DT
- semester: Cũng nên map từ data Unitime 
- registerSubjects: danh sachs môn đăng ký, lấy data kịp thì lấy nguyên trên Ui, không thì map 

# Response successful description

```
    {
    "data": {
        "status": "FAIL",
        "studentStatus": "NORMAL",
        "subjectChecks": [
            {
                "subjectId": "CO3",
                "subjectName": "CCC",
                "checkResult": "FAIL",
                "failReasons": [
                    {
                        "subjectDesId": "CO1",
                        "conditionType": 1
                    },
                    {
                        "subjectDesId": "CO2",
                        "conditionType": 1
                    }
                ]
            }
        ],
        "checkMinCreditResult": "PASS"
    }
}
```
- status: FAIL hoặc PASS (pass thì hiển thị lên UI, fail thì có 3 loại có thể fail -> lấy danh sách hiển thị UI)
- studentStatus: NORMAL hoặc NOT_PERMIT_REGISTER_COURSE
- subjectChecks: Danh sách check các môn đăng ký (trong đó fail reason là một list). Như vd trên thì môn CO3 bị fail do thiếu môn Tiên quyết là CO1 và CO2
- checkMinCreditResult: PASS hoặc FAIL 


# Response fail (Khi gặp lỗi kiểu không có data cho sv đó, tìm ko có config số tín chỉ tối thiểu)

VD:
```
{
    "error": {
        "code": 503,
        "reason": "NOT_FOUND_STUDENT_STATUS",
        "domain": "register_course_check"
    }
}
```




  
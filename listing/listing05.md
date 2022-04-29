Что выведет программа? Объяснить вывод программы.
```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод программы:
error

Переменной err присваивается путем вызова функции test тип customError, который удовлетворяет интерфейсу error,
поэтому может использоваться. Однако при присвоении переменная err будет не с типом nil, а с типом *customError, 
и, соответственно, перестанет быть равна nil.
```
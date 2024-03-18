package contexts

import (
	"context"
	"fmt"
	"time"
)

// Контекст служит для двух целей:
// а) Хранение некоторых значений (произвольных)
// б) Сообщение о завершении

// Базовые контексты - Background и TODO

func Cons() {
	// Основной, родительский контекст
	ctx := context.Background()

	// Пример создания контекста
	// Прямое добавление значений в контекст считается анти-паттерном
	withValue := context.WithValue(ctx, "id", 15)
	fmt.Println("WithValue", withValue)

	// Завершающийся контекст. Возвращает контекст и CancelFunc (функция, завершающая контекст)
	withCancel, cancelFunc := context.WithCancel(ctx)
	fmt.Println("withCancel error:", withCancel.Err())
	cancelFunc()
	fmt.Println("withCancel error:", withCancel.Err())

	// Deadline-контекст. Завершается в течение определенного времени, указанного в качестве аргумента
	deadline, deadlineFunc := context.WithDeadline(ctx, time.Now().Add(time.Second*2))
	defer deadlineFunc()

	// endTime возвращает дату и время завершения процесса, ok возвращает true в случае успешного завершения, иначе false
	endTime, ok := deadline.Deadline()
	fmt.Println("Deadline:", endTime, ok)

	// Вывод ошибки и результата работы канала Done
	fmt.Println(deadline.Err())
	fmt.Println(<-deadline.Done())

	withTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	fmt.Println("withTimeout:", withTimeout)

	// TODO-контекст. Используется для тестов
	toDo := context.TODO()
	fmt.Println(toDo)
}

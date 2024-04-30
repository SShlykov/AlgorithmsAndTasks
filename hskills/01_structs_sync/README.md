<div id="top" style="font-weight: bolder; font-size: 1.5em">Первое занятие: Синтаксис + Синхронизация </div>

<div style="margin: 10px 0; ">
    <div style="padding: 10px 0 0 10px; font-weight: bolder; font-size: 1.2em">TOC</div>
    <ul>
        <li>Задание 1 (<a href="#task1">link</a>)</li>
        <li>Задание 2 (<a href="#task2">link</a>)</li>
        <li>Задание 3 (<a href="#task3">link</a>)</li>
        <li>Задание 4 (<a href="#task4">link</a>)</li>
        <li>Задание 5 (<a href="#task5">link</a>)</li>
        <li>Задание 6 (<a href="#task6">link</a>)</li>
        <li>Задание 7 (<a href="#task7">link</a>)</li>
        <li>Задание 8 (<a href="#task8">link</a>)</li>
        <li>Задание 9 (<a href="#task9">link</a>)</li>
        <li>Задание 10 (<a href="#task10">link</a>)</li>
    </ul>
</div>

<div id="task1" style="font-weight: bold">Задание 1</div>
<pre>
package main

import "fmt"

func main() {
    s := "hello"
    s[0] = 'H'
    fmt.Println(s)
}
</pre>

__Что происходит:__<br>
_S является строкой, в первый байт строки мы пытаемся заменить на литерал руны_<br>
_Это сработало бы во многих языках, но Golang запрещает это делать т.к. строки не мутабельны_<br>

__В результате:__<br>
_Скорее всего компилятор не даст собрать такую программу_<br>
_Кроме того, тут присутствует не совместимость типов, т.к. по дефолту строка - набор UTF8 байт, а мы пытаемся подставить rune_<br>

__Как решить:__<br>
_через strings Buffer_<br>
_разрезать строку и объединить через пакет fmt_<br>

<div style="text-align: right">(<a href="#top">top</a>)</div>

-----

<div id="task2" style="font-weight: bold">Задание 2</div>
<pre>
package main

func main() {
    ch := make(chan int, 1)
    for i := 0; i < 10; i++ {
        select {
        case x := <-ch:
            println(x)
        case ch <- i:
        }
    }
}
</pre>

__Что происходит:__<br>
_Создается буферизированный канал из int, далее по циклу вызывается select_<br>
_Сам select равновероятно обрабатывает 2 ситуации - чтение из канала и запись в него_<br>

__Что может произойти:__<br>
_Идеальная ситуация - паттерн запись - чтение, число записано и, в следующей итерации прочитано_<br>
_Хотя селект не гарантирует последовательность операции, может произойти так, что буфер заполнится; и следующая попытка записать заблокирует поток исполнения и будет ожидать прочтения_<br>
_На самом деле не так, селект умнее_<br>

<div style="text-align: right">(<a href="#top">top</a>)</div>

-----

<div id="task3" style="font-weight: bold">Задание 3</div>
<pre>
package main

import "fmt"

func add(s []string) {
    s = append(s, "x")
}

func main() {
    s := []string{"a", "b", "c"}
    r := s[1:2]
    add(r)
    fmt.Println(s)
}
</pre>

__Что происходит (len/cap):__ <br>
_Инициализируется слайс из строк s(3,3), выбирается фрагмент из слайса r(1,2);_<br>
_К "подслайсу" применяется функция, которая потенциально должна добавить "x"; После чего мы должны вывести s_<br>

__Какие проблемы:__<br>
_добавление в "подслайс" изменит исходный слайс (len < cap), s предается включая ссылку из r_<br>

__Как исправить:__<br>
_Смотря что мы хотим сделать... Выглядит как будто мы хотим сделать копию r и обновить ее._<br>
_В целом такая идея плоха, сайд-эффекты всегда приводят к сомнительному поведению; лучше сделать так, чтобы функция вернула значение и записала_<br>
_но если очень хочется - создаем r через copy(); передаем в add ссылку; в функции обновляем значение_<br>

<div style="text-align: right">(<a href="#top">top</a>)</div>

-----

<div id="task4" style="font-weight: bold">Задание 4</div>
<pre>
package main

func main() {
    digits := []int{1, 2, 3, 4, 5}
    for _, d := range digits {
        defer println(d)
    }
}
</pre>

__Что происходит:__<br>
_Берем список из чисел и заполняем defer; должно вывести цифры в обратном порядке;_<br>
_println из фичей, лучше не использовать;_<br>

__Какие проблемы:__<br>
_потенциальная проблема с протечкой данных, т.к. println(d) вызывается внутри цикла (не явная очередь исполнения),_<br>
_в простом примере очевиден вывод (обратный), но в усложненных задачах можно и пропустить;_<br>

__Как исправлять:__<br>
_завернуть в callback; вызвать без defer_<br>

<div style="text-align: right">(<a href="#top">top</a>)</div>

-----

<div id="task5" style="font-weight: bold">Задание 5</div>
<pre>
package main

func main() {
    defer func() {
        if err := recover(); err != nil {
            println(err)
        }
    }()
    go func() {
        panic(123)
    }()
    time.Sleep(time.Hour)
}
</pre>

__Что происходит:__<br>
_объявили в defer callback который призван к восстановлению в случае паники в рутине main_<br>
_запустили отдельную рутину и вызвали панику, замкнули рутину main с помощью sleep_<br>
_Паника будет запущена в рутине и приложение остановится, так как в ней нет обработки паники_<br>

__Как исправлять:__<br>
_перенести обработку паники в рутину, которая паникует\может запаниковать_<br>

<div style="text-align: right">(<a href="#top">top</a>)</div>

----

<div id="task6" style="font-weight: bold">Задание 6</div>
<pre>
package main

import (
    "fmt"
    "runtime"
)

func main() {
    runtime.GOMAXPROCS(1)
    ch := 0
    go func() {
        ch = 1
    }()
    for ch == 0 {
    }
	fmt.Println("finish")
}
</pre>

__Что происходит:__<br>
_ограничили нашу программу 1 физическим потоком; объявили "общую" переменную для main и еще одной рутиной_<br>
_запустили отдельную рутину, которая изменит переменную; Замкнули рутину main в цикл, который остановится только после изменения переменной_<br>
_По идее программа не должна дойти до finish, тк. Всего 1 поток. Но по факту, в данный момент у Go вытесняющая многозадачность_<br>

__Как исправлять:__<br>
_Не писать так - тут потенциальная гонка; компилятор может не обновить значение_<br>
_Использовать канал, atomic, mutex, wg_<br>

<div style="text-align: right">(<a href="#top">top</a>)</div>

---

<div id="task7" style="font-weight: bold">Задание 7</div>
<pre>
package main

import "fmt"

func main() {
    values := []int{}
	for _, v := range []int{1, 2, 3, 4, 5} {
		values = append(values, v)
	}
	pointers := []*int{}
	for _, v := range values {
		pointers = append(pointers, &v)
	}
	for _, v := range pointers {
		fmt.Println(*v)
	}
}
</pre>

__Что происходит:__<br>
_присваиваем пустой слайс из интов в переменную values, переносим значения из цикла (1..5) внутрь values_<br>
_присваиваем пустой слайс из ссылок на числовое значение, переносим ссылки на значение итератора внутрь слайса_<br>
_По циклу печатаем разыменованное из предыдущего слайса значение; Ссылка (на v) во второй итерации одинаковая и мы получим 5-5-5-5-5 по итогу_<br>

__Как исправлять:__<br>
_Создавать новую ссылку из values; не использовать v;_<br>

<div style="text-align: right">(<a href="#top">top</a>)</div>

---

<div id="task8" style="font-weight: bold">Задание 8</div>
<pre>
package main

import "fmt"

func main() {
    a := unsafe.Pointer(&struct{}{})
    b := unsafe.Pointer(&struct{}{})
    if print {
        fmt.Println(a)
        fmt.Println(b)
    }
    // Что напечатает этот вызов
    // 1. Если print == true
    // 2. Если print == false
    // 3. Будет ли разница, почему
    fmt.Println(a == b)
}
</pre>

__Что происходит:__<br>
_Создаем 2 переменные, в которые складываем ссылку на значение инициализированной пустой структуры;_<br>
_внешняя переменная контролирует запись в лог этих ссылок; проверяется равенство этих ссылок_<br>
_1) напечатается 2 ссылки на пустую структуру и true; 2) напечатается только true_<br>
_согласно спецификации одинаковые нулевые значения могут иметь одинаковую ссылку в памяти; с другой стороны не понятно, оптимизирует ли значения компилятор_<br>

<div style="text-align: right">(<a href="#top">top</a>)</div>

---

<div id="task8" style="font-weight: bold">Задание 8</div>

<pre>
package main

import (
	"fmt"
	"math/rand"
	"time"
)


// REVIEW: не стоит использовать магические цифры для предсказуемости поведения
// REVIEW: rand.Intn(1000) - [0..999] приведет к непредсказуемости во время модульного тестирования + CI
// REVIEW: следует обратить внимание, лучше либо вынести в конфиг или в параметр функции
func fetchDataWrong(source string, data chan<- string) {
    // Симуляция ответа запроса разного по продолжительности.
    time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

    // Симуляция ошибки в 30%.
    // REVIEW: следует использовать rand.Float64 чтобы задавать процент явно
    // REVIEW: 3, 10 - магические числа
    if rand.Intn(10) < 3 {
        // REVIEW: сформировать структуру под ошибку и вынести значение
        data <- fmt.Sprintf("Failed to fetch data from %s", source)
        return
    }

    data <- fmt.Sprintf("Data from %s", source)
}

// REVIEW: Слишком сложная функция; 
// REVIEW: Необходимо разделить функциональность конфигурации, управления go-рутинами и логикой
// REVIEW: Нужно обоснование, почему функция использует один и тот-же канал для успешной обработки и ошибок;
func main() {
    sources := []string{"Source1", "Source2", "Source3", "Source4", "Source5"}

    // REVIEW: нужно обоснование почему используется канал из 100 записей, тогда как всего 5 рутин (используется больше памяти, чем необходимо);
    // REVIEW: Для такой ситуации можно не тратить память и использовать не буферизированный канал.    
    // REVIEW: нет опций чтобы его закрыть; Как обосновывается то, что не будет утечки go-рутин?
    data := make(chan string, 100)

    for _, source := range sources {
        go fetchDataWrong(source, data)
    }

    // REVIEW: Ошибки и успехи пишутся в один канал. Стоит создать структуру передаваемому параметру\ошибке.  
    for i := 0; i < len(sources); i++ {
        fmt.Println(<-data)
    }
}

</pre>

<div style="text-align: right">(<a href="#top">top</a>)</div>

---

<div id="task9" style="font-weight: bold">Задание 9</div>

<pre>
package throttle

import (
    "time"
    "sync"
)

// HOF и замыкание
func Throttle(f func(), duration time.Duration) func() {
    var once sync.Once 
    var mu sync.Mutex
    var lastTime time.Time

    return func() {
        mu.Lock()
        defer mu.Unlock()

        once.Do(func() {
            lastTime = time.Now().Add(-duration)
        })

        if time.Since(lastTime) < duration {
            return
        }

        f()

        lastTime = time.Now()
    }
}
</pre>

<div style="text-align: right">(<a href="#top">top</a>)</div>

---

<div id="task10" style="font-weight: bold">Задание 10</div>

<pre>
package cache

import "sync"

type item[T any] struct {
	value T
}

type InMemoryCache[Key comparable, Val any] struct {
	mu    sync.RWMutex
	cache map[Key]item[Val]
}

func NewInMemoryCache[Key comparable, Val any]() *InMemoryCache[Key, Val] {
	return &InMemoryCache[Key, Val]{
		cache: make(map[Key]item[Val]),
	}
}

func (c *InMemoryCache[Key, Val]) Put(key Key, value Val) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = item[Val]{value: value}
}

func (c *InMemoryCache[Key, Val]) Get(key Key) (Val, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.cache[key]
	if ok {
		return v.value, true
	}

	return v.value, false
}

func (c *InMemoryCache[Key, Val]) GetOrCreate(key Key, value Val) Val {
	if v, ok := c.Get(key); ok {
		return v
	}

	c.Put(key, value)
	return value
}
</pre>

<div style="text-align: right">(<a href="#top">top</a>)</div>
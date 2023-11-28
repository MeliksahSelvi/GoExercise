package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func main() {
	//dataTypes()
	//typeConversion()
	//shadowing()
	//strConv()
	//ifStatements()
	//switchCaseStatement()
	//switchCasefallthrough()
	//forLoopAndGoto()
	//nonIdiomaticIfStatement()
	//idiomaticIfStatement()
	//executeTwoFunctions()
	//readValueFromKeyboard()
	//blankIdentifier()
	//doubleReturnFunction
	//checkError()
	//produceRandomNumber()
	//arrays()
	//slice()
	//copyArrayIsPassByValue()
	//copySliceIsPassByReference()
	//copySliceFromArrayAndMakeChanges()
	//appendSlice()
	//appendElementsToEmptySlice()
	//appendSliceFromAnotherSliceAndMakeChangesOldSlice()
	//deleteSliceElementsFromStart()
	//deleteSliceElementsFromEnd()
	//declaredSliceAndDeclaredMakeSlice()
	//mapExample()
	//addElementToMapAfterInitialize()
	//deleteElementFromMap()
	//copyMapIsPassByReference()
	//mapRange()
	//copyMethodIsPassByValue()
	//mapExampleWithArray()
	//varStructExample()
	//typeStructExample()
	//typeStructRangeExample()
	//structCopyIsPassByValue()
	//anonymousStructExample()
	//definitionType()
	//pointerExample()
	//changeValueByPointer()
	//pointerParameterMethod()
	//receiverMethod()
	//pointerReceiver()
	//interfaceExample()
	//polymorphism()
	//goroutineExample()
	//goroutineExample2()
	//goroutinesWithWaitGroup()
	//channel()
	//deadLockWithChannel()
	//infinitiveLoopWithChannel()
	//deferKeyword()
	//doPanic()
	//handlePanic()
	//raceCondition()
	//raceConditionSolutionWithMutex()
	//raceConditionWithAtomic()
	//raceConditionWithAtomicSolution()
	//unbuffChan()
	//goRoutineUndeterministic()
	//bufferedChan()
	//bufferedChannelsUndeterministic()
	//selectChannel()
	//multipleSelectCase()
	//multipleSelectCaseSolution()
}

func multipleSelectCaseSolution() {
	chan1 := make(chan int, 1)
	chan1 <- 1

	chan2 := make(chan int, 1)
	chan2 <- 2
	var done bool
	for !done {
		select {
		case c1Val := <-chan1:
			fmt.Println(c1Val)
		case c2Val := <-chan2:
			fmt.Println(c2Val)
		default:
			done = true
		}
	}
}

func multipleSelectCase() {
	chan1 := make(chan int, 1)
	chan1 <- 1

	chan2 := make(chan int, 1)
	chan2 <- 2

	select {
	case c1Val := <-chan1:
		fmt.Println(c1Val)
	case c2Val := <-chan2:
		fmt.Println(c2Val)
	}
	//select case yapısında birden fazla case varsa select aralarında random seçim yapar.
}

func selectChannel() {
	chan1 := make(chan int, 1)
	chan1 <- 1

	select {
	case c1Val := <-chan1:
		fmt.Println(c1Val)
	}
}

func bufferedChannelsUndeterministic() {
	bufferedChannel := make(chan int, 3)

	go func() {
		for {
			val := <-bufferedChannel
			fmt.Println(val)
		}
	}()
	bufferedChannel <- 1
	bufferedChannel <- 2
	bufferedChannel <- 3
	/*
		channel değer atama işlemlerinin hepsi unblocking çalıştığı için ana goroutine onları beklemiyor.
		eğer 3 den fazla atama işlemi yaparsak 3 den sonrası blocking durumuna girecek.
		bu durum sayesinde ana goroutine bazı goroutine'leri beklemiş olacak.
	*/
}

func bufferedChan() {
	bufferedChannel := make(chan int, 2)

	bufferedChannel <- 1
	bufferedChannel <- 2
	//bufferedChan <- 3

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
	//fmt.Println(<-bufferedChan)
	/*
		en fazla 2 değer unblocking bir şekilde işlenir.
		2.den sonraki değerler ise unbuffered channels gibi blocking olur. 2.işleme kadar asynchorous çalışır
	*/
}

func goRoutineUndeterministic() {
	unBuffChannel := make(chan int)

	//reader goroutine
	go func() {
		value := <-unBuffChannel
		fmt.Println(value)
	}()

	//writer goroutine
	go func(unBufChan chan int) {
		unBuffChannel <- 1
	}(unBuffChannel)

	time.Sleep(time.Second)
	//goroutine'lerde output işlemleri undeterministic'dir. yani kod bloğuna onu yazdık diye çalıştığını ekranda görmemiz kesin değildir.bunu dengelemek için sleep ekledik.
}

func unbuffChan() {
	unbuffChannel := make(chan int)
	//fmt.Println(unbuffChannel)
	go func(unbuff chan int) {
		value := <-unbuff //bu kod satırı veri gelene kadar blocking'dir
		fmt.Println(value)
	}(unbuffChannel)
	unbuffChannel <- 5
	//unbuffered channels synchorized çalışır.input ve output işlemleri blockingdir.
}

func raceConditionWithAtomicSolution() {
	raceTest := &RaceTest{}

	wg2 := &sync.WaitGroup{}
	wg2.Add(10000)

	for i := 0; i < 10000; i++ {
		go increment(raceTest, wg2)
	}

	wg2.Wait()

	fmt.Println(raceTest)
	//increment func içindeki yorum satırındaki gibi artırma yaparsak racecondition durumu oluşur.
}

func increment(rt *RaceTest, wg2 *sync.WaitGroup) {
	//rt.Val+=1
	atomic.AddInt32(&rt.Val, 1)
	wg2.Done()
}

type RaceTest struct {
	Val int32
}

func raceConditionWithAtomic() {
	wg.Add(2)

	var val int32 = 0

	go func() {
		for i := 0; i < 100000; i++ {
			atomic.AddInt32(&val, 1)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			atomic.AddInt32(&val, 1)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(val)
}

func raceConditionSolutionWithMutex() {
	wg.Add(2)
	val := 0

	lock := sync.Mutex{}

	go func() {
		for i := 0; i < 100000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(val)
	//burada lock yaparak bir goroutine bir value üzerinde çalışırken diğer goroutine'lerin çalışmasını engelleyen bir mekanizma kurduk.
}

func raceCondition() {
	wg.Add(2)
	val := 0
	go func() {
		for i := 0; i < 100000; i++ {
			val++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			val++
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(val)
	// iki goroutine de aynı anda aynı veriyi kullanmaya çalıştığı için veri kaybı oldu.
}

func goroutineExample2() {
	for i := 0; i < 10; i++ {
		go func(value int) {
			fmt.Println(value)
		}(i)
	}
	time.Sleep(3 * time.Second)
	/*
		Bu sefer güncel i değerleri ile goroutine'ler çalıştı çünkü bu sefer onlara i'nin güncel halini value olarak copy işlemi ile verdik
	*/
}

func goroutineExample() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(3 * time.Second)
	/*
		Burada sürekli 10 değeri basıldı sebebi şu şekilde açıklanabilir. Döngü her döndüğünde bir goroutine açıldı
		ve döngü bittiğinde bu goroutine'ler schedule oldu. goroutine'ler i'nin referansındaki son değer ile çalıştılar.
		i'nin son değeri döngü biterken 10 olmuştu.
	*/
}

func handlePanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	for i := 0; i < 3; i++ {
		fmt.Println("i-> ", i)
		time.Sleep(1 * time.Second)
		if i == 1 {
			panic("Beklenmeyen bir hata oluştu")
		}
	}
}

func doPanic() {
	//uygulamayı panic ile durdurduk.
	for i := 0; i < 3; i++ {
		fmt.Println("i-> ", i)
		time.Sleep(1 * time.Second)
		if i == 1 {
			panic("Beklenmeyen bir hata oluştu")
		}
	}
}

func deferKeyword() {
	defer fmt.Println("uygulama duracak")
	defer fmt.Println("func bitti")

	for i := 0; i < 3; i++ {
		fmt.Println("i->", i)
	}

	//func bittiğinde deferler ters sıra ile çalışır.
}

func infinitiveLoopWithChannel() {
	c := time.After(5 * time.Second)
	for {
		b := false

		select {
		case <-c:
			b = true
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
		if b {
			break
		}
	}
}

func appendElementsToEmptySlice() {
	mySlice := make([]int, 0, 3)
	fmt.Println(mySlice)
	//append aynı zamanda boyutu belli olmayan yani length değeri 0 olan array'lere eleman eklemeyi de sağlar
	mySlice = append(mySlice, 7, 9)
	fmt.Println(mySlice)
}

func deadLockWithChannel() {
	myChannel := make(chan string)
	myChannel <- "ali"
	fmt.Println(<-myChannel)
	/*
		channel başka bir goroutine tarafından işlenene kadar bulunduğu goroutine'yi bloklar.
		biz burada main goroutine üzerinde channel işlemi yaptığımız için onu işleyecek başka bir goroutine olmadığı için deadlock hatası alırız.
	*/
}

func channel() {
	/*
		Normalde goroutine açtığımız yerde return değeri olmamalı. çünkü ana routine yeni açılan goroutineye bağımlı olmuş oluyor.
		golang bu duruma izin vermiyor. Bu durumu çözebilmek için goroutine'ler arası iletişim sağlayacak bir yapıya ihtiyaç var.
		bu yapı channel.
	*/
	myChannel := make(chan string)
	go func(chan1 chan string) {
		chan1 <- "merhaba"
	}(myChannel)
	fmt.Println("channel adres", myChannel)
	fmt.Println("channel değer", <-myChannel)
}

func goroutinesWithWaitGroup() {
	/*
		önce add methodu ile kaç tane goroutine bekleyeceğini waitgroup'a söylüyoruz.
		ardından Wait methodunun çağrıldığı kod satırı işlendiğince goroutine blocklanıyor
		done methodu çağrıldığı zaman wait methodunun oluşturduğu block kalkıyor.
	*/
	wg.Add(1)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Print("X")
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println()
	func() {
		for i := 0; i < 20; i++ {
			fmt.Print("Y")
		}
	}()
}

func polymorphism() {
	t1 := triangle{5, 8}
	sq1 := square{7}
	c1 := circle{3}
	showShapesAreas(t1, sq1, c1)
}

func showShapesAreas(shapes ...mathShape) {
	for _, shape := range shapes {
		fmt.Println("Alan -> ", shape.mathShapeArea())
	}
}

type mathShape interface {
	mathShapeArea() float64
}

type triangle struct {
	a, h float64
}

func (t triangle) mathShapeArea() float64 {
	return (t.a * t.h) / 2
}

type square struct {
	x float64
}

func (s square) mathShapeArea() float64 {
	return s.x * s.x
}

type circle struct {
	r float64
}

func (c circle) mathShapeArea() float64 {
	return math.Pi * c.r * c.r
}

func interfaceExample() {
	r1 := rectangle{3, 8}

	interfaceFunction(r1)
}

type shape interface {
	area() float64
	circumference() float64
}

func pointerReceiver() {
	k := Kullanici{
		Ad:    "Melikşah",
		Soyad: "Selvi",
		Takipci: []Kullanici{
			{
				Ad:    "Takipçi1",
				Soyad: "1",
			},
			{
				Ad:    "Takipçi2",
				Soyad: "2",
			},
		},
	}

	t := Kullanici{
		Ad:    "Takipçi3",
		Soyad: "3",
	}

	fmt.Println(k.TakipciSayisi())
	k.TakipciEkle(t)
	fmt.Println(k.TakipciSayisi())
	/*
		Kullanici struct'ının TakipçiEkle functionu parametre olarak gelen kullanici üzeride değişiklik yapacağı için
		pointer receiver olarak tasarlanırsa değişiklik kalıcı olacaktır.Parametre olarak gelen structun valuesini değil
		adress'ini pass yapmış oluyoruz.
	*/
}

type Kullanici struct {
	Ad      string `json:"adi"`
	Soyad   string `json:"soyadi"`
	Takipci []Kullanici
}

func (k Kullanici) TakipciSayisi() int {
	return len(k.Takipci)
}

func (k *Kullanici) TakipciEkle(t Kullanici) {
	if k.Takipci == nil {
		k.Takipci = []Kullanici{}
	}
	k.Takipci = append(k.Takipci, t)
}

func receiverMethod() {
	r1 := rectangle{4, 5}

	fmt.Println("Area -> ", r1.area())
	fmt.Println("Circumference -> ", r1.circumference())
	//burada area ve circumference receiver olarak adlandırılır.
}

func interfaceFunction(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.circumference())
	fmt.Printf("%T", s)
	fmt.Println()
}

type rectangle struct {
	a, b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func (r rectangle) circumference() float64 {
	return 2 * (r.a + r.b)
}

func pointerParameterMethod() {
	x := 5
	fmt.Println(x)
	double(&x)
	fmt.Println(x)
}

func double(num *int) {
	fmt.Println(num)
	*num *= 2
	fmt.Println(*num)
}

func changeValueByPointer() {
	x1 := 10
	x2 := &x1
	fmt.Println(x1, x2)
	fmt.Println(x1, *x2)
	*x2 = 3
	fmt.Println(x1, *x2)
}

func pointerExample() {
	x := 22
	fmt.Println(x)
	fmt.Println(&x)          //x'in bellekteki adresi
	fmt.Println(*(&x))       // x'in bellekteki adresinin gösterdiği value yani x
	fmt.Println(&(*(&x)))    // x'in bellekteki adresinin gösterdiği value'nin bellekteki adresi yani &x
	fmt.Println(*(&(*(&x)))) // x'in bellekteki adresinin gösterdiği value'nin bellekteki adresinin valuesi yani x
}

func definitionType() {
	type mile float64
	m1 := mile(7.2)
	m2 := mile(2.1)
	fmt.Println(m1 + m2)
}

func anonymousStructExample() {
	theBoss := struct {
		name  string
		money bool
	}{name: "The Boss", money: true}

	fmt.Println(theBoss)
}

func structCopyIsPassByValue() {
	type employee struct {
		name string
		age  int
		kids []string
	}
	e1 := employee{
		"Ali", 26, []string{"Mehmet", "Burak"},
	}
	e2 := e1
	fmt.Println("e1->", e1)
	fmt.Println("e2->", e2)
	fmt.Println("e2 name değişti")
	e2.name = "Veli"
	fmt.Println("e1->", e1)
	fmt.Println("e2->", e2)
}

func typeStructRangeExample() {
	type employee struct {
		name string
		age  int
		kids []string
	}

	e1 := employee{
		"Ali", 26, []string{"Mehmet", "Burak"},
	}
	for index, value := range e1.kids {
		fmt.Println(index, value)
	}
}

func typeStructExample() {
	type employee struct {
		name      string
		age       int
		isMarried bool
	} //bu type struct function dışında tanımlamak mantıklı.
	e1 := employee{
		"Ali", 25, false,
	}

	e2 := employee{
		"Kerem", 45, true,
	}
	fmt.Println(e1)
	fmt.Println(e2)
}

func varStructExample() {
	var employee struct {
		name      string
		age       int
		isMarried bool
	}
	fmt.Printf("%#v\n", employee)
	fmt.Println(employee)
	employee.name = "Melik"
	employee.age = 26
	employee.isMarried = false
	fmt.Printf("%#v\n", employee)
	fmt.Println(employee)

	var employee2 struct {
		name      string
		age       int
		isMarried bool
	}

	fmt.Printf("%#v\n", employee2)
	fmt.Println(employee2)
	employee2.name = "Kerem"
	employee2.age = 68
	employee2.isMarried = true
	fmt.Printf("%#v\n", employee2)
	fmt.Println(employee2)
}

func mapExampleWithArray() {
	myMap := map[string][]string{
		"Yazar1": {"Kitap1", "Kitap2"},
		"Yazar2": {"Kitap3", "Kitap4"},
	}

	for yazar, kitaplari := range myMap {
		fmt.Println("Yazar -> ", yazar)
		fmt.Println("Kitapları :")
		for key, value := range kitaplari {
			fmt.Println(key+1, "-", value)
		}
	}
}

func copyMethodIsPassByValue() {
	mySlice := []int{1, 3, 5}
	mySlice2 := make([]int, 2)
	fmt.Println("mySlice ->", mySlice)
	fmt.Println("mySlice2 ->", mySlice2)
	copy(mySlice2, mySlice)
	fmt.Println("mySlice mySlice2'ye kopyalandı.")
	fmt.Println("mySlice ->", mySlice)
	fmt.Println("mySlice2 ->", mySlice2)
	mySlice[0] = 0
	fmt.Println("mySlice değişti")
	fmt.Println("mySlice ->", mySlice)
	fmt.Println("mySlice2 ->", mySlice2)
	fmt.Println("mySlice2 değişiklikten etkilenmedi")
	//mySlice2 değişiklikten etkilenmedi çünkü copy ile çoğaltılan slice'ler; copy olduğu slice'nin üzerine kurulduğu array'den değil yeni oluşturulan bir array üzerine kurulur
}

func mapRange() {
	myMap := map[string]int{
		"Ahmet": 40,
		"Veli":  18,
		"Kenan": 26,
		"İrem":  0,
	}

	for k, v := range myMap {
		fmt.Println(k, v)
	}
}

func copyMapIsPassByReference() {
	myMap := map[string]int{
		"Ahmet": 40,
		"Veli":  18,
		"Kenan": 26,
		"İrem":  0,
	}
	myMap2 := myMap
	fmt.Println("myMap ->", myMap)
	fmt.Println("myMap2 ->", myMap2)
	delete(myMap2, "Kenan")
	fmt.Println("myMap2'den Kenan silindi")
	fmt.Println("myMap ->", myMap)
	fmt.Println("myMap2 ->", myMap2)
	//değişiklikten myMap de etkilendi.
}

func deleteElementFromMap() {
	myMap := map[string]int{
		"Ahmet": 40,
		"Veli":  18,
		"Kenan": 26,
		"İrem":  0,
	}
	fmt.Println(myMap)
	delete(myMap, "Kenan")
	fmt.Println(myMap)
}

func addElementToMapAfterInitialize() {
	myMap := map[string]int{
		"Ahmet": 40,
		"Veli":  18,
		"Kenan": 26,
		"İrem":  0,
	}

	fmt.Println(myMap)
	myMap["Ceren"] = 15
	fmt.Println(myMap)
}

func mapExample() {
	myMap := map[string]int{
		"Ahmet": 40,
		"Veli":  18,
		"Kenan": 26,
		"İrem":  0,
	}

	fmt.Println(myMap["İremm"]) //bu eleman yok o yüzden 0 zerovalue veriyor. aslında irem değerimiz de 0 burada ufak bir açık oluşuyor.
	value, ok := myMap["İremm"] //bu açıktan kurtularak bir elemanın gerçekten olup olmadığını anlamak için ok değerini kullanabiliriz
	//_, ok := myMap["İremm"]
	fmt.Println(value, ok)
}

func declaredSliceAndDeclaredMakeSlice() {
	var mySlice []int
	fmt.Printf("%#v", mySlice)
	fmt.Println()
	mySlice2 := make([]int, 0)
	fmt.Printf("%#v", mySlice2)
}

func deleteSliceElementsFromEnd() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)
	mySlice = mySlice[:len(mySlice)-2]
	fmt.Println(mySlice)
}

func deleteSliceElementsFromStart() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)
	mySlice = mySlice[2:]
	fmt.Println(mySlice)
}

func appendSliceFromAnotherSliceAndMakeChangesOldSlice() {
	mySlice := []int{1, 3, 5}
	//otherSlice:=[]int{7,9};
	//mySlice2:=append(mySlice,otherSlice...);
	mySlice2 := append(mySlice, 7, 9)
	fmt.Println("mySlice ->", mySlice)
	fmt.Println("mySlice2 ->", mySlice2)
	mySlice[0] = 0
	fmt.Println("mySlice değişti")
	fmt.Println("mySlice ->", mySlice)
	fmt.Println("mySlice2 ->", mySlice2)
	fmt.Println("mySlice2 değişiklikten etkilenmedi")
	//mySlice2 değişiklikten etkilenmedi çünkü append ile türetilen slice'ler; türetildiği slice'nin üzerine kurulduğu array'den değil yeni oluşturulan bir array üzerine kurulur
}

func appendSlice() {
	mySlice := []int{1, 3, 5}
	fmt.Println(mySlice)
	mySlice = append(mySlice, 7, 9)
	fmt.Println(mySlice)
}

func copySliceFromArrayAndMakeChanges() {
	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mySlice := underArray[2:5]
	mySlice2 := underArray[3:]
	mySlice3 := underArray[:6]
	mySlice4 := underArray[:]
	fmt.Println("arr->", underArray)
	fmt.Println("mySlice->", mySlice)
	fmt.Println("mySlice2->", mySlice2)
	fmt.Println("mySlice3->", mySlice3)
	fmt.Println("mySlice4->", mySlice4)
	mySlice[0] = 100
	fmt.Println("mySlice değişti")
	fmt.Println("arr->", underArray)
	fmt.Println("mySlice->", mySlice)
	fmt.Println("mySlice2->", mySlice2)
	fmt.Println("mySlice3->", mySlice3)
	fmt.Println("mySlice4->", mySlice4)
	fmt.Println("sliceler değişince türediği array de değişti.array değişince ondan türeyen diğer slice'lar da değişti")
}

func copySliceIsPassByReference() {
	//golang pass by value'dir fakat slice'ler pass by reference'dir
	mySlice := []string{"ali", "veli"}
	mySlice2 := mySlice
	fmt.Println("ilk slice", mySlice)
	fmt.Println("ikinci slice", mySlice2)
	mySlice2[0] = "kerem"
	fmt.Println("ikinci slice değişti")
	fmt.Println("ilk slice", mySlice)
	fmt.Println("ikinci slice", mySlice2)
	fmt.Println("değişiklikten ilk slice etkilendi.")
}

func copyArrayIsPassByValue() {
	myArr := [3]int{1, 2, 3}
	myArr2 := myArr
	fmt.Println("ilk array", myArr)
	fmt.Println("ikinci array", myArr2)
	myArr2[0] = 4
	fmt.Println("ikinci array değişti")
	fmt.Println("ilk array", myArr)
	fmt.Println("ikinci array", myArr2)
	fmt.Println("değişiklikten ilk array etkilenmedi")
}

func slice() {
	//slice'ler array'lerin genişletilmiş halidir.

	mySliceWithLiteral := []int{1, 2, 3} //slice, literal yöntemi ile verdiğimiz elemanlar ile doluyor
	fmt.Println(mySliceWithLiteral)
	fmt.Println(len(mySliceWithLiteral))
	fmt.Println(cap(mySliceWithLiteral))

	mySliceWithMake := []int{}
	mySliceWithMake = make([]int, 2) //slice, make yöntemi ile verdiğimiz kapasite kadar int türünde boş elemanlar ile doluyor
	fmt.Println(mySliceWithMake)
	fmt.Println(len(mySliceWithMake))
	fmt.Println(cap(mySliceWithMake))
}

func arrays() {
	//cities := [3]string{"istanbul", "ankara", "izmir"}
	cities := [...]string{"istanbul", "ankara", "izmir"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	}

	for index, city := range cities {
		fmt.Println(index, city)
	}
}

func produceRandomNumber() {
	result := numRand(1, 100)
	println("random number -> ", result)
}

func numRand(min, max int) int {
	//rand.Seed(time.Now().Unix()) // döngüde kullansa idik her çalıştırdığımızda yeni bir rakam üretebilmesi için bu satırı yazardık.
	return rand.Intn(max-min) + min
}

func checkError() {
	result, err := evenNumber(10)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Girdiğiniz sayi ->", result)
	}
}

func evenNumber(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("Hata: Girdiğiniz sayı çift değil")
	}

	return num, nil //nil javadaki null
}

func doubleReturnFunction() { //sağdaki parantezde neyi döneceğimizi belirliyoruz.
	bolum, kalan := divideAction(104, 5)
	fmt.Println("bölüm ->", bolum)
	fmt.Println("kalan ->", kalan)
}

func divideAction(bolunen, bolen int) (int, int) { //sağdaki parantezde neyi döneceğimizi belirliyoruz.
	bolum := bolunen / bolen
	kalan := bolunen % bolen
	return bolum, kalan
}

func blankIdentifier() {
	fmt.Print("Bir sayı giriniz: ")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') //normalde ReadString methodu 2 tane değer return eder. Biz ikinci değeri blank identifier ile alıyoruz.
	println("girdiğiniz sayı ", value)
}

func readValueFromKeyboard() {
	fmt.Print("Bir sayı giriniz: ")
	reader := bufio.NewReader(os.Stdin)   //klavyeden girilen değeri okuyabilmek için reader oluşturuyoruz
	value, err := reader.ReadString('\n') //kod buraya geldiği andan alt satıra geçene kadar ki geçen süre boyunca girilen değerleri alıyoruz.
	println("girdiğiniz sayı ", value)
	println("hata  ", err)
}

func executeTwoFunctions() {
	x, y := 7, 15
	sum := sumTwoNumber(x, y)
	fmt.Println(sum)
}

func sumTwoNumber(x int, y int) int {
	//burada x ve y parametre
	return x + y
}

func nonIdiomaticIfStatement() {
	if x := 20; x%2 == 0 {
		fmt.Println("x çifttir")
	} else {
		println("x tektir.")
	}
}

func idiomaticIfStatement() {
	x := 20
	if x%2 == 0 {
		fmt.Println("x çifttir")
		return
	}
	fmt.Println("x tektir.")
}

func switchCasefallthrough() {
	//fallthrough true case çalıştıktan sonra switch ifadesinden çıkmaktansa öteki case'yi de kontrol eder.
	switch x := 25; {

	case x < 20:
		println("x 20'den küçük")
		fallthrough
	case x < 30:
		println("x 30'den küçük")
		fallthrough
	case x < 50:
		println("x 50'den küçük")
		break
	}
}

func forLoopAndGoto() {
	for i := 0; i < 10; i++ {
		if i == 7 {
			goto FALSE
		}
	}
FALSE:
	fmt.Println("go to work")
}

func switchCaseStatement() {
	name := "kerem"
	switch name {
	case "ali":
		println("name  ali -> " + name)
		break
	case "mehmet":
		println("name  mehmet ->" + name)
		break
	default:
		println("name ne mehmet ne ali ->" + name)
	}
}

func ifStatements() {
	x := 16
	if x%2 == 0 {
		println("mod 2 evet")
	} else {
		fmt.Println("mode 2 hayır")
	}
}

func strConv() {
	x := 65
	y := strconv.Itoa(x)

	fmt.Println(x)
	fmt.Println(y)
}

func shadowing() {
	x := 5
	if true {
		x := 10
		fmt.Println(x)
		//shadowing if scope içerisinde x değişkeni üzerinde shadowing yaptık.
	}
	fmt.Println(x)
}

func typeConversion() {
	x := 10.1
	y := 15.0

	fmt.Println(x + y)
	fmt.Println(int(x) + int(y))
}

func dataTypes() {
	name, surName, age, ogrenciMi := "ali", "demir", 25, false
	//0 dan buraya kadar
	var deger uint8 = 255
	var deger2 uint16 = 65535
	var deger3 uint32 = 4294967295
	var deger4 uint64 = 18446744073709551615

	//- ve + değerler dahil
	var sayi int8 = 127
	var sayi2 int16 = 32767
	var sayi3 int32 = 2147483647
	var sayi4 int64 = 9223372036854775807

	var ondalikli float32 = 12.34
	var ondalikli2 float64 = 12.34
	karisik := complex64(123.56)
	karisik2 := complex128(123.56)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", surName)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", ogrenciMi)
	fmt.Printf("%T\n", deger)
	fmt.Printf("%T\n", deger2)
	fmt.Printf("%T\n", deger3)
	fmt.Printf("%T\n", deger4)
	fmt.Printf("%v\n", sayi)
	fmt.Printf("%v\n", sayi2)
	fmt.Printf("%v\n", sayi3)
	fmt.Printf("%v\n", sayi4)
	fmt.Println(ondalikli)
	fmt.Println(ondalikli2)
	fmt.Println(karisik)
	fmt.Println(karisik2)
}

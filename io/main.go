package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	//writeFile()
	//writeFile2()
	//readFile()
	//readFileBufio()
	//readFileSeek()
	//readLineByLine()
	//multiReader()
	//multiWriter()
	//pipe()
	//teeReader()
}

func teeReader() {
	/*
		Normalde bir veriyi streamden alırken stream bir kere okunabilir.Çünkü okunduktan sonra da veriler kaybolur.
		TeeReader sayesinde okuma işlemi bittikten sonra verileri kaybetmeden datayı tekrar işleyebiliriz.
	*/
	sReader := strings.NewReader("test string")
	tReader := io.TeeReader(sReader, os.Stdout) //os.Stdout bir writer consol çıktısı verir.

	fmt.Println("starting")
	readedBytes, _ := io.ReadAll(tReader) //burada okudu ve io.TeeReader() func içine parametre geçtiğimiz writer'ı işleme soktu.
	fmt.Println("\nreaded string")
	fmt.Println(string(readedBytes))
}

func pipe() {
	reader, writer := io.Pipe()
	done := make(chan struct{})

	go read(reader, done)
	go write(writer)

	<-done
	//pipe bir veriyi yazarken aynı zamanda okumak için kullanılabilir.
}

func write(writer *io.PipeWriter) {
	/*
		i'yi gelen data olarak düşünelim
	*/
	i := 0
	for {
		if i == 10 {
			writer.Close()
			break
		}
		writer.Write([]byte(string(i)))
		i++
		time.Sleep(time.Millisecond * 100)
	}

}

func read(reader *io.PipeReader, done chan struct{}) {
	buff := make([]byte, 1024)
	/*
		bazen pipe read işleminde verilerin birazını okuduktan sonra hata alınır ve okuma kapanır. buna uygun bir okuma kodu yazdık.
	*/

	for {
		readed, err := reader.Read(buff)
		if readed == 0 {
			if err == io.EOF {
				done <- struct{}{}
				break
			}
			if err != nil {
				fmt.Println(err)
				done <- struct{}{}
				break
			}
		} else {
			fmt.Println(buff[:readed])

			if err == io.EOF || err != nil {
				fmt.Println(err)
				done <- struct{}{}
				break
			}
		}
	}
}

func multiWriter() {
	testFile, _ := os.Create("testFile.txt")
	testFile2, _ := os.Create("testFile2.txt")
	mWriter := io.MultiWriter(os.Stdout, testFile, testFile2)

	size, err := mWriter.Write([]byte("Multi writer example"))
	if err != nil {
		fmt.Println(err)
	}
	testFile.Close()
	testFile2.Close()
	fmt.Println(size) //yazılan metinin size bilgisi
}

func multiReader() {
	str1 := strings.NewReader("First Reader String\n")
	str2 := strings.NewReader("Second Reader String\n")
	mReader := io.MultiReader(str1, str2)

	io.Copy(os.Stdout, mReader)
}

func readLineByLine() {
	os.WriteFile("lines.txt", []byte("line 1\nline 2\nline 3\n"), os.ModePerm)

	testFile, err := os.Open("lines.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(testFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readFileSeek() {
	f, err := os.Open("testFile.txt")
	if err != nil {
		fmt.Println(err)
	}
	f.Seek(5, 1)
	readByte := make([]byte, 4)
	f.Read(readByte)

	fmt.Println(string(readByte))
}

func readFileBufio() {
	readTest, err := os.Open("testFile.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		bufReader := bufio.NewReader(readTest)
		io.Copy(os.Stdout, bufReader)
	}

}

func readFile() {
	//file tamamı döndü
	readTest, err := os.ReadFile("testFile.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(readTest))
	}
}

func writeFile() {
	//varsa açar öncekileri siler yoksa oluşturur params -> fileName,veri,dosyayı hangi yetki ile açacak
	err := os.WriteFile("testFile.txt", []byte("1234567890"), os.ModePerm) //modeperm sınırsız yetki gibi bir yetki
	if err != nil {
		fmt.Println(err)
	}
}

func writeFile2() {
	f, err := os.Create("testFile.txt")
	if err != nil {
		fmt.Println(err)
	}

	f.Write([]byte("1234567890\n"))
}

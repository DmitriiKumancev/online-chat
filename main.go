package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// listener, который прослушивает порт.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Println("Ошибка прослушивания порта: 2000: Ошибка: ", err.Error())
	}

	for {
		// принимаем соединения
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Ошибка: %s при подключении к адресу: %s\n",
				err.Error(), conn.RemoteAddr())
		} else {
			log.Printf("Успешное подключение к удаленному адресу: %s\n",
				conn.RemoteAddr())
		}

		// обработка записи на сервер
		go HandleIncoming(conn)

		// обработка чтения с сервера
		go HandleOutgoing(conn)

	}
}

// обработчик входящих запросов
func HandleIncoming(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte("Ваше соединение с нашим сервером успешно!\n"))
	for {
		// буфер для хранения сообщений на стороне клиента
		buf := make([]byte, 64)

		if err != nil {
			log.Printf("Запись не удалась")
		}
		// считываем клиента и сохраняем в buf
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Соединение разорвано!")
		}

		// выводим на консоль сервера
		fmt.Printf("client: %s\n", string(buf[:n]))

	}
}

func HandleOutgoing(conn net.Conn) {
	defer conn.Close()
	for {
		// чтение из локальной консоли
		reader := bufio.NewReader(os.Stdin);

		line, err := reader.ReadString('\n' );
		if err != nil {
			log.Printf("Ошибка чтения строки\n" );
		}

		// запись в консоль клиента 
		_, err = conn.Write([] byte(fmt.Sprintf("server: %s\n", line)));
		if err != nil {
			log.Printf("Подключение к адресу: %s закрыто.\n" , conn.RemoteAddr());
		}

	}
}

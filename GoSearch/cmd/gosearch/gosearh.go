package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"go-core-4/GoSearch/pkg/crawler"
	"go-core-4/GoSearch/pkg/crawler/membot"
	"go-core-4/GoSearch/pkg/crawler/spider"
	"go-core-4/GoSearch/pkg/index"
)

func main() {
	searchFlag := flag.String("s", "", "слово для поиска")
	useMemBot := flag.Bool("mem", false, "использовать тестовые данные membot вместо сканирования сети")
	flag.Parse()

	var docs []crawler.Document
	var err error

	if *useMemBot {
		fmt.Println("Используем локальный эмулятор membot...")
		bot := membot.New()
		docs, err = bot.Scan("", 1)
		if err != nil {
			log.Fatalf("ошибка сканирования membot: %v", err)
		}
	} else {
		sites := []string{
			"https://go.dev",
			"https://golang.org",
		}

		sp := spider.New()
		fmt.Println("Сканирование сайтов...")
		for _, site := range sites {
			fmt.Printf("Сканирование: %s (глубина 1)...\n", site)
			siteDocs, scanErr := sp.Scan(site, 1)
			if scanErr != nil {
				fmt.Printf("Предупреждение: не удалось отсканировать %s: %v\n", site, scanErr)
				continue
			}
			docs = append(docs, siteDocs...)
		}

		// Если сеть недоступна или нет документов, используем membot в качестве fallback
		if len(docs) == 0 {
			fmt.Println("Сетевое сканирование не вернуло документов. Переключаемся на membot...")
			bot := membot.New()
			docs, _ = bot.Scan("", 1)
		}
	}

	// 1. Добавляем уникальный номер (ID) для каждого документа
	for i := range docs {
		docs[i].ID = i
	}

	// 2. Сортируем срез документов по номерам (ID) с помощью стандартной библиотеки
	sort.Slice(docs, func(i, j int) bool {
		return docs[i].ID < docs[j].ID
	})

	fmt.Printf("Всего проиндексировано документов: %d\n", len(docs))

	// 3. Создаем и заполняем обратный индекс из пакета index
	idx := index.New()
	idx.Add(docs)

	// 4. Поиск по запросу
	if *searchFlag != "" {
		performSearch(idx, docs, *searchFlag)
		return
	}

	// Интерактивный режим поиска через консоль
	fmt.Println("\n=== Поисковый движок GoSearch ===")
	fmt.Println("Введите слово для поиска (или 'exit' для выхода):")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}
		query := strings.TrimSpace(scanner.Text())
		if query == "" {
			continue
		}
		if strings.ToLower(query) == "exit" || strings.ToLower(query) == "quit" {
			fmt.Println("Завершение работы.")
			break
		}

		performSearch(idx, docs, query)
	}
}

// performSearch осуществляет поиск слова по индексу, а затем
// находит документы с помощью бинарного поиска по отсортированному срезу документов.
func performSearch(idx *index.Service, docs []crawler.Document, word string) {
	// Задача №2: Поиск по индексу возвращает срез номеров (ID) документов
	docIDs := idx.Search(word)
	if len(docIDs) == 0 {
		fmt.Printf("По запросу %q ничего не найдено.\n", word)
		return
	}

	fmt.Printf("Найдено совпадений: %d\n", len(docIDs))

	// Задача №3: Бинарный поиск по отсортированному срезу документов по их ID
	for _, id := range docIDs {
		indexPos := index.BinarySearch(docs, id)
		if indexPos != -1 {
			doc := docs[indexPos]
			fmt.Printf("  [%d] %s — %s\n", doc.ID, doc.Title, doc.URL)
		}
	}
}

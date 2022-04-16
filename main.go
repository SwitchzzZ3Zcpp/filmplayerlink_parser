package main

import (
    //"fmt"
    "github.com/gocolly/colly"
    "strconv"
    "os"
    "github.com/briandowns/spinner"
    //"strings"
    "bufio"
	"log"
	"time"
    // "database/sql"
    // _ "github.com/mattn/go-sqlite3"
    
    
)

func main() {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) 

        s.Color("bgBlack", "bold", "fgGreen")

        s.Prefix = "Обновляем список фильмов по жанру 'Horror' " 
        s.Start()
        time.Sleep(2 * time.Second)
    
	c := colly.NewCollector(
        colly.AllowedDomains("new.lordfilms-s.art"),
    )

    c.OnHTML(".th-item", func(e *colly.HTMLElement) {
       	links := e.ChildAttrs("a", "href")

		file, err := os.OpenFile("filmlinks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		 
			if err != nil {
				log.Fatalf("Не удалось создать файл: %s", err)
			}
		 
			datawriter := bufio.NewWriter(file)
		 
			for _, data := range links {
				_, _ = datawriter.WriteString(data + "\n")
			}
		 
			datawriter.Flush()
			file.Close()
    })
      
    for pages := 1; pages < 320; pages++ {
        c.Visit("http://new.lordfilms-s.art/f-uzhas/page/" + strconv.Itoa(pages) + "/")        
    }
    s.Stop()
    frameParse()
    //readFile()
}

func frameParse() {

		file, err := os.Open("filmlinks.txt")
	  
	    if err != nil {
	        log.Fatalf("Я не смог открыть :(")
	  
	    }

	    scanner := bufio.NewScanner(file)
	    
	    scanner.Split(bufio.ScanLines)
	    var text []string
	  
	    for scanner.Scan() {
	        text = append(text, scanner.Text())
	    }
	  
	    file.Close()
	  
	    ///////////////////////////////////////

	 b := colly.NewCollector(
     	 colly.AllowedDomains("new.lordfilms-s.art"),
     )

		s := spinner.New(spinner.CharSets[43], 100*time.Millisecond) 

        s.Color("bgBlack", "bold", "fgGreen")

        s.Prefix = "Сохраняем ссылки на плееры фильмов по жанру 'Horror' " 
        s.Start()
        time.Sleep(2 * time.Second)

     b.OnHTML(".tabs-b", func(e *colly.HTMLElement) {
         links := e.ChildAttrs("iframe", "src")
         // fmt.Println("\nВот ссылка на плеер фильма: ", links)

		file, err := os.OpenFile("fplayerlinks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				 
					if err != nil {
						log.Fatalf("Не удалось создать файл: %s", err)
					}
				 
					datawriter := bufio.NewWriter(file)
				 
					for _, data := range links {
						_, _ = datawriter.WriteString(data + "\n")
					}
				 
					datawriter.Flush()
					file.Close()
     })

	 for _, each_ln := range text {
     	b.Visit(each_ln)
	 }
	 s.Stop()
    
 }

// func readFile() {
	// file, err := os.Open("filmlinks.txt")
	  // 
	    // if err != nil {
	        // log.Fatalf("failed to open")
	  // 
	    // }
// 
	    // scanner := bufio.NewScanner(file)
	    // 
	    // scanner.Split(bufio.ScanLines)
	    // var text []string
	  // 
	    // for scanner.Scan() {
	        // text = append(text, scanner.Text())
	    // }
	  // 
	    // file.Close()
	  // 
	    // for _, each_ln := range text {
	        // fmt.Println(each_ln)
	    // }
// }

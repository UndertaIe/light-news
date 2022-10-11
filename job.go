package main

import (
	"log"
	"sync"
	"time"
)

type Job func()

func NewJob(r *Rule) Job {
	job := func() {
		now := time.Now()
		log.Printf("job[DataSource: %s] is running", r.DataSource)
		ps := parserRegistry.GetParser(r.Parser)
		models, err := ps.Parse(r)
		if err != nil {
			log.Printf("parser.Parse(r) args:[%v], err msg:[%v]\n", r, err)
		}
		err = dataCheck.Check(models)
		if err != nil {
			log.Printf("dataCheck.Check(models) args:[%v], err msg:[%v]\n", r.DataSource, err)
		}
		syncStore(r, models)
		log.Printf("job[DataSource: %s] has pawned %v items, elapsed time: %v", r.DataSource, len(models), time.Since(now))
	}
	return job
}

// 连接数过多，运行多个go协程时将*gorm.DB复制了多次
func asyncStore(r *Rule, models []*NewsModel) {
	store := storerRegistry.GetStorer(r.StoreType)
	var wg sync.WaitGroup
	for _, m := range models {
		wg.Add(1)
		go func(mm *NewsModel) {
			err := store.Store(mm)
			if err != nil {
				log.Printf("asyncStore args:[%v], err msg:[%v]\n", mm, err)
			}
			wg.Done()
		}(m)
	}
	wg.Wait()
}

func syncStore(r *Rule, models []*NewsModel) {
	store := storerRegistry.GetStorer(r.StoreType)
	for _, m := range models {
		err := store.Store(m)
		if err != nil {
			log.Printf("asyncStore args:[%v], err msg:[%v]\n", m, err)
		}
	}
}

var storerRegistry StorerRegistry

type StorerRegistry struct {
	stores sync.Map // {StoreType: Storer}
}

func (sr *StorerRegistry) GetStorer(t StoreType) Storer {
	if storerI, ok := sr.stores.Load(t); ok {
		storer, _ := storerI.(Storer)
		return storer
	}
	storer, err := SelectStorer(t)
	if err != nil {
		log.Fatalf("SelectStorer(t) err args:[%v], msg:[%v]\n", t, err)
	}
	log.Println("init store success. StoreType: ", t)
	sr.stores.Store(t, storer)
	return storer
}

var parserRegistry ParserRegistry

type ParserRegistry struct {
	stores sync.Map // {StoreType: Storer}
}

func (pr *ParserRegistry) GetParser(t ParserType) Parser {
	if parserI, ok := pr.stores.Load(t); ok {
		ps, _ := parserI.(Parser)
		return ps
	}
	ps, err := SelectParser(t)
	if err != nil {
		log.Fatalf("SelectParser(t) err args:[%v], msg:[%v]\n", t, err)
	}
	pr.stores.Store(t, ps)
	return ps
}

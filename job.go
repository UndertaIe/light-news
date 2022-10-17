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
			return
		}
		err = dataCheck.Check(models)
		if err != nil {
			log.Printf("dataCheck.Check(models) args:[%v], err msg:[%v]\n", r.DataSource, err)
			return
		}
		asyncWrite(r, models)
		log.Printf("job[DataSource: %s] has pawned %v items, write to %v, elapsed time: %v", r.DataSource, len(models), r.StoreType, time.Since(now))
	}
	return job
}

func syncWrite(store Storer, r *Rule, models []*NewsModel) {
	for _, m := range models {
		err := store.Store(m)
		if err != nil {
			log.Printf("syncStore args:[%v], err msg:[%v]\n", m, err)
		}
	}
}

func asyncWrite(r *Rule, models []*NewsModel) {
	stores := storerRegistry.GetStorers(r.StoreType)
	var wg sync.WaitGroup
	for _, store := range stores {
		go func(store Storer) {
			defer wg.Done()
			syncWrite(store, r, models)
		}(store)

		wg.Add(1)
	}
	wg.Wait()
}

var storerRegistry StorerRegistry

type StorerRegistry struct {
	stores sync.Map // {StoreType: Storer}
}

func (sr *StorerRegistry) GetStorers(st []StoreType) []Storer {
	var storers []Storer
	for _, t := range st {
		if storerI, ok := sr.stores.Load(t); ok {
			storers = append(storers, storerI.(Storer))
			continue
		}
		storer, err := SelectStorer(t)
		if err != nil {
			log.Fatalf("SelectStorer(t) err args:[%v], msg:[%v]\n", t, err)
		}
		log.Println("init store success. StoreType: ", t)
		sr.stores.Store(t, storer)
		storers = append(storers, storer)
	}
	return storers
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

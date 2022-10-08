package main

import (
	"log"
	"sync"

	"github.com/robfig/cron/v3"
)

var (
	DefaultScheduler *Scheduler = NewScheduler()
)

type Scheduler struct {
	rules    sync.Map // URL: *Rule
	jobs     sync.Map // URL: cron.EntryID
	cr       *cron.Cron
	addCh    chan *Rule
	deleteCh chan *Rule
	updateCh chan *Rule
}

func NewScheduler() *Scheduler {
	var sch = &Scheduler{
		cr:       cron.New(cron.WithSeconds()),
		addCh:    make(chan *Rule, 1),
		deleteCh: make(chan *Rule, 1),
		updateCh: make(chan *Rule, 1),
	}
	sch.loop()
	return sch
}

func (s *Scheduler) Start() {
	s.cr.Start()

}
func (s *Scheduler) loop() {
	go func() {
		for {
			var r *Rule
			select {
			case r = <-s.addCh:
				s.schedule(r)
			case r = <-s.updateCh:
				s.removeJob(r)
				s.schedule(r)
			case r = <-s.deleteCh:
				s.removeJob(r)
			}
		}
	}()
}

func (s *Scheduler) AddJob(r Rule) {
	s.addCh <- &r
}

func (s *Scheduler) UpdateJob(r Rule) {
	s.updateCh <- &r
}

func (s *Scheduler) DeleteJob(r Rule) {
	s.deleteCh <- &r
}

func (s *Scheduler) schedule(r *Rule) {
	if _, ok := s.rules.Load(r.Key()); ok {
		s.UpdateJob(*r)
		return
	}
	// 将Job添加到cron
	eid, err := s.cr.AddFunc(r.Cron, NewJob(r))
	if err != nil {
		log.Println("s.cr.AddFunc(r.Cron, NewJob(r)) error: ", eid, err)
	}
	log.Printf("s.cr.AddFunc(r.Cron, NewJob(r)) msg: %v, entityID: %v\n", r, eid)
	s.jobs.Store(r.Key(), eid)
	s.rules.Store(r.Key(), r)
}
func (s *Scheduler) removeJob(r *Rule) {
	jobKey := r.Key()
	eid, ok := s.jobs.Load(jobKey)
	if !ok {
		log.Println("job not found, Rule: ", r)
		return
	}
	jobid, _ := eid.(cron.EntryID)
	s.cr.Remove(jobid)
	s.jobs.Delete(r.Key())
	s.rules.Delete(r.Key())
}

func (s *Scheduler) ListJob(r Rule) []Rule {
	var rules []Rule
	f := func(k, v any) bool {
		vv, _ := v.(*Rule)
		rules = append(rules, *vv)
		return true
	}
	s.rules.Range(f)
	return rules
}

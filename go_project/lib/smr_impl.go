package qse

import (
	log "github.com/sirupsen/logrus"
)

func NewSMRConfig[
	ATOM comparable,
	IDENT any,
	SORT any,
	MODEL any,
	SCTX SMTSolvedContext[MODEL],
	SYS SMTSystem[
		IdLiteral[ATOM],
		IDENT,
		SORT,
		MODEL,
		SCTX,
	],
](
	in_canidates chan SMRDNFClause[ATOM, IDENT, SORT],
	out_models chan MODEL,
	sys SYS,
) (
	smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS],
) {
	smr_config = SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]{
		in_canidates: in_canidates,
		out_models:   out_models,
		sys:          sys,
		unfinished:   NewSMRUnfinishedArray[ATOM, IDENT, SORT](),
	}
	return
}

func NewSMRUnfinishedArray[
	ATOM comparable,
	IDENT any,
	SORT any,
]() (unfinished SMRUnfinishedArray[ATOM, IDENT, SORT]) {
	backing_nocopy := TrustingNoCopySMRUnfinishedArray[ATOM, IDENT, SORT]{
		arr: make([]SMRDNFClause[ATOM, IDENT, SORT], 0),
	}
	unfinished = SMRUnfinishedArray[ATOM, IDENT, SORT]{
		&backing_nocopy,
	}
	return
}

func NewSMRIsSleeping() (is_sleeping SMRIsSleeping) {
	backing_nocopy := TrustingNoCopySMRIsSleeping{
		is: false,
	}
	is_sleeping = SMRIsSleeping{
		&backing_nocopy,
	}
	return
}

func (is_sleeping SMRIsSleeping) Check() (is bool) {
	is_sleeping.mu.Lock()
	defer is_sleeping.mu.Unlock()
	is = is_sleeping.is
	return
}

func (is_sleeping SMRIsSleeping) Sleep() (was bool) {
	is_sleeping.mu.Lock()
	defer is_sleeping.mu.Unlock()
	was = is_sleeping.is
	is_sleeping.is = true
	return
}

func (is_sleeping SMRIsSleeping) Wake() (was bool) {
	is_sleeping.mu.Lock()
	defer is_sleeping.mu.Unlock()
	was = is_sleeping.is
	is_sleeping.is = false
	return
}

func (smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]) Start() {
	wakeup_chan := make(chan struct{})
	is_sleeping := NewSMRIsSleeping()
	eternal_slumber := NewSMRIsSleeping()
	eternal_slumber.Wake()
	go func() {
		defer log.Info("[smr/go1] Entered eternal slumber. ")
		defer close(smr_config.out_models)
		defer close(wakeup_chan)
	runSMRLoop:
		for {
			log.Info("[smr/go1] SMR iteration goroutine awake. ")
			if eternal_slumber.Check() {
				log.Info("[smr/go1] Performing all current work before eternal slumber. ")
				for !smr_config.RunSMR() {
				}
				log.Info("[smr/go1] Entering eternal slumber. ")
				break runSMRLoop
			}
			for !smr_config.RunSMR() {
			}
			is_sleeping.Sleep()
			log.Info("[smr/go1] SMR iteration goroutine asleep. ")
			<-wakeup_chan
		}
	}()
	go func() {
		defer func() {
			log.Info("[smr/go2] Preparing for eternal slumber. ")
			eternal_slumber.Sleep()
			wakeup_chan <- struct{}{}
			log.Info("[smr/go2] Requesting eternal slumber. ")
		}()
		log.Info("[smr/go2] Waiting for canidates. ")
		for canidate := range smr_config.in_canidates {
			log.Info("[smr/go2] Received canidate. ")
			smr_config.unfinished.Append(canidate)
			if !is_sleeping.Wake() {
				log.Info("[smr/go2] Waking up SMR iteration goroutine. ")
				wakeup_chan <- struct{}{}
			} else {
				log.Info("[smr/go2] SMR iteration goroutine is already awake. ")
			}
		}
	}()
}

func (unfinished *TrustingNoCopySMRUnfinishedArray[ATOM, IDENT, SORT]) Append(
	elems ...SMRDNFClause[ATOM, IDENT, SORT],
) {
	unfinished.mu.Lock()
	defer unfinished.mu.Unlock()
	unfinished.arr = append(unfinished.arr, elems...)
}

func (smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]) RunSMR() (done bool) {
	log.Info("[smr/SMRConfig.RunSMR] Acquiring lock. ")
	smr_config.unfinished.mu.Lock()
	defer smr_config.unfinished.mu.Unlock()
	done = (len(smr_config.unfinished.arr) == 0)
	if done {
		log.Info("[smr/SMRConfig.RunSMR] Lock acquired, no work right now. ")
		return
	}
	log.Info("[smr/SMRConfig.RunSMR] Lock aquired, running an SMR iteration. ")
	smr_config.SMRIterationUnfinishedUnlocked()
	return
}

func (smr_config SMRConfig[ATOM, IDENT, SORT, MODEL, SCTX, SYS]) SMRIterationUnfinishedUnlocked() {
	finished := make([]int, 0)
	for i := range smr_config.unfinished.arr {
		elem := smr_config.unfinished.arr[i]
		combined := make([]IdLiteral[ATOM], 0)
		combined = append(combined, elem.conjunction_r...)
		combined = append(combined, elem.conjunction_f...)
		sctx := smr_config.sys.CheckSat(
			combined,
			elem.free_funs,
		)
		is_sat_ptr := sctx.IsSat()
		if is_sat_ptr == nil {
			log.Info("[smr/SMRConfig.SMRIteration...] Formula invalid. ")
			finished = append(finished, i)
			continue
		}
		if *is_sat_ptr {
			log.Info("[smr/SMRConfig.SMRIteration...] Formula sat. Sending.")
			smr_config.out_models <- *sctx.GetModel()
			finished = append(finished, i)
		} else {
			log.Info("[smr/SMRConfig.SMRIteration...] Formula unsat. ")
			mus := *sctx.ExtractMUS()
			InsertionSortInPlace(mus)
			DedupSortedInPlace(&mus)
			offset := 0
		smrReductionLoop:
			for _, index := range mus {
				if index >= len(elem.conjunction_r) {
					continue smrReductionLoop
				}
				index := index - offset
				SpliceOutReclaim(&elem.conjunction_r, index)
				offset++
			}
		}
	}
	InsertionSortInPlace(finished)
	DedupSortedInPlace(&finished)
	offset := 0
	for _, index := range finished {
		index := index - offset
		SpliceOutReclaim(
			&smr_config.unfinished.arr,
			index,
		)
		offset++
	}
}

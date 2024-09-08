package cronutil

import (
	"context"
	"errors"
	"github.com/robfig/cron/v3"
	"log"
	"sync"
)

var (
	c         *cron.Cron
	onceStart *sync.Once
	onceStop  *sync.Once
)

func init() {
	c = cron.New()
	onceStart = &sync.Once{}
	onceStop = &sync.Once{}
}

func AddJob(spec string, jobFunc func()) error {
	if jobFunc == nil {
		return errors.New("job function is nil")
	}

	entryId, err := c.AddFunc(spec, jobFunc)
	if err != nil {
		return err
	}
	log.Printf("Cron job (%v) with spec '%v' attached", entryId, spec)

	return nil
}

func RunInBackground() {
	onceStart.Do(func() {
		log.Printf("[CRON] Starting cron service in the background...")
		c.Start()
		log.Printf("[CRON] Cron service started")
	})
}

func StopFromBackground(ctx context.Context) {
	onceStop.Do(func() {
		log.Printf("[CRON] Stoping cron service from the background...")
		cronStopCtx := c.Stop()
		select {
		case <-ctx.Done():
			log.Printf("[CRON] Cron service failed to stop. %v", ctx.Err())
		case <-cronStopCtx.Done():
			log.Printf("[CRON] Cron service stopped")
		}
	})
}

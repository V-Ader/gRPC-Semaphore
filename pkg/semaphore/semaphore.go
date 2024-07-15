package semaphore

import (
	"fmt"
	"log"
	"sync"
)

type Semaphore struct {
	permits int
	mutex   sync.Mutex
}

func NewSemaphore(permits int) *Semaphore {
	return &Semaphore{permits: permits}
}

func (semaphore *Semaphore) Increase(increaseValue int) {

	semaphore.mutex.Lock()
	defer semaphore.mutex.Unlock()

	semaphore.permits += increaseValue
	log.Printf("Increased by %d. current value: %d\n", increaseValue, semaphore.permits)

}

func (semaphore *Semaphore) Decrease(decreaseValue int) error {
	semaphore.mutex.Lock()
	defer semaphore.mutex.Unlock()

	if semaphore.permits >= decreaseValue {
		semaphore.permits -= decreaseValue
		log.Printf("Decreased by %d. current value: %d\n", decreaseValue, semaphore.permits)

		return nil
	}
	return fmt.Errorf("not enough permits: requested %d, available %d", decreaseValue, semaphore.permits)
}

package counter

import "sync"

// CounterState representa o estado do contador
type CounterState struct {
	Value int
}

// CounterObserver define como os componentes serão notificados de mudanças
type CounterObserver interface {
	OnCounterChanged(state CounterState)
}

// CounterService contém a lógica de negócio do contador e gerencia observers
type CounterService struct {
	state     CounterState
	observers []CounterObserver
	mutex     sync.RWMutex
}

// NewCounterService cria uma nova instância do serviço
func NewCounterService() *CounterService {
	return &CounterService{
		state:     CounterState{Value: 0},
		observers: make([]CounterObserver, 0),
	}
}

// AddObserver adiciona um observer para receber notificações
func (cs *CounterService) AddObserver(observer CounterObserver) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	cs.observers = append(cs.observers, observer)
}

// RemoveObserver remove um observer
func (cs *CounterService) RemoveObserver(observer CounterObserver) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	for i, obs := range cs.observers {
		if obs == observer {
			cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
			break
		}
	}
}

// Increment aumenta o contador em 1
func (cs *CounterService) Increment() {
	cs.mutex.Lock()
	cs.state.Value++
	cs.mutex.Unlock()
	cs.notifyObservers()
}

// Decrement diminui o contador em 1
func (cs *CounterService) Decrement() {
	cs.mutex.Lock()
	cs.state.Value--
	cs.mutex.Unlock()
	cs.notifyObservers()
}

// Reset zera o contador
func (cs *CounterService) Reset() {
	cs.mutex.Lock()
	cs.state.Value = 0
	cs.mutex.Unlock()
	cs.notifyObservers()
}

// GetValue retorna o valor atual (thread-safe)
func (cs *CounterService) GetValue() int {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	return cs.state.Value
}

// GetState retorna uma cópia do estado completo (thread-safe)
func (cs *CounterService) GetState() CounterState {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	return CounterState{Value: cs.state.Value}
}

// SetValue define um valor específico
func (cs *CounterService) SetValue(value int) {
	cs.mutex.Lock()
	cs.state.Value = value
	cs.mutex.Unlock()
	cs.notifyObservers()
}

// notifyObservers notifica todos os observers sobre mudanças
func (cs *CounterService) notifyObservers() {
	cs.mutex.RLock()
	stateCopy := CounterState{Value: cs.state.Value}
	observers := make([]CounterObserver, len(cs.observers))
	copy(observers, cs.observers)
	cs.mutex.RUnlock()

	// Notifica fora do lock para evitar deadlocks
	for _, observer := range observers {
		observer.OnCounterChanged(stateCopy)
	}
}

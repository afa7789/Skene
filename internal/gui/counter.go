package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/afa7789/skene/internal/counter"
	"github.com/afa7789/skene/internal/localization"
)

// CounterWidget representa o widget do contador na GUI
type CounterWidget struct {
	// Widgets da UI
	valueLabel   *widget.Label
	incrementBtn *widget.Button
	decrementBtn *widget.Button
	resetBtn     *widget.Button
	container    *fyne.Container

	// Serviço de negócio (injetado)
	counterService *counter.CounterService
}

// NewCounterWidget cria um novo widget de contador
func NewCounterWidget(counterService *counter.CounterService) *CounterWidget {
	cw := &CounterWidget{}

	// Usa o serviço fornecido (injeção de dependência)
	cw.counterService = counterService

	// Se adiciona como observer para receber notificações
	cw.counterService.AddObserver(cw)

	// Inicializa os widgets
	cw.initWidgets()
	cw.layoutWidgets()

	return cw
}

// initWidgets inicializa todos os widgets da UI
func (cw *CounterWidget) initWidgets() {
	// Label para mostrar o valor
	cw.valueLabel = widget.NewLabel(cw.formatValue(0))

	// Botões com callbacks que chamam o serviço
	cw.incrementBtn = widget.NewButton(localization.T("counter_increment"), func() {
		cw.counterService.Increment() // Chama serviço
	})

	cw.decrementBtn = widget.NewButton(localization.T("counter_decrement"), func() {
		cw.counterService.Decrement() // Chama serviço
	})

	cw.resetBtn = widget.NewButton(localization.T("counter_reset"), func() {
		cw.counterService.Reset() // Chama serviço
	})
}

// layoutWidgets organiza os widgets no layout
func (cw *CounterWidget) layoutWidgets() {
	// Container com botões horizontais
	buttonsContainer := container.NewHBox(
		cw.decrementBtn,
		cw.resetBtn,
		cw.incrementBtn,
	)

	// Container principal vertical
	cw.container = container.NewVBox(
		widget.NewCard(
			localization.T("counter_title"),
			"",
			container.NewVBox(
				cw.valueLabel,
				buttonsContainer,
			),
		),
	)
}

// GetContainer retorna o container principal
func (cw *CounterWidget) GetContainer() *fyne.Container {
	return cw.container
}

// OnCounterChanged implementa counter.CounterObserver
// Este método é chamado pelo serviço quando o estado muda
func (cw *CounterWidget) OnCounterChanged(state counter.CounterState) {
	// Atualiza apenas a apresentação
	cw.valueLabel.SetText(cw.formatValue(state.Value))
}

// formatValue formata o valor para exibição
func (cw *CounterWidget) formatValue(value int) string {
	return fmt.Sprintf(localization.T("counter_value"), value)
}

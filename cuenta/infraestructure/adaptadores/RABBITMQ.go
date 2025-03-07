package adaptadores

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/streadway/amqp"
)

// Estructura de la transacción
type Transaction struct {
	ID        string  `json:"id"`
	From      int     `json:"from"`
	To        int     `json:"to"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
	Timestamp string  `json:"timestamp"`
}

// Publicar una transacción en RabbitMQ y esperar respuesta
func PublishTransaction(transaction Transaction) (bool, error) {
	// 1️⃣ Conexión a RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@13.217.71.115:5672/")
	if err != nil {
		return false, fmt.Errorf("Error conectando a RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return false, fmt.Errorf("Error abriendo canal: %v", err)
	}
	defer ch.Close()

	// 2️⃣ Declarar la cola de transacciones
	_, err = ch.QueueDeclare("transactions", true, false, false, false, nil)
	if err != nil {
		return false, fmt.Errorf("Error declarando la cola de transacciones: %v", err)
	}

	// 3️⃣ Convertir la transacción a JSON
	body, err := json.Marshal(transaction)
	if err != nil {
		return false, fmt.Errorf("Error serializando JSON: %v", err)
	}

	// 4️⃣ Publicar la transacción
	err = ch.Publish(
		"",
		"transactions",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return false, fmt.Errorf("Error publicando mensaje: %v", err)
	}

	fmt.Println("📤 Transacción enviada:", transaction.ID)

	// 5️⃣ Esperar respuesta en `transactions_responses`
	responseChan := make(chan bool)

	go func() {
		respConn, _ := amqp.Dial("amqp://guest:guest@13.217.71.115:5672/")
		defer respConn.Close()

		respCh, _ := respConn.Channel()
		defer respCh.Close()

		// Declarar la cola de respuestas
		queue, _ := respCh.QueueDeclare("transactions_responses", true, false, false, false, nil)

		msgs, _ := respCh.Consume(queue.Name, "", true, false, false, false, nil)

		for msg := range msgs {
			var response map[string]interface{}
			_ = json.Unmarshal(msg.Body, &response)

			if response["id"] == transaction.ID {
				status := response["status"].(string)
				if status == "success" {
					responseChan <- true
				} else {
					responseChan <- false
				}
				break
			}
		}
	}()

	// 6️⃣ Manejo de tiempo de espera
	select {
	case success := <-responseChan:
		if success {
			fmt.Println("✅ Confirmación recibida. La transferencia fue exitosa.")
			return true, nil
		} else {
			fmt.Println("❌ Error en el procesamiento de la transacción.")
			return false, fmt.Errorf("Error en el procesamiento")
		}
	case <-time.After(10 * time.Second): // Tiempo máximo de espera
		fmt.Println("⏳ No se recibió respuesta en el tiempo esperado.")
		return false, fmt.Errorf("Timeout esperando respuesta")
	}
}

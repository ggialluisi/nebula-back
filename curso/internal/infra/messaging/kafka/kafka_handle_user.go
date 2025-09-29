package kafka

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/ggialluisi/nebula-back/curso/internal/domain/dto"
	"github.com/ggialluisi/nebula-back/curso/internal/domain/repository"
	"github.com/ggialluisi/nebula-back/curso/internal/domain/usecase"
)

type PessoaKafkaHandlers struct {
	PessoaRepository repository.PessoaRepositoryInterface
}

func NewPessoaKafkaHandlers(
	PessoaRepository repository.PessoaRepositoryInterface,
) *PessoaKafkaHandlers {
	return &PessoaKafkaHandlers{
		PessoaRepository: PessoaRepository,
	}
}

// Handler para mensagens Sarama
func (h *PessoaKafkaHandlers) CreateOrUpdatePessoa(msg *sarama.ConsumerMessage) error {
	var inputDto dto.PessoaInputDTO

	err := json.Unmarshal(msg.Value, &inputDto)
	if err != nil {
		log.Printf("❌ Erro ao decodificar mensagem: %v", err)
		return err
	}

	uc := usecase.NewSavePessoaUseCase(h.PessoaRepository)

	_, err = uc.ExecuteCreateOrUpdatePessoa(inputDto)
	if err != nil {
		log.Printf("❌ Erro ao executar usecase Pessoa: %v", err)
		return err
	}

	log.Printf("✅ Pessoa criada/atualizada: %+v", inputDto)
	return nil
}

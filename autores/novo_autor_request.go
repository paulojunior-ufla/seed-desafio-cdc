package autores

import "jornada-dev-eficiente/casa-do-codigo/validator"

// Pilar: não ligamos parâmetros de requisição externa com objetos de domínio diretamente, assim como não serializamos objetos de domínio para respostas de API.
type NovoAutorRequest struct {
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	Descricao string `json:"descricao"`
}

// Pilar: protegemos as bordas do sistemas como se não houvesse amanhã. Principalmente a mais externa
func (r *NovoAutorRequest) Validate(v *validator.Validator) {
	v.Check(r.Nome != "", "Nome", "Não pode ser vazio")
	v.Check(r.Email != "", "Email", "Não pode ser vazio")
	v.CheckEmail(r.Email, "Email", "Deve ser válido")
	v.Check(r.Descricao != "", "Descrição", "Não pode ser vazia")
	v.Check(len(r.Descricao) <= 400, "Descrição", "Não pode conter mais de 400 caracteres")
}

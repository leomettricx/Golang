package main

type Paciente struct {
	Id            string `json:"id"`
	Nome          string `json:"nome"`
	Endereco      string `json:"endereco"`
	Cep           int    `json:"cep"`
	Cidade        string `json:"cidade"`
	Estado        string `json:"estado"`
	Telefone1     string `json:"telefone1"`
	Telefone2     string `json:"telefone2"`
	Email         string `json:"email"`
	Nascimento    string `json:"nascimento"`
	CPF           string `json:"cpf"`
	TipoSanguineo string `json:"tiposanguineo"`
}

type SeguroSaude struct {
	Id       string `json:"id"`
	Validade string `json:"validade"`
	Tipo     string `json:"tipo"`
	Numero   int    `json:"numero"`
	Status   string `json:"status"`
}

type Receita struct {
	Id         string `json:"id"`
	Nome       string `json:"nome"`
	Dosagem    string `json:"dosagem"`
	Frequencia string `json:"frequencia"`
}

type Consulta struct {
	Id          string `json:"id"`
	Tipo        string `json:"tipo"`
	Data        string `json:"data"`
	Hora        string `json:"hora"`
	Responsavel string `json:"responsavel"`
	Status      string `json:"status"`
	Descricao   string `json:"descricao"`
}

type Exames struct {
	Id            string   `json:"id"`
	Tipo          string   `json:"tipo"`
	Data          string   `json:"data"`
	Laboratorio   string   `json:"local"`
	Status        string   `json:"status"`
	Justificativa string   `json:"justificativa"`
	Itens         []string `json:"itens"`
}

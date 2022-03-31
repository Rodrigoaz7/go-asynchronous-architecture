package endereco

type Endereco struct {
	Rua    string `json:"rua"`
	Cep    string `json:"cep"`
	Bairro string `json:"bairro"`
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
	Pais   string `json:"pais"`
}

type Message struct {
	Data            Endereco `json:"data"`
	MessageSoftware string   `json:"message_software"`
}

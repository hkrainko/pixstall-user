package model

type ListJWTCredentialsResponse struct {
	Data []JWTCredential `json:"data"`
}

type JWTCredential struct {
	CreatedAt    int64       `json:"created_at"`
	ID           string      `json:"id"`
	Tags         interface{} `json:"tags"`
	Secret       string      `json:"secret"`
	RSAPublicKey interface{} `json:"rsa_public_key"`
	Consumer     Consumer    `json:"consumer"`
	Key          string      `json:"key"`
	Algorithm    string      `json:"algorithm"`
}

type Consumer struct {
	ID string `json:"id"`
}

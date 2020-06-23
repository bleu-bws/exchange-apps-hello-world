package entities

import "time"

// Token -
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// Profile -
type Profile struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  struct {
		Comprovanterenda      int         `json:"comprovanterenda"`
		Comprovanteresidencia int         `json:"comprovanteresidencia"`
		TaxIDNumber           string      `json:"taxIDNumber"`
		Datanascimento        time.Time   `json:"datanascimento"`
		Email                 string      `json:"email"`
		Emailverificado       bool        `json:"emailverificado"`
		Emailexecucao         bool        `json:"emailexecucao"`
		Emaillogin            bool        `json:"emaillogin"`
		Emailorderbook        bool        `json:"emailorderbook"`
		Fotodocumento         int         `json:"fotodocumento"`
		Fotodocumentoback     int         `json:"fotodocumentoback"`
		Googleauth            bool        `json:"googleauth"`
		Nome                  string      `json:"nome"`
		Endereco              string      `json:"endereco"`
		Cep                   string      `json:"cep"`
		Complemento           interface{} `json:"complemento"`
		Cidade                string      `json:"cidade"`
		Estado                string      `json:"estado"`
		Pais                  string      `json:"pais"`
		Selfie                int         `json:"selfie"`
		Telefone              string      `json:"telefone"`
		Zencard               int         `json:"zencard"`
		Numero                interface{} `json:"numero"`
		Bairro                interface{} `json:"bairro"`
		Fee                   int         `json:"fee"`
		MarketMaker           bool        `json:"market_maker"`
		UserLevel             struct {
			Level              string `json:"level"`
			UserData           bool   `json:"userData"`
			TwoFA              bool   `json:"twoFA"`
			Email              bool   `json:"email"`
			Phone              bool   `json:"phone"`
			Document           bool   `json:"document"`
			Selfie             bool   `json:"selfie"`
			ProofOfResidence   bool   `json:"proofOfResidence"`
			ProofOfIncome      bool   `json:"proofOfIncome"`
			Corporate          bool   `json:"corporate"`
			WithdrawLimitDay   int    `json:"withdrawLimitDay"`
			WithdrawLimitMonth int    `json:"withdrawLimitMonth"`
		} `json:"userLevel"`
		Corporative        bool `json:"corporative"`
		WithdrawLimitDay   int  `json:"withdraw_limit_day"`
		WithdrawLimitMonth int  `json:"withdraw_limit_month"`
	} `json:"result"`
}

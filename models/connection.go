package models

type Instance struct {
	ID               int    `json:"id"`
	Apikey           string `json:"apikey"`
	Phone            string `json:"phone"`
	Name             string `json:"name"`
	Label            string `json:"label"`
	Platform         string `json:"platform"`
	Status           string `json:"status"`
	Version          string `json:"version"`
	TariffID         int    `json:"tariff_id"`
	TariffPlane      string `json:"tariff_plane"`
	PipelineID       int    `json:"pipelineId"`
	StageID          int    `json:"stageId"`
	ChatID           string `json:"chat_id"`
	InstanceID       string `json:"instanceId"`
	ChatToken        string `json:"chat_token"`
	ChatKey          string `json:"chat_key"`
	DateAdd          int    `json:"date_add"`
	DateTrial        int    `json:"date_trial"`
	DatePay          int    `json:"date_pay"`
	DateSubscription int    `json:"date_subscription"`
	IsPremium        int    `json:"is_premium"`
	IsGroup          bool   `json:"is_group"`
	Strategy         string `json:"strategy"`
	ResponsibleID    []int  `json:"responsibleId"`
	CreateNewIfClose bool   `json:"createNewIfClose"`
	WID              string `bson:"wid"`
	LanguageCode     string `bson:"language_code"`
	Pushname         string `bson:"pushname,omitempty"`
}

type Status struct {
	State string `json:"state"`
}

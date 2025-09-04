package Ads

import (
	"Vk-api-go/types"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type UpdateClientOrdSubagent struct {
	OrdSubagentType string `json:"type"` //Тип подрядчика. Возможные значения:
	//   • person — физическое лицо.
	//   • individual — индивидуальный предприниматель.
	//   • legal — юридическое лицо.
	Name  string `json:"name"`  //Название подрядчика.
	Inn   string `json:"inn"`   //ИНН подрядчика. Длина строки: 10 или 12 символов.
	Phone string `json:"phone"` //Телефон подрядчика. Длина строки: от 7 до 20 символов.
}

// OrdDataSpecification Объект описывает оператора рекламных данных.
type UpdateClientOrdDataSpecification struct {
	ClientType string `json:"client_type"` //Тип клиента. Возможные значения:
	//• person — физическое лицо.
	//• individual — индивидуальный предприниматель.
	//• legal — юридическое лицо.
	ClientName     string                  `json:"client_name"`     //Полное имя клиента или название организации клиента.
	Inn            string                  `json:"inn"`             //ИНН клиента. Длина строки: 10 или 12 символов.
	Phone          string                  `json:"phone"`           //Телефон клиента. Длина строки: от 7 до 20 символов.
	AgencyPhone    string                  `json:"agency_phone"`    //Номер телефона агентства. Длина строки: от 7 до 20 символов. Доступно с версии 5.194
	Subagent       UpdateClientOrdSubagent `json:"subagent"`        //Данные о подрядчике в формате JSON (ord_subagent).
	ContractNumber string                  `json:"contract_number"` //Номер контракта.
	ContractDate   string                  `json:"contract_date"`   //Дата контракта в формате DD.MM.YYYY.
	ContractType   string                  `json:"contract_type"`   //Тип контракта. Примеры:
	//   • Договор оказания услуг.
	//   • Дополнительное соглашение.
	//   • Посреднический договор.
	ContractObject    string `json:"contract_object"`     //Предмет контракта. Вы можете скопировать информацию из текста договора — пункт «Предмет договора» — или кратко опишите суть договора (что должны сделать заказчик и исполнитель).
	WithVat           bool   `json:"with_vat"`            //Информация о том, подлежит ли НДС оплате в рамках этого контракта.
	IsSubagentEnabled bool   `json:"is_subagent_enabled"` //false, если нужно удалить данные об имеющемся подрядчике
}

// UpdateClientsSpecification Объект описывает клиента рекламного агентства.
type UpdateClientsSpecification struct {
	Name     string                           `json:"name"`      //Название клиента. Длина строки: от 3 до 60 символов.
	DayLimit int                              `json:"day_limit"` //Дневной лимит в рублях.
	AllLimit int                              `json:"all_limit"` //Общий лимит в рублях.
	OrdData  UpdateClientOrdDataSpecification `json:"ord_data"`  //Данные оператора рекламных данных в формате JSON (ord_data_specification).
}

// UpdateClientsRequest структура для передачи в функцию
type UpdateClientsRequest struct {
	UpdateClientsSpecification []UpdateClientsSpecification `json:"client_specification"`
}

// UpdateClientsSerialize сериализует массив UpdateClientsSpecification в JSON
func UpdateClientsSerialize(UpdateClientsSpecification []UpdateClientsSpecification) (string, error) {
	request := UpdateClientsRequest{
		UpdateClientsSpecification: UpdateClientsSpecification,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации: %w", err)
	}

	return string(jsonData), nil
}

// UpdateClients Метод создаёт клиентов рекламного агентства. Доступен только для рекламных агентств.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) UpdateClients(ctx context.Context, accountID int, data string) (types.VkResponse, error) {
	params := url.Values{}
	params.Set("account_id", strconv.Itoa(accountID))
	params.Set("data", data)
	VkRequest := types.VkRequest{
		Method: "ads.updateClients",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}

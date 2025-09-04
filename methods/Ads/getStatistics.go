package Ads

import (
	"Vk-api-go/types"
	"context"
	"net/url"
	"strconv"
)

type GetStatisticsOptions struct {
	accountId int    //Идентификатор рекламного кабинета.
	idsType   string //Тип запрашиваемых объектов, которые перечислены в параметре ids:
	//• ad — объявления;
	//• campaign — кампании;
	//• client — клиенты;
	//• office — кабинет.
	ids string //Перечисленные через запятую ID запрашиваемых объявлений,
	// кампаний, клиентов или кабинета, в зависимости от того, что указано в параметре ids_type. Максимум 2000 объектов.
	period string //Способ группировки данных по датам:
	//• day — статистика по дням;
	//• week — статистика по неделям;
	//• month — статистика по месяцам;
	//• year — статистика по годам;
	//• overall — статистика за всё время;
	//Временные ограничения задаются параметрами date_from и date_to.
	dateFrom string //Начальная дата выводимой статистики. Используется разный формат дат для разного значения параметра period:
	//• day: YYYY-MM-DD, пример: 2011-09-27 - 27 сентября 2011.
	//• 0 — день создания;
	//• week: YYYY-MM-DD, пример: 2011-09-27 - считаем статистику, начиная с понедельника той недели, в которой находится заданный день.
	//• 0 — неделя создания;
	//• month: YYYY-MM, пример: 2011-09 - сентябрь 2011.
	//• 0 — месяц создания;
	//• year: YYYY, пример: 2011 - 2011 год.
	//• 0 — год создания;
	//• overall: 0
	dateTo string //Конечная дата выводимой статистики. Используется разный формат дат для разного значения параметра period:
	//• day: YYYY-MM-DD, пример: 2011-09-27 - 27 сентября 2011.
	//• week: YYYY-MM-DD, пример: 2011-09-27 - считаем статистику до воскресения той недели, в которой находится данный день.
	//• 0 — текущая неделя;
	//• month: YYYY-MM, пример: 2011-09 - сентябрь 2011.
	//• 0 — текущий месяц;
	//• year: YYYY, пример: 2011 - 2011 год.
	//• 0 — текущий год;
	//• overall: 0
	statsFields string //Дополнительные статистики:
	//• views_times — распределение количества показов на пользователя. Доступно для типов ad и campaign, содержащих только рекламные записи.
}

type GetStatisticsOption func(*GetStatisticsOptions)

func GetStatisticsWithStatsFields(statsFields string) GetStatisticsOption {
	return func(o *GetStatisticsOptions) {
		o.statsFields = statsFields
	}
}

// GetStatistics Возвращает статистику показателей эффективности по рекламным объявлениям, кампаниям, клиентам или всему кабинету.
// Для вызова метода можно использовать:
// •ключ доступа пользователя (требуются права доступа: ads)
func (am *AddMethods) GetStatistics(ctx context.Context, accountId int, idsType string, ids string, period string, dateFrom string, dateTo string, opts ...GetStatisticsOption) (types.VkResponse, error) {

	options := &GetStatisticsOptions{}

	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}

	params.Set("account_id", strconv.Itoa(accountId))
	params.Set("ids_type", idsType)
	params.Set("ids", ids)
	params.Set("period", period)
	params.Set("date_from", dateFrom)
	params.Set("date_to", dateTo)

	if options.statsFields != "" {
		params.Set("stats_fields", options.statsFields)
	}

	VkRequest := types.VkRequest{
		Method: "ads.getStatistics",
		Params: params,
	}
	return am.methods.Call(ctx, VkRequest)
}

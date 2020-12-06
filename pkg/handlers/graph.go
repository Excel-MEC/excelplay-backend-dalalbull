package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/httperrors"
	"github.com/Excel-MEC/excelplay-backend-dalalbull/pkg/strconst"
)

// Graph sends back stock history data of a specific company to be displayed in a graph on the client
func Graph(db *database.DB, env *env.Config) httperrors.Handler {
	type company struct {
		Symbol string `json:"company"`
	}
	type datapoint struct {
		Time         int     `json:"X"`
		CurrentPrice float32 `json:"Y"`
	}
	type data struct {
		Datapoints []datapoint `json:"graph_data"`
	}

	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		var c company
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.InvalidJSON, http.StatusBadRequest}
		}

		var stockHistoryData []database.StockHistory
		err = db.GetStockDataHistory(c.Symbol, &stockHistoryData)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not get stock data of the company", http.StatusInternalServerError}
		}

		var graphData data
		for i := len(stockHistoryData) - 1; i >= 0; i-- {
			var d datapoint
			d.Time = stockHistoryData[i].Time.Hour()*60 + stockHistoryData[i].Time.Minute()
			d.CurrentPrice = stockHistoryData[i].CurrentPrice
			graphData.Datapoints = append(graphData.Datapoints, d)
		}

		jsonRes, err := json.Marshal(graphData)
		if err != nil {
			return &httperrors.HTTPError{r, err, strconst.JSONEncodingFail, http.StatusInternalServerError}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonRes)
		return nil
	}
}

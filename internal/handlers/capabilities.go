package handlers

import (
	"net/http"
)

func GetCapabilitiesHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Формируем параметры сервера")
	// var layers []models.Layer
	// db := db.GetConnection()
	// if err := db.Find(&layers).Error; err != nil {
	// 	http.Error(w, "DB not work", http.StatusBadRequest)
	// 	return
	// }

	// // Собираем стили для каждого слоя
	// layerStyles := make(map[uint][]models.Style)
	// var styles []models.Style
	// if err := db.Find(&styles).Error; err == nil {
	// 	for _, style := range styles {
	// 		layerStyles[style.LayerID] = append(layerStyles[style.LayerID], style)
	// 	}
	// }

	// w.Header().Set("Content-Type", "application/xml")
	// xml.NewEncoder(w).Encode(wmts)
}

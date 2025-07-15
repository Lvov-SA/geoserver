package loader

// import (
// 	"fmt"
// 	"log"

// 	"github.com/dhconnelly/rtreego"
// 	"github.com/lukeroth/gdal"
// )

// type GeoTIFFTile struct {
// 	X, Y, Z int
// 	Data    []float64
// }

// type rtreeTile struct {
// 	z, x, y int
// 	bounds  rtreego.Rect
// }

// func (t *rtreeTile) Bounds() rtreego.Rect {
// 	return t.bounds
// }

// var (
// 	geoDataset *gdal.Dataset
// 	tileCache  map[string]GeoTIFFTile
// 	rtree      *rtreego.Rtree
// )

// func init() {
// 	// 1. Открываем GeoTIFF
// 	dataset, err := gdal.Open("../resourse/map/NE1_LR_LC.tif", gdal.ReadOnly)
// 	if err != nil {
// 		log.Fatal("Ошибка загрузки GeoTIFF:", err)
// 	}
// 	geoDataset = &dataset

// 	// 2. Инициализируем кеш и R-дерево
// 	tileCache = make(map[string]GeoTIFFTile)
// 	rtree = rtreego.NewTree(2, 25, 50)

// 	// 3. Получаем размеры растра
// 	xSize := geoDataset.RasterXSize()
// 	ySize := geoDataset.RasterYSize()
// 	gt := geoDataset.GeoTransform()

// 	// 4. Разбиваем на тайлы (пример для z=0)
// 	tileSize := 256
// 	for x := 0; x < xSize; x += tileSize {
// 		for y := 0; y < ySize; y += tileSize {
// 			// Вычисляем границы в координатах растра
// 			minX := gt[0] + float64(x)*gt[1] + float64(y)*gt[2]
// 			maxX := gt[0] + float64(x+tileSize)*gt[1] + float64(y)*gt[2]
// 			minY := gt[3] + float64(x)*gt[4] + float64(y)*gt[5]
// 			maxY := gt[3] + float64(x)*gt[4] + float64(y+tileSize)*gt[5]

// 			rect, _ := rtreego.NewRect(
// 				rtreego.Point{minX, minY},
// 				[]float64{maxX - minX, maxY - minY},
// 			)

// 			tile := &rtreeTile{
// 				z:      0,
// 				x:      x / tileSize,
// 				y:      y / tileSize,
// 				bounds: rect,
// 			}
// 			rtree.Insert(tile)
// 		}
// 	}
// }

// func Check() {
// 	fmt.Println("Инициализация завершена. Тайлов в R-дереве:", rtree.Size())
// }

// func getTile(z, x, y int) (GeoTIFFTile, error) {
// 	// Проверяем кеш
// 	key := fmt.Sprintf("%d/%d/%d", z, x, y)
// 	if tile, ok := tileCache[key]; ok {
// 		return tile, nil
// 	}

// 	// 1. Получаем bbox тайла
// 	bbox := tileToGeoTIFFBounds(z, x, y)

// 	// 2. Читаем данные из GeoTIFF
// 	width := 256  // Ширина тайла
// 	height := 256 // Высота тайла
// 	buffer := make([]float64, width*height)

// 	// Читаем данные первого канала (для многоканальных растров нужен цикл)
// 	band := geoDataset.RasterBand(1)
// 	err := band.IO(gdal.Read, int(bbox.MinX), int(bbox.MinY), width, height, buffer, width, height, 0, 0)
// 	if err != nil {
// 		return GeoTIFFTile{}, err
// 	}

// 	// 3. Сохраняем в кеш
// 	tile := GeoTIFFTile{
// 		X:    x,
// 		Y:    y,
// 		Z:    z,
// 		Data: buffer,
// 	}
// 	tileCache[key] = tile

// 	return tile, nil
// }

// // Переводит координаты тайла (z,x,y) в bbox в системе координат GeoTIFF
// func tileToGeoTIFFBounds(z, x, y int) gdal.b {
// 	// Получаем геотрансформ (перевод пиксельных координат в географические)
// 	gt := geoDataset.GeoTransform()

// 	// Размер тайла в пикселях (например, 256x256)
// 	tileSize := 256
// 	minX := gt[0] + float64(x*tileSize)*gt[1]
// 	maxX := gt[0] + float64((x+1)*tileSize)*gt[1]
// 	minY := gt[3] + float64(y*tileSize)*gt[5]
// 	maxY := gt[3] + float64((y+1)*tileSize)*gt[5]

// 	return gdal.Bounds{MinX: minX, MaxX: maxX, MinY: minY, MaxY: maxY}
// }

// internal/cli/flags.go
package cli

import (
	"flag"
	"csv-parser/pkg"
	"csv-parser/internal/config"
)

func ParseFlags() *config.Config {
	var (
		filePath     string
		requiredRaw  string
		validateRaw  string
		rangeRaw     string
		verbose      bool
		filterRow    string
		sortRaw      string
		header       bool
	)

	flag.StringVar(&filePath, "file", "", "Путь к CSV-файлу")
	flag.StringVar(&requiredRaw, "required", "", "Обязательные поля (через запятую)")
	flag.StringVar(&validateRaw, "validate-type", "", "Проверка типов (пример: Age:int,Price:float)")
	flag.StringVar(&rangeRaw, "range", "", "Диапазоны значений (пример: Age:18-65)")
	flag.BoolVar(&verbose, "verbose", false, "Вывод подробной информации")
	flag.StringVar(&filterRow, "filter", "", "Фильтрация по строке (пример: Age>100)")
	flag.StringVar(&sortRaw, "sort", "", "Сортировка по колонке (пример: Age:desc или Name:asc)")
	flag.BoolVar(&header, "header", true, "Учитывать ли заголовки в первой строке")


	flag.Parse()

	return &config.Config{
		FilePath:     filePath,
		Required:     pkg.ParseCommaList(requiredRaw),
		ValidateType: pkg.ParseKeyValueMap(validateRaw, ":"),
		Range:        pkg.ParseRangeMap(rangeRaw),
		Verbose:      verbose,
		Filter:    	  filterRow,
		Sort:         sortRaw,
		Header:       header,
	}
}

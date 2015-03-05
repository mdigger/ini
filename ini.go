package ini

import (
	"bufio"
	"io"
	"strings"
)

// Section описывает секцию ini-файла в виде ключ - значение.
type Section map[string]string

// Config описывает разобранное содержимое .ini файла в виде раздел - ключ - значение.
type Config map[string]Section

// Parse читает содержимое конфигурационного ini-файла и возвращает его в разобранном виде.
func Parse(ini io.Reader) (Config, error) {
	section := Section{} // текущая секуция
	config := Config{    // пустое описание конфигурации
		"": section, // определяем корневую секцию
	}
	scanner := bufio.NewScanner(ini) // инициализируем чтение из потока
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())          // читаем строку и избавляемся от пробелов в начале и конце строки
		if idx := strings.IndexAny(line, "#;"); idx > -1 { // удаляем комментарии
			line = strings.TrimSpace(line[:idx])
		}
		if line == "" { // пропускаем пустые строки
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") { // разбираем раздел
			sectionName := strings.TrimSpace(line[1 : len(line)-1]) // новое имя секции
			if config[sectionName] == nil {                         // создаем секцию, если ее не было
				config[sectionName] = Section{}
			}
			section = config[sectionName] // запоминаем текущую секцию
			continue
		}
		keyValue := strings.SplitN(line, "=", 2)     // разделяем ключ = значение
		if len(keyValue) != 2 || keyValue[0] == "" { // игнорируем пустые ключи
			continue
		}
		key := strings.TrimSpace(keyValue[0])   // название ключа
		value := strings.TrimSpace(keyValue[1]) // значение ключа
		section[key] = value                    // сохраняем ключ и значение в текущей секции
	}
	if err := scanner.Err(); err != nil { // в случае ошибки чтения возвращаем ее описание
		return nil, err
	}
	return config, nil
}

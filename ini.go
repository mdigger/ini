package ini

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Section описывает секцию ini-файла в виде ключ - значение.
type Section map[string]string

// Config описывает разобранное содержимое .ini файла в виде раздел - ключ - значение.
type Config map[string]Section

// Parse читает содержимое конфигурационного ini-файла и возвращает его в разобранном виде.
func Parse(ini io.Reader) (Config, error) {
	config := Config{
		"": Section{}, // корневая секция
	}
	scanner := bufio.NewScanner(ini)
	sectionName := ""
	for scanner.Scan() {
		switch line := strings.TrimSpace(scanner.Text()); {
		case strings.HasPrefix(line, ";"): // комментарий - пропускаем
		case strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]"): // раздел
			sectionName = strings.TrimSpace(line[1 : len(line)-1])
			config[sectionName] = Section{}
		default: // ключ = значение
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 && parts[0] != "" {
				config[sectionName][strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning ini: %v", err)
	}
	return config, nil
}

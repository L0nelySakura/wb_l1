package main

import (
	"fmt"
)

// Универсальный интерфейс для всех баз данных
type Database interface {
    Query(sql string) []string
}

// Реализация MySQL
type MySQL struct{}

func (m *MySQL) ExecuteMySQLQuery(query string) []map[string]string {
    fmt.Printf("MySQL выполняет запрос: %s\n", query)
    return []map[string]string{
        {"id": "1", "name": "John"},
        {"id": "2", "name": "Jane"},
    }
}


// Реализация PostgreSQL
type PostgreSQL struct{}

func (p *PostgreSQL) RunPGQuery(cmd string) map[string]interface{} {
    fmt.Printf("PostgreSQL выполняет: %s\n", cmd)
    return map[string]interface{}{
        "rows": []string{"row1", "row2", "row3"},
        "count": 3,
    }
}

type MySQLAdapter struct {
    mysql *MySQL
}

func NewMySQLAdapter() *MySQLAdapter {
    return &MySQLAdapter{mysql: &MySQL{}}
}

func (a *MySQLAdapter) Query(sql string) []string {
    // Конвертируем результат MySQL в универсальный формат
    result := a.mysql.ExecuteMySQLQuery(sql)
    
    var rows []string
    for _, row := range result {
        rows = append(rows, fmt.Sprintf("MySQL: %v", row))
    }
    return rows
}

// PostgresAdapter адаптирует PostgreSQL
type PostgresAdapter struct {
    postgres *PostgreSQL
}
func NewPostgresAdapter() *PostgresAdapter {
    return &PostgresAdapter{postgres: &PostgreSQL{}}
}

func (a *PostgresAdapter) Query(sql string) []string {
    result := a.postgres.RunPGQuery(sql)
    
    rows := result["rows"].([]string)
    var converted []string
    for _, row := range rows {
        converted = append(converted, fmt.Sprintf("PostgreSQL: %s", row))
    }
    return converted
}

// Таким образом мы можем приводить разные базы данных разных видов к 1 общему случаю
// Из плюсов можно выделить: упрощает тестирование, позволяет повторно использовать уже написанный код
// Из минусов можно выделить: снижает производительность, добовляет дополнительные классы
package where

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type Expr interface {
	toQueryAndArgs() (query interface{}, args []interface{})
}

type Eq map[string]interface{}

func (b Eq) toQueryAndArgs() (query interface{}, args []interface{}) {
	var exprs []string

	sortedKeys := getSortedKeys(b)
	for _, key := range sortedKeys {
		val := b[key]
		exprs = append(exprs, fmt.Sprintf("%s = ?", key))
		args = append(args, val)
	}

	query = strings.Join(exprs, " AND ")
	return
}

type Neq map[string]interface{}

func (b Neq) toQueryAndArgs() (query interface{}, args []interface{}) {
	var exprs []string

	sortedKeys := getSortedKeys(b)
	for _, key := range sortedKeys {
		val := b[key]
		exprs = append(exprs, fmt.Sprintf("%s != ?", key))
		args = append(args, val)
	}

	query = strings.Join(exprs, " AND ")
	return
}

type Gt map[string]interface{}

func (b Gt) toQueryAndArgs() (query interface{}, args []interface{}) {
	var exprs []string

	sortedKeys := getSortedKeys(b)
	for _, key := range sortedKeys {
		val := b[key]
		exprs = append(exprs, fmt.Sprintf("%s > ?", key))
		args = append(args, val)
	}

	query = strings.Join(exprs, " AND ")
	return
}

type GtOrEq map[string]interface{}

func (b GtOrEq) toQueryAndArgs() (query interface{}, args []interface{}) {
	var exprs []string

	sortedKeys := getSortedKeys(b)
	for _, key := range sortedKeys {
		val := b[key]
		exprs = append(exprs, fmt.Sprintf("%s >= ?", key))
		args = append(args, val)
	}

	query = strings.Join(exprs, " AND ")
	return
}

type Lt map[string]interface{}

func (b Lt) toQueryAndArgs() (query interface{}, args []interface{}) {
	var exprs []string

	sortedKeys := getSortedKeys(b)
	for _, key := range sortedKeys {
		val := b[key]
		exprs = append(exprs, fmt.Sprintf("%s < ?", key))
		args = append(args, val)
	}

	query = strings.Join(exprs, " AND ")
	return
}

type LtOrEq map[string]interface{}

func (b LtOrEq) toQueryAndArgs() (query interface{}, args []interface{}) {
	var exprs []string

	sortedKeys := getSortedKeys(b)
	for _, key := range sortedKeys {
		val := b[key]
		exprs = append(exprs, fmt.Sprintf("%s <= ?", key))
		args = append(args, val)
	}

	query = strings.Join(exprs, " AND ")
	return
}

type Between map[string][]interface{}

func (b Between) toQueryAndArgs() (query interface{}, args []interface{}) {
	var exprs []string

	sortedKeys := getSortedKeys2(b)
	for _, key := range sortedKeys {
		val := b[key]
		exprs = append(exprs, fmt.Sprintf("%s BETWEEN ? AND ?", key))
		args = append(args, val...)
	}

	query = strings.Join(exprs, " AND ")
	return
}

type In map[string]interface{}

func (b In) toQueryAndArgs() (query interface{}, args []interface{}) {
	var exprs []string

	sortedKeys := getSortedKeys(b)
	for _, key := range sortedKeys {
		val := b[key]
		exprs = append(exprs, fmt.Sprintf("%s IN (?)", key))
		args = append(args, val)
	}

	query = strings.Join(exprs, " AND ")
	return
}

type Like map[string]interface{}

func (b Like) toQueryAndArgs() (query interface{}, args []interface{}) {
	var exprs []string

	sortedKeys := getSortedKeys(b)
	for _, key := range sortedKeys {
		val := b[key]
		exprs = append(exprs, fmt.Sprintf("%s LIKE ?", key))
		args = append(args, val)
	}

	query = strings.Join(exprs, " AND ")
	return
}

type Or []Expr

func (b Or) toQueryAndArgs() (query interface{}, args []interface{}) {
	var exprs []string

	for _, expr := range b {
		expr, arg := expr.toQueryAndArgs()
		exprs = append(exprs, fmt.Sprintf("%s", expr))
		args = append(args, arg...)
	}

	query = "(" + strings.Join(exprs, " OR ") + ")"
	return
}

type And []Expr

func (b And) toQueryAndArgs() (query interface{}, args []interface{}) {
	var exprs []string

	for _, expr := range b {
		expr, arg := expr.toQueryAndArgs()
		exprs = append(exprs, fmt.Sprintf("%s", expr))
		args = append(args, arg...)
	}

	query = "(" + strings.Join(exprs, " AND ") + ")"
	return
}

func ToQueryAndArgs(exprs []Expr) (string, []interface{}) {
	if len(exprs) == 0 {
		return "", nil
	}

	var buf bytes.Buffer
	var args []interface{}

	for _, e := range exprs {
		expr, arg := e.toQueryAndArgs()
		buf.WriteString(expr.(string))
		buf.WriteString(" AND ")
		args = append(args, arg...)
	}

	query := buf.String()
	return query[:len(query)-5], args
}

func getSortedKeys(exp map[string]interface{}) []string {
	sortedKeys := make([]string, 0, len(exp))
	for k := range exp {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

func getSortedKeys2(exp map[string][]interface{}) []string {
	sortedKeys := make([]string, 0, len(exp))
	for k := range exp {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

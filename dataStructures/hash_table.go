package dataStructures

import (
	"fmt"
)

const Size = 10 // Размер массива хеш-таблицы

// Entry представляет элемент хеш-таблицы
type Entry struct {
	key   string
	value string
	next  *Entry
}

// HashTable структура для хранения данных
type HashTable struct {
	buckets [Size]*Entry
}

// хеш-функция: вычисляет индекс для ключа
func hash(key string) int {
	hash := 0
	for _, ch := range key {
		hash += int(ch)
	}
	return hash % Size
}

// Метод вставки данных в хеш-таблицу
func (ht *HashTable) Insert(key, value string) {
	index := hash(key)
	entry := ht.buckets[index]

	// Если нет элементов в данном индексе, добавляем новый
	if entry == nil {
		ht.buckets[index] = &Entry{key: key, value: value}
		return
	}

	// Ищем, существует ли уже элемент с таким ключом
	for entry != nil {
		if entry.key == key {
			entry.value = value // Обновляем значение, если ключ уже существует
			return
		}
		if entry.next == nil {
			break
		}
		entry = entry.next
	}

	// Добавляем новый элемент в цепочку
	entry.next = &Entry{key: key, value: value}
}

// Метод получения данных по ключу
func (ht *HashTable) Get(key string) (string, bool) {
	index := hash(key)
	entry := ht.buckets[index]

	// Ищем элемент с данным ключом
	for entry != nil {
		if entry.key == key {
			return entry.value, true
		}
		entry = entry.next
	}
	return "", false // Ключ не найден
}

// Метод удаления элемента по ключу
func (ht *HashTable) Delete(key string) bool {
	index := hash(key)
	entry := ht.buckets[index]

	if entry == nil {
		return false
	}

	// Если первый элемент в списке имеет нужный ключ
	if entry.key == key {
		ht.buckets[index] = entry.next
		return true
	}

	// Ищем элемент с данным ключом
	prev := entry
	for entry != nil {
		if entry.key == key {
			prev.next = entry.next
			return true
		}
		prev = entry
		entry = entry.next
	}

	return false // Ключ не найден
}

func works() {
	ht := &HashTable{}

	// Добавляем элементы
	ht.Insert("name", "Gleb")
	ht.Insert("age", "17")
	ht.Insert("city", "Minsk")

	// Получаем элементы
	if value, ok := ht.Get("name"); ok {
		fmt.Println("Name:", value)
	} else {
		fmt.Println("Name not found")
	}

	if value, ok := ht.Get("city"); ok {
		fmt.Println("City:", value)
	} else {
		fmt.Println("City not found")
	}

	// Удаляем элемент
	if ht.Delete("age") {
		fmt.Println("Age deleted")
	} else {
		fmt.Println("Age not found")
	}

	// Проверяем после удаления
	if value, ok := ht.Get("age"); ok {
		fmt.Println("Age:", value)
	} else {
		fmt.Println("Age not found")
	}
}

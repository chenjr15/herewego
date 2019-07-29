package maps

import (
	"sort"
	"testing"
)

func TestMapDrinks(t *testing.T) {
	drinkMap := map[string]string{
		"Juice":     "果汁",
		"Lemon tea": "肥仔快乐茶",
		"Yogurt":    "酸奶",
		"Coffee":    "咖啡",
		"Milk":      "牛奶",
		"Cola":      "肥仔快乐水",
	}
	t.Log("All Drinks:")
	for k, _ := range drinkMap {
		t.Log(k)
	}
	keys := make([]string, len(drinkMap))
	i := 0
	t.Log("\nDrinks -> Translation:")
	for k, v := range drinkMap {
		t.Logf("%s \t-> \t%s ", k, v)
		keys[i] = k
		i++

	}

	t.Log("\nSorted:")
	sort.Strings(keys)
	for _, k := range keys {
		t.Logf("%s \t->\t %s ", k, drinkMap[k])

	}
}
